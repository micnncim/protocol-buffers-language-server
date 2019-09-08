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

package server

import (
	"context"
	"fmt"
	"net"
	"sync"

	"github.com/go-language-server/jsonrpc2"
	"github.com/go-language-server/protocol"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
)

type state int

const (
	stateCreated = state(iota)
	stateInitializing
	stateInitialized // Set once the server has received Initialize Request
	stateShutdown    // Set once the server has received Initialized Request
)

type Server struct {
	Conn   *jsonrpc2.Conn
	Client protocol.ClientInterface

	state   state
	stateMu *sync.RWMutex

	session source.Session

	logger *zap.Logger
}

var _ protocol.ServerInterface = (*Server)(nil)

type Option func(*Server)

func WithLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func New(ctx context.Context, session source.Session, stream jsonrpc2.Stream, opts ...Option) *Server {
	s := &Server{
		state:   stateCreated,
		stateMu: &sync.RWMutex{},
		session: session,
		logger:  zap.NewNop(),
	}
	for _, opt := range opts {
		opt(s)
	}

	jsonrpcOpts := []jsonrpc2.Options{
		jsonrpc2.WithCanceler(protocol.Canceller),
		jsonrpc2.WithCapacity(protocol.DefaultBufferSize),
		jsonrpc2.WithLogger(s.logger.Named("jsonrpc2")),
	}
	s.Conn, s.Client = protocol.NewServer(ctx, s, stream, zap.NewNop(), jsonrpcOpts...)
	s.logger = s.logger.Named("server")

	return s
}

// RunServerOnPort starts a server on the given port and does not exit.
// This function exists for debugging purposes.
func RunServerOnPort(ctx context.Context, session source.Session, port int, handler func(s *Server), opts ...Option) error {
	return RunServerOnAddress(ctx, session, fmt.Sprintf(":%v", port), handler, opts...)
}

// RunServerOnPort starts a server on the given port and does not exit.
// This function exists for debugging purposes.
func RunServerOnAddress(ctx context.Context, session source.Session, addr string, handler func(s *Server), opts ...Option) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		handler(New(ctx, session, jsonrpc2.NewStream(conn, conn), opts...))
	}
}

func (s *Server) Run(ctx context.Context) (err error) {
	return s.Conn.Run(ctx)
}

// Initialize implements initialize method.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	return s.initialize(ctx, params)
}

// Initialized implements initialized method.
// https://microsoft.github.io/language-server-protocol/specification#initialized
func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	return s.initialized(ctx, params)
}

// Shutdown implements shutdown method.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (s *Server) Shutdown(ctx context.Context) (err error) {
	return s.shutdown(ctx)
}

// Exit implements exit method.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (s *Server) Exit(ctx context.Context) (err error) {
	return s.exit(ctx)
}

func (s *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	err = notImplemented("CodeAction")
	return
}

func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	err = notImplemented("CodeLens")
	return
}

func (s *Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	err = notImplemented("CodeLensResolve")
	return
}

func (s *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	err = notImplemented("ColorPresentation")
	return
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	err = notImplemented("Completion")
	return
}

func (s *Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	err = notImplemented("CompletionResolve")
	return
}

func (s *Server) Declaration(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	err = notImplemented("Declaration")
	return
}

// Definition implements textDocument/definition method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_definition
func (s *Server) Definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	return s.definition(ctx, params)
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	err = notImplemented("DidChange")
	return
}

func (s *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	err = notImplemented("DidChangeConfiguration")
	return
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	err = notImplemented("DidChangeWatchedFiles")
	return
}

// DidChangeWorkspaceFolders implements workspace/didChangeWorkspaceFolders method.
// https://microsoft.github.io/language-server-protocol/specification#workspace_didChangeWorkspaceFolders
func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	return s.changeWorkspace(ctx, params.Event)
}

// DidClose implements textDocument/didClose method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didClose
func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	return s.didClose(ctx, params)
}

// DidOpen implements textDocument/didOpen method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didOpen
func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	return s.didOpen(ctx, params)
}

// DidSave implements textDocument/didSave method.
// https://microsoft.github.io/language-server-protocol/specification#textDocument_didSave
func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	return s.didSave(ctx, params)
}

func (s *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	err = notImplemented("DocumentColor")
	return
}

func (s *Server) DocumentHighlight(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.DocumentHighlight, err error) {
	err = notImplemented("DocumentHighlight")
	return
}

func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	err = notImplemented("DocumentLink")
	return
}

func (s *Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	err = notImplemented("DocumentLinkResolve")
	return
}

func (s *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []protocol.DocumentSymbol, err error) {
	err = notImplemented("DocumentSymbol")
	return
}

func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	err = notImplemented("ExecuteCommand")
	return
}

func (s *Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	err = notImplemented("FoldingRanges")
	return
}

func (s *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	err = notImplemented("Formatting")
	return
}

func (s *Server) Hover(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.Hover, err error) {
	err = notImplemented("Hover")
	return
}

func (s *Server) Implementation(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	err = notImplemented("Implementation")
	return
}

func (s *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	err = notImplemented("OnTypeFormatting")
	return
}

func (s *Server) PrepareRename(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.Range, err error) {
	err = notImplemented("PrepareRename")
	return
}

func (s *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	err = notImplemented("RangeFormatting")
	return
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	err = notImplemented("References")
	return
}

func (s *Server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	err = notImplemented("Rename")
	return
}

func (s *Server) SignatureHelp(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.SignatureHelp, err error) {
	err = notImplemented("SignatureHelp")
	return
}

func (s *Server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	err = notImplemented("Symbols")
	return
}

func (s *Server) TypeDefinition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	err = notImplemented("TypeDefinition")
	return
}

func (s *Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	err = notImplemented("WillSave")
	return
}

func (s *Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	err = notImplemented("WillSaveWaitUntil")
	return
}

func notImplemented(method string) error {
	return jsonrpc2.NewError(jsonrpc2.MethodNotFound, "method %q not implemented", method)
}
