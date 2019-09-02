package server

import (
	"context"

	"github.com/go-language-server/protocol"
)

type Message struct {
}

func (s *Server) Definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	// FIXME: Fix protoSet appropriately.

	return
}
