package server

import (
	"context"

	"github.com/go-language-server/protocol"
)

func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	// uri := params.TextDocument.URI

	// p := s.protoSet.GetProtoByFilename(uri.Filename())

	// f, _ := p.GetFieldByLine(params.Position.Line)
	// t := f.Protobuf().ProtoField.Type
	// m, _ := p.GetMessageByName(t)

	return
}
