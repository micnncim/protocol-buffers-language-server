// Copyright 2019 The Protocol Buffers Language Server Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

	GetPackageByName(name string) (*Package, bool)
	GetMessageByName(name string) (Message, bool)
	GetEnumByName(name string) (Enum, bool)
	GetServiceByName(name string) (Service, bool)

	GetPackageByLine(line int) (*Package, bool)
	GetMessageByLine(line int) (Message, bool)
	GetEnumByLine(line int) (Enum, bool)
	GetServiceByLine(line int) (Service, bool)

	GetMessageFieldByLine(line int) (*MessageField, bool)
	GetEnumFieldByLine(line int) (*EnumField, bool)
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

// NewProto returns Proto initialized by provided *protobuf.Proto.
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
func (p *proto) GetPackageByName(name string) (*Package, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	pkg, ok := p.packageNameToPackage[name]
	return pkg, ok
}

// GetMessageByName gets message by provided name.
// This ensures thread safety.
func (p *proto) GetMessageByName(name string) (Message, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	m, ok := p.messageNameToMessage[name]
	return m, ok
}

// GetEnumByName gets enum by provided name.
// This ensures thread safety.
func (p *proto) GetEnumByName(name string) (Enum, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	e, ok := p.enumNameToEnum[name]
	return e, ok
}

// GetServiceByName gets service by provided name.
// This ensures thread safety.
func (p *proto) GetServiceByName(name string) (Service, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	s, ok := p.serviceNameToService[name]
	return s, ok
}

// GetPackageByLine gets Package by provided line.
// This ensures thread safety.
func (p *proto) GetPackageByLine(line int) (*Package, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	pkg, ok := p.lineToPackage[line]
	return pkg, ok
}

// GetMessageByLine gets message by provided line.
// This ensures thread safety.
func (p *proto) GetMessageByLine(line int) (Message, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	m, ok := p.lineToMessage[line]
	return m, ok
}

// GetEnumByLine gets enum by provided line.
// This ensures thread safety.
func (p *proto) GetEnumByLine(line int) (Enum, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	e, ok := p.lineToEnum[line]
	return e, ok
}

// GetServiceByLine gets service by provided line.
// This ensures thread safety.
func (p *proto) GetServiceByLine(line int) (Service, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	s, ok := p.lineToService[line]
	return s, ok
}

// GetMessageFieldByLine gets message field by provided line.
// This ensures thread safety.
func (p *proto) GetMessageFieldByLine(line int) (*MessageField, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, message := range p.lineToMessage {
		f, ok := message.GetFieldByLine(line)
		if ok {
			return f, true
		}
	}
	return nil, false
}

// GetEnumFieldByLine gets enum field by provided line.
// This ensures thread safety.
func (p *proto) GetEnumFieldByLine(line int) (*EnumField, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, enum := range p.lineToEnum {
		f, ok := enum.GetFieldByLine(line)
		if ok {
			return f, true
		}
	}
	return nil, false
}
