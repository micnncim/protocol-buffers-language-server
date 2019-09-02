package parser

import (
	"os"

	"github.com/emicklei/proto"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

func ParseProtos(filenames ...string) (*registry.ProtoSet, error) {
	protoSet := registry.NewProtoSet()

	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		parser := proto.NewParser(f)
		p, err := parser.Parse()
		if err != nil {
			return nil, err
		}
		protoSet.Append(p)
	}

	return protoSet, nil
}
