package registry

import "github.com/emicklei/proto"

type Package struct {
	protoPackage *proto.Package
}

func newPackage(protoPackage *proto.Package) *Package {
	return &Package{
		protoPackage: protoPackage,
	}
}
