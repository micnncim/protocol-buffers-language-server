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
	GetFile(uri uri.URI) (ProtoFile, bool)

	// Shutdown closes this view, and detaches it from it's session.
	Shutdown(ctx context.Context)
}

type view struct {
	id      int64
	session *session

	// name is the user visible name of this view.
	name string

	// folder is the root of this view.
	folder uri.URI

	// keep track of files by uri and by basename, a single file may be mapped
	// to multiple uris, and the same basename may map to multiple files
	uriToProtoFile      map[uri.URI]ProtoFile
	basenameToProtoFile map[string][]ProtoFile

	mu *sync.RWMutex
}

var _ View = (*view)(nil)

func (v *view) Session() Session {
	return v.session
}

func (v *view) Name() string {
	return v.name
}

func (v *view) Folder() uri.URI {
	return v.folder
}

func (v *view) GetFile(uri uri.URI) (ProtoFile, bool) {
	v.mu.RLock()
	defer v.mu.RUnlock()

	f, ok := v.uriToProtoFile[uri]
	return f, ok
}

// TODO: Implement.
func (v *view) Shutdown(ctx context.Context) {
	panic("implement me")
}
