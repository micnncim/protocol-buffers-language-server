// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"github.com/go-language-server/uri"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

// ProtoFile represents a source file of any type.
type ProtoFile interface {
	URI() uri.URI
	View() View
	ProtoSet() registry.ProtoSet
}
