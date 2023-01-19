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

// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"

	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
)

func (s *Server) changeWorkspace(ctx context.Context, event protocol.WorkspaceFoldersChangeEvent) error {
	for _, folder := range event.Removed {
		view, ok := s.session.View(folder.Name)
		if !ok {
			continue
		}
		if err := view.Shutdown(ctx); err != nil {
			return err
		}
	}

	for _, folder := range event.Added {
		s.addView(ctx, folder.Name, uri.File(folder.URI))
	}
	return nil
}

func (s *Server) addView(ctx context.Context, name string, uri uri.URI) {
	view := source.NewView(s.session, name, uri)
	s.session.AddView(ctx, view)
}
