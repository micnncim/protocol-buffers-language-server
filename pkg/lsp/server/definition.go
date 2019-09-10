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

package server

import (
	"context"

	"github.com/go-language-server/protocol"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/logging"
	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
)

// TODO: Match position with line and column.
// Currently matches with only line.
func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	logger := logging.FromContext(ctx)
	logger = logger.With(zap.Any("params", params))

	uri := params.TextDocument.URI
	filename := uri.Filename()

	view, ok := s.session.ViewOf(uri)
	if !ok {
		logger.Warn("view not found", zap.String("filename", filename))
		return
	}

	f, err := view.GetFile(uri)
	if err != nil {
		logger.Error("file not found", zap.String("filename", filename))
		return
	}

	protoFile, ok := f.(source.ProtoFile)
	if !ok {
		return
	}

	proto := protoFile.Proto()

	line := int(params.Position.Line)
	field, ok := proto.GetMessageFieldByLine(line)
	if !ok {
		logger.Warn("field not found", zap.Int("line", line))
		return
	}

	typ := field.ProtoField.Type
	// TODO: Search the requested proto file and imported proto files.
	m, ok := proto.GetMessageByName(typ)
	if !ok {
		logger.Warn("message not found", zap.String("name", typ))
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
