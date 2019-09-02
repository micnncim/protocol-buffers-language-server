package server

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/go-language-server/jsonrpc2"
	"github.com/go-language-server/protocol"
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

	logger *zap.Logger
}

var _ protocol.ServerInterface = (*Server)(nil)

type Option func(*Server)

func WithLogger(logger *zap.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func NewServer(ctx context.Context, stream jsonrpc2.Stream, opts ...Option) *Server {
	s := &Server{
		state:   stateCreated,
		stateMu: &sync.RWMutex{},
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

func (s *Server) Run(ctx context.Context) (err error) {
	return s.Conn.Run(ctx)
}

func (s *Server) CodeAction(ctx context.Context, params *protocol.CodeActionParams) (result []protocol.CodeAction, err error) {
	panic("not implement yet")
}

func (s *Server) CodeLens(ctx context.Context, params *protocol.CodeLensParams) (result []protocol.CodeLens, err error) {
	panic("not implement yet")
}

func (s *Server) CodeLensResolve(ctx context.Context, params *protocol.CodeLens) (result *protocol.CodeLens, err error) {
	panic("not implement yet")
}

func (s *Server) ColorPresentation(ctx context.Context, params *protocol.ColorPresentationParams) (result []protocol.ColorPresentation, err error) {
	panic("not implement yet")
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
	panic("not implement yet")
}

func (s *Server) CompletionResolve(ctx context.Context, params *protocol.CompletionItem) (result *protocol.CompletionItem, err error) {
	panic("not implement yet")
}

func (s *Server) Declaration(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	panic("not implement yet")
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidChangeConfiguration(ctx context.Context, params *protocol.DidChangeConfigurationParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidChangeWorkspaceFolders(ctx context.Context, params *protocol.DidChangeWorkspaceFoldersParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) (err error) {
	panic("not implement yet")
}

func (s *Server) DocumentColor(ctx context.Context, params *protocol.DocumentColorParams) (result []protocol.ColorInformation, err error) {
	panic("not implement yet")
}

func (s *Server) DocumentHighlight(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.DocumentHighlight, err error) {
	panic("not implement yet")
}

func (s *Server) DocumentLink(ctx context.Context, params *protocol.DocumentLinkParams) (result []protocol.DocumentLink, err error) {
	panic("not implement yet")
}

func (s *Server) DocumentLinkResolve(ctx context.Context, params *protocol.DocumentLink) (result *protocol.DocumentLink, err error) {
	panic("not implement yet")
}

func (s *Server) DocumentSymbol(ctx context.Context, params *protocol.DocumentSymbolParams) (result []protocol.DocumentSymbol, err error) {
	panic("not implement yet")
}

func (s *Server) ExecuteCommand(ctx context.Context, params *protocol.ExecuteCommandParams) (result interface{}, err error) {
	panic("not implement yet")
}

func (s *Server) FoldingRanges(ctx context.Context, params *protocol.FoldingRangeParams) (result []protocol.FoldingRange, err error) {
	panic("not implement yet")
}

func (s *Server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implement yet")
}

func (s *Server) Hover(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.Hover, err error) {
	panic("not implement yet")
}

func (s *Server) Implementation(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	panic("not implement yet")
}

func (s *Server) OnTypeFormatting(ctx context.Context, params *protocol.DocumentOnTypeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implement yet")
}

func (s *Server) PrepareRename(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.Range, err error) {
	panic("not implement yet")
}

func (s *Server) RangeFormatting(ctx context.Context, params *protocol.DocumentRangeFormattingParams) (result []protocol.TextEdit, err error) {
	panic("not implement yet")
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) (result []protocol.Location, err error) {
	panic("not implement yet")
}

func (s *Server) Rename(ctx context.Context, params *protocol.RenameParams) (result *protocol.WorkspaceEdit, err error) {
	panic("not implement yet")
}

func (s *Server) SignatureHelp(ctx context.Context, params *protocol.TextDocumentPositionParams) (result *protocol.SignatureHelp, err error) {
	panic("not implement yet")
}

func (s *Server) Symbols(ctx context.Context, params *protocol.WorkspaceSymbolParams) (result []protocol.SymbolInformation, err error) {
	panic("not implement yet")
}

func (s *Server) TypeDefinition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	panic("not implement yet")
}

func (s *Server) WillSave(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (err error) {
	panic("not implement yet")
}

func (s *Server) WillSaveWaitUntil(ctx context.Context, params *protocol.WillSaveTextDocumentParams) (result []protocol.TextEdit, err error) {
	panic("not implement yet")
}
