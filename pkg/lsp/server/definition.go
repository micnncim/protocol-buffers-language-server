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
	"fmt"

	"github.com/go-language-server/protocol"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/logging"
	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// TODO: Match position with line and column.
// Currently matches with only line.
func (s *Server) definition(ctx context.Context, params *protocol.TextDocumentPositionParams) (result []protocol.Location, err error) {
	logger := logging.FromContext(ctx)
	logger = logger.With(zap.Any("params", params))

	uri := params.TextDocument.URI
	filename := uri.Filename()

	v := s.session.ViewOf(uri)

	f, err := v.GetFile(uri)
	if err != nil {
		logger.Error("file not found", zap.String("filename", filename))
		return
	}

	protoFile, ok := f.(source.ProtoFile)
	if !ok {
		logger.Warn("not proto file", zap.String("filename", filename))
		return
	}

	var protos []registry.Proto
	proto := protoFile.Proto()
	protos = append(protos, proto)

	line := int(params.Position.Line) + 1
	field, ok := proto.GetMessageFieldByLine(line)
	if !ok {
		logger.Warn("field not found", zap.Int("line", line))
		return
	}

	imports := proto.Imports()

	for _, imp := range imports {
		logger.Debug("import", zap.String("import", fmt.Sprintf("%#v", imp.ProtoImport.Filename)))

		f, err := v.FindFileByRelativePath(imp.ProtoImport.Filename)
		if err != nil {
			logger.Warn("failed to find file by import path", zap.Error(err))
			continue
		}
		pf, ok := f.(source.ProtoFile)
		if !ok {
			logger.Warn("not proto file", zap.String("filename", filename))
			continue
		}
		protos = append(protos, pf.Proto())
	}

	// TODO: Search the requested proto file and imported proto files.

	logger.Debug("protos", zap.String("protos", fmt.Sprintf("%#v", protos)))

	var m registry.Message
	var found bool
	typ := field.ProtoField.Type
	for _, p := range protos {
		m, found = p.GetMessageByName(typ)
		if found {
			break
		}
	}
	if !found {
		logger.Warn("message not found", zap.String("message_name", typ))
		return
	}

	line, column := m.Protobuf().Position.Line, m.Protobuf().Position.Column

	result = []protocol.Location{
		{
			URI: uri,
			Range: protocol.Range{
				Start: protocol.Position{
					Line:      float64(line) - 1,
					Character: float64(column) - 1,
				},
			},
		},
	}

	return
}
