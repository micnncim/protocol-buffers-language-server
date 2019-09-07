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

	"go.uber.org/zap"

	"github.com/go-language-server/protocol"
)

// TODO: Match position with line and column.
// Currently matches with only line.
func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) { //nolint:unparam
	s.logger = s.logger.With(zap.Any("params", params))

	uri := params.TextDocument.URI

	view, ok := s.session.ViewOf(uri)
	if !ok {
		s.logger.Info("view not found", zap.String("filename", uri.Filename()))
		return
	}

	protoFile, ok := view.GetFile(uri)
	if !ok {
		s.logger.Info("file not found", zap.String("filename", uri.Filename()))
		return
	}

	proto, ok := protoFile.ProtoSet().GetProtoByFilename(uri.Filename())
	if !ok {
		s.logger.Info("file not found", zap.String("filename", uri.Filename()))
		return
	}

	field, ok := proto.GetMessageFieldByLine(int(params.Position.Line))
	if !ok {
		s.logger.Info("field not found")
		return
	}

	m, ok := proto.GetMessageByName(field.ProtoField.Type)
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
