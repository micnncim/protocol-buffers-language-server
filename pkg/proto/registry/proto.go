package registry

import (
	"github.com/emicklei/proto"
)

type ProtoSet struct {
	protos map[string]*Proto
}

type Proto struct {
	ProtoProto *proto.Proto

	PackageNameToPackage map[string]*Package
	MessageNameToMessage map[string]*Message
	EnumNameToEnum       map[string]*Enum
	ServiceNameToService map[string]*Service

	LineToPackage map[int]*Package
	LineToMessage map[int]*Message
	LineToEnum    map[int]*Enum
	LineToService map[int]*Service
}

func NewProtoSet(protos ...*proto.Proto) *ProtoSet {
	protoSet := &ProtoSet{
		protos: make(map[string]*Proto),
	}
	for _, p := range protos {
		protoSet.protos[p.Filename] = NewProto(p)
	}
	return protoSet
}

func (p *ProtoSet) Append(proto *proto.Proto) {
	p.protos[proto.Filename] = NewProto(proto)
}

func NewProto(protoProto *proto.Proto) *Proto {
	p := &Proto{
		ProtoProto: protoProto,

		PackageNameToPackage: make(map[string]*Package),
		MessageNameToMessage: make(map[string]*Message),
		EnumNameToEnum:       make(map[string]*Enum),
		ServiceNameToService: make(map[string]*Service),

		LineToPackage: make(map[int]*Package),
		LineToMessage: make(map[int]*Message),
		LineToEnum:    make(map[int]*Enum),
		LineToService: make(map[int]*Service),
	}

	for _, el := range protoProto.Elements {
		switch v := el.(type) {

		case *proto.Package:
			pkg := newPackage(v)
			p.PackageNameToPackage[v.Name] = pkg
			p.LineToPackage[v.Position.Line] = pkg

		case *proto.Message:
			m := newMessage(v)
			p.MessageNameToMessage[v.Name] = m
			p.LineToMessage[v.Position.Line] = m

		case *proto.Enum:
			e := newEnum(v)
			p.EnumNameToEnum[v.Name] = e
			p.LineToEnum[v.Position.Line] = e

		case *proto.Service:
			s := newService(v)
			p.ServiceNameToService[v.Name] = s
			p.LineToService[v.Position.Line] = s

		default:

		}
	}

	return p
}

func (p *Proto) GetPackageByLine(line int) *Package {
	return p.LineToPackage[line]
}

func (p *Proto) GetMessageByLine(line int) *Message {
	return p.LineToMessage[line]
}

func (p *Proto) GetEnumByLine(line int) *Enum {
	return p.LineToEnum[line]
}

func (p *Proto) GetServiceByLine(line int) *Service {
	return p.LineToService[line]
}
