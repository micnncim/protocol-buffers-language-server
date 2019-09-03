package registry

import "github.com/emicklei/proto"

type Package struct {
	ProtoPackage *proto.Package
}

func NewPackage(protoPackage *proto.Package) *Package {
	return &Package{
		ProtoPackage: protoPackage,
	}
}
