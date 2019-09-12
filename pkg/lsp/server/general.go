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
	"os"
	"path/filepath"

	"github.com/go-language-server/jsonrpc2"
	"github.com/go-language-server/protocol"
	"github.com/go-language-server/uri"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/config"
	"github.com/micnncim/protocol-buffers-language-server/pkg/logging"
)

func (s *Server) initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	logger := logging.FromContext(ctx)
	logger = logger.With(zap.Any("params", params))
	logger.Debug("start initialize")
	defer logger.Debug("end initialize")

	s.stateMu.RLock()
	state := s.state
	s.stateMu.RUnlock()
	if state > stateInitializing {
		err = jsonrpc2.NewError(jsonrpc2.InvalidRequest, "server already initialized")
		return
	}
	s.stateMu.Lock()
	s.state = stateInitializing
	s.stateMu.Unlock()

	folders := params.WorkspaceFolders
	if len(folders) == 0 {
		rootURI := params.RootURI
		if rootURI == "" {
			err = errors.New("single file mode not supported yet")
			return
		}
		folders = []protocol.WorkspaceFolder{
			{
				URI:  rootURI.Filename(),
				Name: filepath.Base(rootURI.Filename()),
			},
		}
	}

	for _, folder := range folders {
		s.addView(ctx, folder.Name, uri.File(folder.URI))
	}

	cfg := config.DefaultLSPConfig

	result = &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    float64(cfg.TextDocumentSyncKind),
			},
			HoverProvider: false,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: []string{"."},
			},
			SignatureHelpProvider: &protocol.SignatureHelpOptions{
				TriggerCharacters: nil,
			},
			DefinitionProvider:              true,
			WorkspaceSymbolProvider:         false,
			DocumentFormattingProvider:      false,
			DocumentRangeFormattingProvider: false,
			RenameProvider:                  nil,
			FoldingRangeProvider:            nil,
			Workspace: &protocol.ServerCapabilitiesWorkspace{
				WorkspaceFolders: &protocol.ServerCapabilitiesWorkspaceFolders{
					Supported:           false,
					ChangeNotifications: nil,
				},
			},
		},
	}

	return
}

func (s *Server) initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	logger := logging.FromContext(ctx)
	logger = logger.With(zap.Any("params", params))
	logger.Debug("start initialized")
	defer logger.Debug("end initialized")

	s.stateMu.Lock()
	s.state = stateInitialized
	s.stateMu.Unlock()
	return
}

func (s *Server) shutdown(ctx context.Context) (err error) {
	logger := logging.FromContext(ctx)
	logger.Debug("start shutdown")
	defer logger.Debug("end shutdown")

	s.stateMu.RLock()
	state := s.state
	s.stateMu.RUnlock()
	if state < stateInitialized {
		err = jsonrpc2.NewError(jsonrpc2.InvalidRequest, "server not initialized")
		return
	}
	s.session.Shutdown(ctx)
	s.stateMu.Lock()
	s.state = stateShutdown
	s.stateMu.Unlock()
	return
}

func (s *Server) exit(ctx context.Context) (err error) {
	logger := logging.FromContext(ctx)
	logger.Debug("start exit")
	defer logger.Debug("end exit")

	s.stateMu.RLock()
	defer s.stateMu.RUnlock()
	if s.state != stateShutdown {
		os.Exit(1)
	}
	os.Exit(0)
	return
}
