package server

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-language-server/protocol"
)

func TestServerDefinition(t *testing.T) {
	cases := []struct {
		name       string
		params     *protocol.TextDocumentPositionParams
		wantResult []protocol.Location
		wantErr    bool
	}{
		{},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			s := NewServer(ctx, nil)

			result, err := s.Definition(ctx, tc.params)

			assert.Equal(t, tc.wantResult, result)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}
