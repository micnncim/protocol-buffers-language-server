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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"fmt"

	"github.com/go-language-server/jsonrpc2"
	"github.com/go-language-server/protocol"
)

func (s *Server) didOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	uri := params.TextDocument.URI
	text := []byte(params.TextDocument.Text)

	v := s.session.ViewOf(uri)
	v.DidOpen(uri, text)

	return nil
}

func (s *Server) didChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	if len(params.ContentChanges) < 1 {
		return jsonrpc2.NewError(jsonrpc2.InternalError, "no content changes provided")
	}

	uri := params.TextDocument.URI
	// TODO: Support incremental change. Currently support only full change.
	text := params.ContentChanges[0].Text

	switch s.config.TextDocumentSyncKind {
	case protocol.None:
		return nil
	case protocol.Full:
	case protocol.Incremental:
		return fmt.Errorf("incremental change is not supported yet")
	}

	v := s.session.ViewOf(uri)
	v.SetContent(ctx, uri, []byte(text))

	return nil
}

func (s *Server) didClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	uri := params.TextDocument.URI

	v := s.session.ViewOf(uri)
	v.DidClose(uri)
	v.SetContent(ctx, uri, nil)

	return nil
}

func (s *Server) didSave(_ context.Context, params *protocol.DidSaveTextDocumentParams) error {
	uri := params.TextDocument.URI

	v := s.session.ViewOf(uri)
	v.DidSave(uri)

	return nil
}
