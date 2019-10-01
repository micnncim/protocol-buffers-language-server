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
	"strings"

	"github.com/go-language-server/protocol"
	"github.com/go-language-server/uri"
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

	v := s.session.ViewOf(uri)

	f, err := v.GetFile(uri)
	if err != nil {
		logger.Error("file not found")
		return
	}

	protoFile, ok := f.(source.ProtoFile)
	if !ok {
		logger.Warn("not proto file")
		return
	}

	proto := protoFile.Proto()
	if len(proto.Packages()) == 0 {
		logger.Error("proto has no package")
		return
	}

	line := int(params.Position.Line) + 1
	field, ok := proto.GetMessageFieldByLine(line)
	if !ok {
		logger.Warn("field not found", zap.Int("line", line))
		return
	}
	typ := field.ProtoField.Type
	// Identify message's package.
	slugs := strings.Split(typ, ".")

	// In the case message's package is same as the package importing its package.

	if len(slugs) == 1 {
		m, ok := proto.GetMessageByName(typ)
		if !ok {
			logger.Warn("message not found", zap.String("message_name", typ))
			return
		}
		loc := messageToLocation(m, uri)
		result = []protocol.Location{loc}
		return
	}

	// In the case message's package is different from the package importing its package.

	protos := make(map[string]registry.Proto)
	imports := proto.Imports()
	for _, imp := range imports {
		f, err := v.FindFileByRelativePath(imp.ProtoImport.Filename)
		if err != nil {
			logger.Warn("failed to find file by import path", zap.Error(err))
			continue
		}
		pf, ok := f.(source.ProtoFile)
		if !ok {
			logger.Warn("not proto file")
			continue
		}
		if len(pf.Proto().Packages()) == 0 {
			logger.Warn("proto has no package")
			continue
		}
		protos[pf.Proto().Packages()[0].ProtoPackage.Name] = pf.Proto()
	}

	logger.Debug("finished import", zap.String("protos", fmt.Sprintf("%#v", protos)))

	typ = slugs[len(slugs)-1]
	pkg := strings.Join(slugs[0:len(slugs)-1], ".")
	logger.Debug("will get proto from map", zap.String("message_name", typ), zap.String("package", pkg))
	p, ok := protos[pkg]
	if !ok {
		logger.Error("proto not found", zap.String("package", pkg))
		return
	}
	logger.Debug("will get message from proto", zap.String("proto", p.Packages()[0].ProtoPackage.Name))
	m, ok := p.GetMessageByName(typ)
	if m == nil || !ok {
		logger.Warn("message not found", zap.String("message", typ))
		return
	}

	loc := messageToLocation(m, uri)
	result = []protocol.Location{loc}
	return
}

func messageToLocation(m registry.Message, uri uri.URI) protocol.Location {
	pb := m.Protobuf()
	line, column := pb.Position.Line, pb.Position.Column
	return protocol.Location{
		URI: uri,
		Range: protocol.Range{
			Start: protocol.Position{
				Line:      float64(line) - 1,
				Character: float64(column) - 1,
			},
		},
	}
}
