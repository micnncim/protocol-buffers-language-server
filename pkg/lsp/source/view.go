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
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-language-server/uri"
)

// View represents a single workspace.
// This is the level at which we maintain configuration like working directory
// and build tags.
type View interface {
	// Session returns the session that created this view.
	Session() Session

	// Name returns the name this view was constructed with.
	Name() string

	// Folder returns the root folder for this view.
	Folder() uri.URI

	// GetFile returns the file object for a given uri.
	GetFile(uri uri.URI) (File, error)

	// Called to set the effective contents of a file from this view.
	SetContent(ctx context.Context, uri uri.URI, content []byte) (wasFirstChange bool, err error)

	// Ignore returns true if this file should be ignored by this view.
	Ignore(uri.URI) bool

	// Shutdown closes this view, and detaches it from it's session.
	Shutdown(ctx context.Context) error
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

	// ignoredURIs is the set of URIs of files that we ignore.
	ignoredURIsMu *sync.RWMutex
	ignoredURIs   map[uri.URI]struct{}

	mu *sync.RWMutex
}

var _ View = (*view)(nil)

func NewView(session Session, name string, folder uri.URI) View {
	return &view{
		id:            viewIndex.Add(1),
		session:       session,
		name:          name,
		folder:        folder,
		filesByURI:    make(map[uri.URI]File),
		filesByBase:   make(map[string][]File),
		ignoredURIsMu: nil,
		ignoredURIs:   nil,
		mu:            &sync.RWMutex{},
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
		fileBase: fileBase{
			uri:  uri,
			view: v,
		},
	}
	v.mapFile(uri, file)

	return file, nil
}

// SetContent sets the Overlay contents for a file.
func (v *view) SetContent(ctx context.Context, uri uri.URI, content []byte) (bool, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if !v.Ignore(uri) {
		return v.session.SetOverlay(uri, content), nil
	}
	return false, nil
}

func (v *view) Ignore(uri uri.URI) (ok bool) {
	v.ignoredURIsMu.Lock()
	_, ok = v.ignoredURIs[uri]
	v.ignoredURIsMu.Unlock()
	return
}

func (v *view) Shutdown(ctx context.Context) error {
	return v.session.RemoveView(ctx, v)
}

func (v *view) findFile(uri uri.URI) (File, error) {
	v.mu.Lock()
	defer v.mu.Unlock()

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
	v.mu.Lock()
	v.filesByURI[uri] = f
	basename := filepath.Base(uri.Filename())
	v.filesByBase[basename] = append(v.filesByBase[basename], f)
	v.mu.Unlock()
}
