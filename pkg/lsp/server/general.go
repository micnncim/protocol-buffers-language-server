package server

import (
	"context"
	"os"

	"github.com/go-language-server/protocol"
	errors "golang.org/x/xerrors"
)

// Initialize resolves Initialize Request.
// https://microsoft.github.io/language-server-protocol/specification#initialize
func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (result *protocol.InitializeResult, err error) {
	s.stateMu.RLock()
	state := s.state
	s.stateMu.RUnlock()
	if state > stateInitializing {
		err = errors.Errorf("server already initialized")
		return
	}
	s.stateMu.Lock()
	s.state = stateInitializing
	s.stateMu.Unlock()

	result = &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: nil,
			HoverProvider:    false,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: []string{"."},
			},
			SignatureHelpProvider: &protocol.SignatureHelpOptions{
				TriggerCharacters: nil,
			},
			DefinitionProvider:              false,
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

// Initialized resolves Initialized Notification.
// https://microsoft.github.io/language-server-protocol/specification#initialized
func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) (err error) {
	s.stateMu.Lock()
	s.state = stateInitialized
	s.stateMu.Unlock()
	return
}

// Shutdown resolves Shutdown Request.
// https://microsoft.github.io/language-server-protocol/specification#shutdown
func (s *Server) Shutdown(ctx context.Context) (err error) {
	s.stateMu.RLock()
	state := s.state
	s.stateMu.RUnlock()
	if state < stateInitialized {
		err = errors.Errorf("server not initialized")
		return
	}
	s.stateMu.Lock()
	s.state = stateShutdown
	s.stateMu.Unlock()
	return
}

// Exit resolves Exit Notification.
// https://microsoft.github.io/language-server-protocol/specification#exit
func (s *Server) Exit(ctx context.Context) (err error) {
	s.stateMu.RLock()
	defer s.stateMu.RUnlock()
	if s.state != stateShutdown {
		os.Exit(1)
	}
	os.Exit(0)
	return
}
