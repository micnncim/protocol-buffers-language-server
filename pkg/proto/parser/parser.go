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

package parser

import (
	"os"

	"github.com/emicklei/proto"
	"go.uber.org/multierr"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// ParseProtos parses protobuf files from filenames and return registry.ProtoSet.
func ParseProtos(filenames ...string) (registry.ProtoSet, error) {
	protoSet := registry.NewProtoSet()

	var errs error
	for _, filename := range filenames {
		func() {
			f, err := os.Open(filename)
			if err != nil {
				multierr.Append(errs, err)
			}
			defer f.Close()
			parser := proto.NewParser(f)
			p, err := parser.Parse()
			if err != nil {
				multierr.Append(errs, err)
			}
			protoSet.Append(p)
		}()
	}

	return protoSet, errs
}
