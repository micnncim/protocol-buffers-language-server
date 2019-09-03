package registry

import (
	"sync"

	"github.com/emicklei/proto"
)

// ProtoSet is a registry for Proto.
type ProtoSet struct {
	Protos map[string]*Proto

	mu *sync.RWMutex
}

// NewProtoSet returns ProtoSet initialized by provided []*proto.Proto.
func NewProtoSet(protos ...*proto.Proto) *ProtoSet {
	protoSet := &ProtoSet{
		Protos: make(map[string]*Proto),
	}
	for _, p := range protos {
		protoSet.Protos[p.Filename] = NewProto(p)
	}
	return protoSet
}

// Append appends Proto to ProtoSet.
// This ensures thread safety.
func (p *ProtoSet) Append(proto *proto.Proto) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Protos[proto.Filename] = NewProto(proto)
}

// GetProtoByFilename gets Proto by provided Filename.
// This ensures thread safety.
func (p *ProtoSet) GetProtoByFilename(filename string) *Proto {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.Protos[filename]
}

// Proto is a registry for *proto.Proto.
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

	mu *sync.RWMutex
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
			pkg := NewPackage(v)
			p.PackageNameToPackage[v.Name] = pkg
			p.LineToPackage[v.Position.Line] = pkg

		case *proto.Message:
			m := NewMessage(v)
			p.MessageNameToMessage[v.Name] = m
			p.LineToMessage[v.Position.Line] = m

		case *proto.Enum:
			e := NewEnum(v)
			p.EnumNameToEnum[v.Name] = e
			p.LineToEnum[v.Position.Line] = e

		case *proto.Service:
			s := NewService(v)
			p.ServiceNameToService[v.Name] = s
			p.LineToService[v.Position.Line] = s

		default:

		}
	}

	return p
}

// GetPackageByName gets Package by provided name.
// This ensures thread safety.
func (p *Proto) GetPackageByName(name string) *Package {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.PackageNameToPackage[name]
}

// GetMessageByName gets Message by provided name.
// This ensures thread safety.
func (p *Proto) GetMessageByName(name string) *Message {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.MessageNameToMessage[name]
}

// GetEnumByName gets Enum by provided name.
// This ensures thread safety.
func (p *Proto) GetEnumByName(name string) *Enum {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.EnumNameToEnum[name]
}

// GetServiceByName gets Service by provided name.
// This ensures thread safety.
func (p *Proto) GetServiceByName(name string) *Service {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.ServiceNameToService[name]
}

// GetPackageByLine gets Package by provided line.
// This ensures thread safety.
func (p *Proto) GetPackageByLine(line int) *Package {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.LineToPackage[line]
}

// GetMessageByLine gets Message by provided line.
// This ensures thread safety.
func (p *Proto) GetMessageByLine(line int) *Message {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.LineToMessage[line]
}

// GetEnumByLine gets Enum by provided line.
// This ensures thread safety.
func (p *Proto) GetEnumByLine(line int) *Enum {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.LineToEnum[line]
}

// GetServiceByLine gets Service by provided line.
// This ensures thread safety.
func (p *Proto) GetServiceByLine(line int) *Service {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.LineToService[line]
}
