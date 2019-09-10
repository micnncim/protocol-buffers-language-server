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
	Version() string
}

// ProtoFile represents a source file of protobuf.
type ProtoFile interface {
	File
	Proto() registry.Proto
}

// FileReader reads file's content and returns the data and hash.
type FileReader interface {
	Read(ctx context.Context) (data []byte, hash string, err error)
}

// FileHandle represents a handle to a specific version of a single file from
// a specific file system.
type FileHandle interface {
	FileReader
	File() File
	FileSystem() FileSystem
}

// FileHandle represents a handle to a specific version of a single protobuf file from
// a specific file system.
type ProtoFileHandle interface {
	FileHandle
	ProtoFile() ProtoFile
}

// FileSystem is the interface to something that provides file contents.
type FileSystem interface {
	// GetFile returns a handle for the specified file.
	GetFile(uri uri.URI) FileHandle
}

// ProtoFileSystem is the interface to something that provides protobuf file contents.
type ProtoFileSystem interface {
	// GetProtoFile returns a handle for the specified file.
	GetProtoFile(uri uri.URI) ProtoFileHandle
}

type protoFile struct {
	fileBase

	registry.Proto
}

type fileBase struct {
	uri      uri.URI
	filename string

	view View

	handleMu *sync.RWMutex
	handle   FileHandle
}

func (f *fileBase) URI() uri.URI {
	return f.uri
}

func (f fileBase) View() View {
	return f.view
}

func (f *fileBase) Handle(ctx context.Context) FileHandle {
	if f.handle == nil {
		f.handleMu.Lock()
		f.handle = f.view.Session().GetFile(f.URI())
		f.handleMu.Unlock()
	}
	return f.handle
}
