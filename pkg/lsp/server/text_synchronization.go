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
	"errors"
	"fmt"

	"github.com/go-language-server/jsonrpc2"
	"github.com/go-language-server/protocol"
)

func (s *Server) didOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	uri := params.TextDocument.URI
	s.session.DidOpen(ctx, uri)
	return nil
}

func (s *Server) didChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	if len(params.ContentChanges) < 1 {
		return jsonrpc2.NewError(jsonrpc2.InternalError, "no content changes provided")
	}

	uri := params.TextDocument.URI

	text, isFullChanged := getChangedText(params.ContentChanges)

	// TODO: Implement logic when isFullChanged
	if isFullChanged {
		switch s.config.TextDocumentSyncKind {
		case protocol.None:
		case protocol.Full:
		case protocol.Incremental:
		}
	}

	view, ok := s.session.ViewOf(uri)
	if !ok {
		return errors.New("view not found")
	}
	if _, err := view.SetContent(ctx, uri, []byte(text)); err != nil {
		return err
	}

	return nil
}

func (s *Server) didClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	uri := params.TextDocument.URI

	s.session.DidClose(uri)
	view, ok := s.session.ViewOf(uri)
	if !ok {
		return fmt.Errorf("view of %s not found", uri.Filename())
	}
	if _, err := view.SetContent(ctx, uri, nil); err != nil {
		return err
	}

	return nil
}

func (s *Server) didSave(_ context.Context, params *protocol.DidSaveTextDocumentParams) error {
	s.session.DidSave(params.TextDocument.URI)
	return nil
}

func getChangedText(changes []protocol.TextDocumentContentChangeEvent) (text string, isFullChanged bool) {
	if len(changes) > 1 {
		return
	}
	// The length of the changes must be 1 at this point.
	if changes[0].Range == nil && changes[0].RangeLength == 0 {
		text, isFullChanged = changes[0].Text, true
		return
	}
	return
}
