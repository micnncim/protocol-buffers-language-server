package server

import (
	"context"

	"github.com/go-language-server/protocol"
)

func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	uri := params.TextDocument.URI

	p := s.protoSet.GetProtoByFilename(uri.Filename())

	f, ok := p.GetMessageFieldByLine(int(params.Position.Line))
	if !ok {
		s.logger.Info("definition not exist")
		return
	}
	t := f.ProtoField.Type
	m, ok := p.GetMessageByName(t)
	if !ok {
		return
	}

	// TODO:
	_ = m

	return
}
