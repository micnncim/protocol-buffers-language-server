package server

import (
	"context"
	"testing"

	"github.com/go-language-server/uri"

	"github.com/go-language-server/protocol"
)

func TestServerDefinition(t *testing.T) {
	ctx := context.Background()
	s := NewServer(ctx, nil)

	result, err := s.Definition(ctx, &protocol.TextDocumentPositionParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: uri.File("proto/echo.proto"),
		},
		Position: protocol.Position{
			Line:      13,
			Character: 14,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
