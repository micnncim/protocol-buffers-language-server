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

	"github.com/go-language-server/uri"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// File represents a source file of any type.
type File interface {
	URI() uri.URI
	View() View
	FileSystem() FileSystem
	Read(ctx context.Context) ([]byte, string, error)

	Saved() bool
	// TODO: Fix appropriate function name.
	SetSaved(saved bool)
}

type ProtoFile interface {
	File
	Proto() registry.Proto
	SetProto(proto registry.Proto)
}

// FileSystem is the interface to something that provides file contents.
type FileSystem interface {
	// GetFile returns a file whose the given uri.
	GetFile(uri uri.URI) (File, error)
}

// file is a file for changed files.
type file struct {
	session Session
	view    View

	uri  uri.URI
	data []byte
	hash string

	// saved is true if a file has been saved on disk.
	saved bool
}

var _ File = (*file)(nil)

type protoFile struct {
	File
	proto registry.Proto
}

var _ ProtoFile = (*protoFile)(nil)

func (f *file) URI() uri.URI {
	return f.uri
}

func (f *file) View() View {
	return f.view
}

func (f *file) FileSystem() FileSystem {
	return f.view
}

func (f *file) Read(context.Context) ([]byte, string, error) {
	return f.data, f.hash, nil
}

func (f *file) Saved() bool {
	return f.saved
}

func (f *file) SetSaved(saved bool) {
	f.saved = saved
}

func (p *protoFile) Proto() registry.Proto {
	return p.proto
}

func (p *protoFile) SetProto(proto registry.Proto) {
	p.proto = proto
}
