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

package server

import (
	"context"

	"github.com/go-language-server/protocol"
)

func (s *Server) didOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	uri := params.TextDocument.URI
	s.session.DidOpen(ctx, uri)
	return
}

func (s *Server) didClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	uri := params.TextDocument.URI

	s.session.DidClose(uri)
	view, ok := s.session.ViewOf(uri)
	if !ok {
		return
	}
	if _, err = view.SetContent(ctx, uri, nil); err != nil {
		return
	}

	return
}

func (s *Server) didSave(_ context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	s.session.DidSave(params.TextDocument.URI)
	return
}
