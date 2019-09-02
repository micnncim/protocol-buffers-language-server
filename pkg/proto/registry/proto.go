package registry

import (
	"github.com/emicklei/proto"
)

type ProtoSet struct {
	protos map[string]*Proto
}

type Proto struct {
	protoProto *proto.Proto

	packageNameToPackage map[string]*Package
	messageNameToMessage map[string]*Message
	enumNameToEnum       map[string]*Enum
	serviceNameToService map[string]*Service

	lineToPackage map[int]*Package
	lineToMessage map[int]*Message
	lineToEnum    map[int]*Enum
	lineToService map[int]*Service
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
		protoProto: protoProto,

		packageNameToPackage: make(map[string]*Package),
		messageNameToMessage: make(map[string]*Message),
		enumNameToEnum:       make(map[string]*Enum),
		serviceNameToService: make(map[string]*Service),

		lineToPackage: make(map[int]*Package),
		lineToMessage: make(map[int]*Message),
		lineToEnum:    make(map[int]*Enum),
		lineToService: make(map[int]*Service),
	}

	for _, el := range protoProto.Elements {
		switch v := el.(type) {

		case *proto.Package:
			pkg := newPackage(v)
			p.packageNameToPackage[v.Name] = pkg
			p.lineToPackage[v.Position.Line] = pkg

		case *proto.Message:
			m := newMessage(v)
			p.messageNameToMessage[v.Name] = m
			p.lineToMessage[v.Position.Line] = m

		case *proto.Enum:
			e := newEnum(v)
			p.enumNameToEnum[v.Name] = e
			p.lineToEnum[v.Position.Line] = e

		case *proto.Service:
			s := newService(v)
			p.serviceNameToService[v.Name] = s
			p.lineToService[v.Position.Line] = s

		default:

		}
	}

	return p
}

func (p *Proto) GetPackageByLine(line int) *Package {
	return p.lineToPackage[line]
}

func (p *Proto) GetMessageByLine(line int) *Message {
	return p.lineToMessage[line]
}

func (p *Proto) GetEnumByLine(line int) *Enum {
	return p.lineToEnum[line]
}

func (p *Proto) GetServiceByLine(line int) *Service {
	return p.lineToService[line]
}
