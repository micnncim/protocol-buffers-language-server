package parser

import (
	"os"

	"github.com/emicklei/proto"
	"go.uber.org/multierr"

	"github.com/micnncim/protocol-buffers-language-server/pkg/proto/registry"
)

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
