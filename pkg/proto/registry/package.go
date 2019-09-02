package registry

import "github.com/emicklei/proto"

type PackageSet struct {
	packageNameToPackage map[string]*Package
}

type Package struct {
	protoPackage *proto.Package

	messageNameToMessage map[string]*Message
	enumNameToEnum       map[string]*Enum
	serviceNameToService map[string]*Service
}

func NewPackage(protoPackage *proto.Package) *Package {
	return &Package{
		protoPackage:         protoPackage,
		messageNameToMessage: make(map[string]*Message),
		enumNameToEnum:       make(map[string]*Enum),
		serviceNameToService: make(map[string]*Service),
	}
}
