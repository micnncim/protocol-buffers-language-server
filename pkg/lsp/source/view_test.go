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

package source

import (
	"context"
	"path/filepath"
	"sync"
	"testing"

	"github.com/go-language-server/uri"
	"github.com/stretchr/testify/assert"

	"github.com/micnncim/protocol-buffers-language-server/pkg/internal/testutil"
)

func TestFindFileByRelativePath(t *testing.T) {
	wd := testutil.Getwd(t)

	cases := []struct {
		name       string
		filesByURI map[uri.URI]File
		path       string
		want       File
		wantErr    bool
	}{
		{
			name: "found from view",
			filesByURI: map[uri.URI]File{
				uri.File(filepath.Join(wd, "proto/user/v1/user.proto")): &protoFile{
					File: &file{
						uri: uri.File(filepath.Join(wd, "proto/user/v1/user.proto")),
					},
				},
			},
			path: "proto/user/v1/user.proto",
			want: &protoFile{
				File: &file{
					uri: uri.File(filepath.Join(wd, "proto/user/v1/user.proto")),
				},
			},
			wantErr: false,
		},
		{
			name:       "not found from view but newly open the file",
			filesByURI: make(map[uri.URI]File),
			path:       "testdata/test.proto",
			want: func() ProtoFile {
				data := testutil.ReadFile(t, "testdata/test.proto")
				return &protoFile{
					File: &file{
						session: &session{},
						uri:     uri.File(filepath.Join(wd, "testdata/test.proto")),
						data:    data,
					},
					proto: parseProto(data),
				}
			}(),
			wantErr: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			v := &view{
				session:    &session{},
				folder:     uri.File(wd),
				filesByURI: tc.filesByURI,
				fileMu:     &sync.RWMutex{},
				openFiles:  make(map[uri.URI]bool),
				openFileMu: &sync.RWMutex{},
			}

			got, err := v.FindFileByRelativePath(tc.path)

			assert.Equal(t, tc.wantErr, err != nil)

			ctx := context.Background()
			wantData, _, err := tc.want.Read(ctx)
			if err != nil {
				t.Fatal(err)
			}
			gotData, _, err := got.Read(ctx)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, wantData, gotData)
			assert.Equal(t, tc.want.URI(), got.URI())
		})
	}
}
