package server

import (
	"context"

	"go.uber.org/zap"

	"github.com/go-language-server/protocol"
)

// TODO: Match position with line and column.
// Currently matches with only line.
func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	s.logger = s.logger.With(zap.Any("params", params))

	uri := params.TextDocument.URI

	p := s.protoSet.GetProtoByFilename(uri.Filename())

	f, ok := p.GetMessageFieldByLine(int(params.Position.Line))
	if !ok {
		s.logger.Info("field not found")
		return
	}
	t := f.ProtoField.Type
	m, ok := p.GetMessageByName(t)
	if !ok {
		s.logger.Info("message not found")
		return
	}
	line, column := m.Protobuf().Position.Line, m.Protobuf().Position.Column

	result = []protocol.Location{
		{
			URI: uri,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      float64(line),
					Character: float64(column),
				},
			},
		},
	}

	return
}
