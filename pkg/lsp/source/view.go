// Copyright 2019 The Protocol Buffers Language Server Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-language-server/uri"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/parser"
	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// View represents a single workspace.
// Views are managed by a session. A view accesses files.
type View interface {
	FileSystem

	// Session returns the session that created this view.
	Session() Session

	// Name returns the name this view was constructed with.
	Name() string

	// Folder returns the root folder for this view.
	Folder() uri.URI

	// Called to set the effective contents of a file from this view.
	SetContent(ctx context.Context, uri uri.URI, content []byte)

	// Ignore returns true if this file should be ignored by this view.
	Ignore(uri.URI) bool

	// Shutdown closes this view, and detaches it from it's session.
	Shutdown(ctx context.Context) error

	// DidOpen is invoked each time a file is opened in the editor.
	DidOpen(uri uri.URI, text []byte)

	// DidSave is invoked each time an open file is saved in the editor.
	DidSave(uri uri.URI)

	// DidClose is invoked each time an open file is closed in the editor.
	DidClose(uri uri.URI)

	// IsOpen can be called to check if the editor has a file currently open.
	IsOpen(uri uri.URI) bool

	// FindFileByRelativePath gets File by relative filepath.
	// e.g.) returns File whose URI `file:///Users/username/proto/echo.proto` by filepath `proto/echo.proto`.
	FindFileByRelativePath(filepath string) (File, error)
}

type view struct {
	id      int64
	session Session

	// name is the user visible name of this view.
	name string

	// folder is the root of this view.
	folder uri.URI

	// keep track of files by uri and by basename, a single file may be mapped
	// to multiple uris, and the same basename may map to multiple files
	filesByURI  map[uri.URI]File
	filesByBase map[string][]File
	fileMu      *sync.RWMutex

	openFiles  map[uri.URI]bool
	openFileMu *sync.RWMutex

	// ignoredURIs is the set of URIs of files that we ignore.
	ignoredURIs  map[uri.URI]struct{}
	ignoredURIMu *sync.RWMutex
}

var _ View = (*view)(nil)

func NewView(session Session, name string, folder uri.URI) View {
	return &view{
		id:           viewIndex.Add(1),
		session:      session,
		name:         name,
		folder:       folder,
		filesByURI:   make(map[uri.URI]File),
		filesByBase:  make(map[string][]File),
		fileMu:       &sync.RWMutex{},
		openFiles:    make(map[uri.URI]bool),
		openFileMu:   &sync.RWMutex{},
		ignoredURIs:  make(map[uri.URI]struct{}),
		ignoredURIMu: &sync.RWMutex{},
	}
}

func (v *view) Session() Session {
	return v.session
}

func (v *view) Name() string {
	return v.name
}

func (v *view) Folder() uri.URI {
	return v.folder
}

func (v *view) GetFile(uri uri.URI) (File, error) {
	f, err := v.findFile(uri)
	if err != nil {
		return nil, err
	}
	if f != nil {
		return f, nil
	}

	file := &protoFile{
		File: &file{
			session: v.Session(),
			view:    v,
			uri:     uri,
		},
	}
	v.mapFile(uri, file)

	return file, nil
}

// SetContent sets the file contents for a file.
func (v *view) SetContent(ctx context.Context, uri uri.URI, data []byte) {
	if !v.Ignore(uri) {
		return
	}

	v.fileMu.Lock()
	defer v.fileMu.Unlock()

	if data == nil {
		delete(v.filesByURI, uri)
		return
	}

	pf := &protoFile{
		File: &file{
			session: v.Session(),
			uri:     uri,
			data:    data,
			hash:    hashContent(data),
		},
	}

	// TODO:
	//  Control times of parse of proto.
	//  Currently it parses every time of file change.
	pf.proto = parseProto(data)

	v.filesByURI[uri] = pf
}

func (v *view) Ignore(uri uri.URI) (ok bool) {
	v.ignoredURIMu.Lock()
	_, ok = v.ignoredURIs[uri]
	v.ignoredURIMu.Unlock()
	return
}

func (v *view) Shutdown(ctx context.Context) error {
	return v.session.RemoveView(ctx, v)
}

func (v *view) DidOpen(uri uri.URI, text []byte) {
	v.openFileMu.Lock()
	v.openFiles[uri] = true
	v.openFileMu.Unlock()
	v.openFile(uri, text)
}

func (v *view) DidSave(uri uri.URI) {
	v.fileMu.Lock()
	if file, ok := v.filesByURI[uri]; ok {
		file.SetSaved(true)
	}
	v.fileMu.Unlock()
}

func (v *view) DidClose(uri uri.URI) {
	v.openFileMu.Lock()
	delete(v.openFiles, uri)
	v.openFileMu.Unlock()
}

func (v *view) IsOpen(uri uri.URI) bool {
	v.openFileMu.RLock()
	defer v.openFileMu.RUnlock()

	open, ok := v.openFiles[uri]
	if !ok {
		return false
	}
	return open
}

func (v *view) FindFileByRelativePath(path string) (File, error) {
	fp := filepath.Join(v.Folder().Filename(), path)
	u := uri.File(fp)

	v.fileMu.RLock()
	f, ok := v.filesByURI[u]
	v.fileMu.RUnlock()
	if ok {
		return f, nil
	}

	// if view has not opened the file yet, view opens and returns the file.
	if err := v.openFileByFilepath(path); err != nil {
		return nil, err
	}
	v.fileMu.RLock()
	f = v.filesByURI[u]
	v.fileMu.RUnlock()

	return f, nil
}

func (v *view) openFile(uri uri.URI, data []byte) {
	pf := &protoFile{
		File: &file{
			session: v.Session(),
			view:    v,
			uri:     uri,
			data:    data,
			hash:    hashContent(data),
		},
	}

	pf.proto = parseProto(data)

	v.fileMu.Lock()
	v.filesByURI[uri] = pf
	v.fileMu.Unlock()

	v.openFileMu.Lock()
	v.openFiles[uri] = true
	v.openFileMu.Unlock()
}

func (v *view) openFileByFilepath(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	v.openFile(uri.File(path), data)

	return nil
}

func (v *view) findFile(uri uri.URI) (File, error) {
	v.fileMu.Lock()
	defer v.fileMu.Unlock()

	if f, ok := v.filesByURI[uri]; ok {
		return f, nil
	}

	filename := uri.Filename()
	basename := filepath.Base(filename)
	targetStat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, nil // the file may exist, return without an error
	}
	for _, f := range v.filesByBase[basename] {
		stat, err := os.Stat(f.URI().Filename())
		if err != nil {
			continue
		}
		if os.SameFile(targetStat, stat) {
			v.mapFile(uri, f)
			return f, nil
		}
	}
	return nil, nil
}

func (v *view) mapFile(uri uri.URI, f File) {
	v.fileMu.Lock()

	v.filesByURI[uri] = f
	basename := filepath.Base(uri.Filename())
	v.filesByBase[basename] = append(v.filesByBase[basename], f)

	v.fileMu.Unlock()
}

func parseProto(data []byte) registry.Proto {
	buf := bytes.NewBuffer(data)
	proto, err := parser.ParseProto(buf)
	if err != nil {
		return nil
	}
	return proto
}
