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
	"strings"

	"go.lsp.dev/protocol"
	"go.uber.org/zap"

	"github.com/micnncim/protocol-buffers-language-server/pkg/logging"
	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/types"
)

func (s *Server) completion(ctx context.Context, params *protocol.CompletionParams) (result *protocol.CompletionList, err error) {
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
		return
	}

	proto := protoFile.Proto()
	var items []protocol.CompletionItem

	// Get completions for field within messages.

	// TODO: Check whether the params.TextDocumentPositionParams.Position is valid.
	// TODO: Sort the items.

	// TODO: Remove this and judge whether target is rpc with better performance.
	buf, _, err := f.Read(ctx)
	if err != nil {
		logger.Error("failed to read file", zap.String("filename", filename))
		return
	}
	line := int(params.Position.Line) + 1
	targetLine := readLine(string(buf), line)
	// If target is not rpc, add build-in types to completion items.
	isRPC := strings.HasPrefix(strings.TrimSpace(targetLine), "rpc")

	if !isRPC {
		for _, t := range types.BuildInProtoTypes {
			items = append(items, protocol.CompletionItem{
				Label:  string(t),
				Detail: "type",
			})
		}
	}

	for _, m := range proto.Messages() {
		items = append(items, protocol.CompletionItem{
			Label:  m.Protobuf().Name,
			Detail: "message",
		})
	}

	if !isRPC {
		for _, e := range proto.Enums() {
			items = append(items, protocol.CompletionItem{
				Label:  e.Protobuf().Name,
				Detail: "enum",
			})
		}
	}

	result = &protocol.CompletionList{
		IsIncomplete: false,
		Items:        items,
	}
	return
}

func readLine(text string, line int) string {
	if line < 1 {
		return ""
	}
	slugs := strings.Split(text, "\n")
	return slugs[line-1]
}
