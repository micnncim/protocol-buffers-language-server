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

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// File represents a source file of any type.
type File interface {
	URI() uri.URI
	View() View
	Handle(ctx context.Context) FileHandle
	// TODO: Fix appropriately.
	Proto() registry.Proto
}

// FileHandle represents a handle to a specific version of a single file from
// a specific file system.
type FileHandle interface {
	// FileSystem returns the file system this handle was acquired from.
	FileSystem() FileSystem

	// Read reads the contents of a file and returns it along with its hash value.
	// If the file is not available, returns a nil slice and an error.
	Read(ctx context.Context) ([]byte, string, error)
}

// FileSystem is the interface to something that provides file contents.
type FileSystem interface {
	// GetFile returns a handle for the specified file.
	GetFile(uri uri.URI) FileHandle
}

// fileBase implements File.
// fileBase holds the common functionality for all files.
// It is intended to be embedded in the file implementations.
type fileBase struct {
	uri uri.URI

	view View

	// TODO: Fix appropriately.
	proto registry.Proto

	handleMu *sync.RWMutex
	handle   FileHandle
}

var _ File = (*fileBase)(nil)

func (f *fileBase) URI() uri.URI {
	return f.uri
}

func (f *fileBase) View() View {
	return f.view
}

func (f *fileBase) Handle(ctx context.Context) FileHandle {
	f.handleMu.Lock()
	defer f.handleMu.Unlock()
	if f.handle == nil {
		f.handle = f.view.Session().GetFile(f.URI())
	}
	return f.handle
}

func (f *fileBase) Proto() registry.Proto {
	return f.proto
}
