package server

import (
	"context"
	"text/scanner"

	"github.com/emicklei/proto"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"

	"github.com/go-language-server/protocol"
)

type Message struct {
}

func (s *Server) Definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	// FIXME: Fix protoSet appropriately.

	p := &registry.Proto{
		ProtoProto: &proto.Proto{
			Filename: params.TextDocument.URI.Filename(),
		},
		MessageNameToMessage: map[string]*registry.Message{
			"Message": {
				ProtoMessage: &proto.Message{
					Position: scanner.Position{
						Line:   16,
						Column: 9,
					},
					Name: "Message",
				},
			},
		},
		LineToMessage: map[int]*registry.Message{
			11: {
				LineToField: map[int]*registry.MessageField{
					13: {
						ProtoField: &proto.NormalField{
							Field: &proto.Field{
								Position: scanner.Position{
									Line: 13,
								},
								Name:     "message",
								Type:     "Message",
								Sequence: 1,
							},
						},
					},
				},
			},
		},
	}

	field, ok := p.LineToMessage[11].LineToField[int(params.Position.Line)]
	if !ok {
		return
	}
	message := p.MessageNameToMessage[field.ProtoField.Type]

	result = []protocol.Location{
		{
			URI: params.TextDocument.URI,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      float64(message.ProtoMessage.Position.Line),
					Character: float64(message.ProtoMessage.Position.Column),
				},
			},
		},
	}

	return
}
