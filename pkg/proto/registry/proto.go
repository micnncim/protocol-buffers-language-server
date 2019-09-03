package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

// ProtoSet is a registry for Proto.
type ProtoSet interface {
	Protos() []Proto
	Append(proto *protobuf.Proto)

	GetProtoByFilename(filename string) Proto
}

type protoSet struct {
	protos map[string]Proto

	mu *sync.RWMutex
}

var _ ProtoSet = (*protoSet)(nil)

// NewProtoSet returns protoSet initialized by provided []*protobuf.Proto.
func NewProtoSet(protos ...*protobuf.Proto) ProtoSet {
	protoSet := &protoSet{
		protos: make(map[string]Proto),
	}
	for _, p := range protos {
		protoSet.protos[p.Filename] = NewProto(p)
	}
	return protoSet
}

func (p *protoSet) Protos() []Proto {
	p.mu.Lock()
	defer p.mu.Unlock()

	protos := make([]Proto, 0, len(p.protos))
	for _, proto := range p.protos {
		protos = append(protos, proto)
	}
	return protos
}

// Append appends Proto to protoSet.
// This ensures thread safety.
func (p *protoSet) Append(proto *protobuf.Proto) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.protos[proto.Filename] = NewProto(proto)
}

// GetProtoByFilename gets Proto by provided Filename.
// This ensures thread safety.
func (p *protoSet) GetProtoByFilename(filename string) Proto {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.protos[filename]
}

// Proto is a registry for protobuf proto.
type Proto interface {
	Protobuf() *protobuf.Proto

	GetPackageByName(name string) *Package
	GetMessageByName(name string) Message
	GetEnumByName(name string) Enum
	GetServiceByName(name string) Service

	GetPackageByLine(line int) *Package
	GetMessageByLine(line int) Message
	GetEnumByLine(line int) Enum
	GetServiceByLine(line int) Service
}

type proto struct {
	protoProto *protobuf.Proto

	packageNameToPackage map[string]*Package
	messageNameToMessage map[string]Message
	enumNameToEnum       map[string]Enum
	serviceNameToService map[string]Service

	lineToPackage map[int]*Package
	lineToMessage map[int]Message
	lineToEnum    map[int]Enum
	lineToService map[int]Service

	mu *sync.RWMutex
}

var _ Proto = (*proto)(nil)

func NewProto(protoProto *protobuf.Proto) Proto {
	proto := &proto{
		protoProto: protoProto,

		packageNameToPackage: make(map[string]*Package),
		messageNameToMessage: make(map[string]Message),
		enumNameToEnum:       make(map[string]Enum),
		serviceNameToService: make(map[string]Service),

		lineToPackage: make(map[int]*Package),
		lineToMessage: make(map[int]Message),
		lineToEnum:    make(map[int]Enum),
		lineToService: make(map[int]Service),
	}

	for _, el := range protoProto.Elements {
		switch v := el.(type) {

		case *protobuf.Package:
			p := NewPackage(v)
			proto.packageNameToPackage[v.Name] = p
			proto.lineToPackage[v.Position.Line] = p

		case *protobuf.Message:
			m := NewMessage(v)
			proto.messageNameToMessage[v.Name] = m
			proto.lineToMessage[v.Position.Line] = m

		case *protobuf.Enum:
			e := NewEnum(v)
			proto.enumNameToEnum[v.Name] = e
			proto.lineToEnum[v.Position.Line] = e

		case *protobuf.Service:
			s := NewService(v)
			proto.serviceNameToService[v.Name] = s
			proto.lineToService[v.Position.Line] = s

		default:

		}
	}

	return proto
}

// Protobuf returns *protobuf.Proto.
func (p *proto) Protobuf() *protobuf.Proto {
	return p.protoProto
}

// GetPackageByName gets Package by provided name.
// This ensures thread safety.
func (p *proto) GetPackageByName(name string) *Package {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.packageNameToPackage[name]
}

// GetMessageByName gets message by provided name.
// This ensures thread safety.
func (p *proto) GetMessageByName(name string) Message {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.messageNameToMessage[name]
}

// GetEnumByName gets enum by provided name.
// This ensures thread safety.
func (p *proto) GetEnumByName(name string) Enum {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.enumNameToEnum[name]
}

// GetServiceByName gets service by provided name.
// This ensures thread safety.
func (p *proto) GetServiceByName(name string) Service {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.serviceNameToService[name]
}

// GetPackageByLine gets Package by provided line.
// This ensures thread safety.
func (p *proto) GetPackageByLine(line int) *Package {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.lineToPackage[line]
}

// GetMessageByLine gets message by provided line.
// This ensures thread safety.
func (p *proto) GetMessageByLine(line int) Message {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.lineToMessage[line]
}

// GetEnumByLine gets enum by provided line.
// This ensures thread safety.
func (p *proto) GetEnumByLine(line int) Enum {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.lineToEnum[line]
}

// GetServiceByLine gets service by provided line.
// This ensures thread safety.
func (p *proto) GetServiceByLine(line int) Service {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return p.lineToService[line]
}
