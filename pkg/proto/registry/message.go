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

type Message interface {
	Protobuf() *protobuf.Message

	GetNestedEnumByName(name string) (Enum, bool)
	GetNestedMessageByName(name string) (Message, bool)

	GetFieldByName(name string) (*MessageField, bool)
	GetOneofFieldByName(name string) (Oneof, bool)
	GetMapFieldByName(name string) (*MapField, bool)

	GetFieldByLine(line int) (*MessageField, bool)
	GetOneofFieldByLine(line int) (Oneof, bool)
	GetMapFieldByLine(line int) (*MapField, bool)
}

type message struct {
	protoMessage *protobuf.Message

	fullyQualifiedName string

	nestedEnumNameToEnum       map[string]Enum
	nestedMessageNameToMessage map[string]Message

	fieldNameToField           map[string]*MessageField
	oneofFieldNameToOneofField map[string]Oneof
	mapFieldNameToMapField     map[string]*MapField

	lineToField      map[int]*MessageField
	lineToOneofField map[int]Oneof
	lineToMapField   map[int]*MapField

	mu *sync.RWMutex
}

var _ Message = (*message)(nil)

func NewMessage(protoMessage *protobuf.Message) *message {
	m := &message{
		protoMessage: protoMessage,

		fullyQualifiedName: "",

		nestedEnumNameToEnum:       make(map[string]Enum),
		nestedMessageNameToMessage: make(map[string]Message),

		fieldNameToField:           make(map[string]*MessageField),
		oneofFieldNameToOneofField: make(map[string]Oneof),
		mapFieldNameToMapField:     make(map[string]*MapField),

		lineToField:      make(map[int]*MessageField),
		lineToOneofField: make(map[int]Oneof),
		lineToMapField:   make(map[int]*MapField),
	}

	for _, e := range protoMessage.Elements {
		switch v := e.(type) {

		case *protobuf.NormalField:
			f := NewMessageField(v)

			m.fieldNameToField[v.Name] = f
			m.lineToField[v.Position.Line] = f

		case *protobuf.Oneof:
			f := NewOneof(v)

			m.oneofFieldNameToOneofField[v.Name] = f
			m.lineToOneofField[v.Position.Line] = f

		case *protobuf.MapField:
			f := NewMapField(v)

			m.mapFieldNameToMapField[v.Name] = f
			m.lineToMapField[v.Position.Line] = f

		default:
		}
	}

	return m
}

// Protobuf returns *protobuf.Proto.
func (m *message) Protobuf() *protobuf.Message {
	return m.protoMessage
}

// GetNestedEnumByName gets enum by provided name.
// This ensures thread safety.
func (m *message) GetNestedEnumByName(name string) (Enum, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	e, ok := m.nestedEnumNameToEnum[name]
	return e, ok
}

// GetNestedMessageByName gets Message by provided name.
// This ensures thread safety.
func (m *message) GetNestedMessageByName(name string) (Message, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	msg, ok := m.nestedMessageNameToMessage[name]
	return msg, ok
}

// GetFieldByName gets MessageField by provided name.
// This ensures thread safety.
func (m *message) GetFieldByName(name string) (*MessageField, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.fieldNameToField[name]
	return f, ok
}

// GetFieldByName gets oneof by provided name.
// This ensures thread safety.
func (m *message) GetOneofFieldByName(name string) (Oneof, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.oneofFieldNameToOneofField[name]
	return f, ok
}

// GetMapFieldByName gets MapField by provided name.
// This ensures thread safety.
func (m *message) GetMapFieldByName(name string) (*MapField, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.mapFieldNameToMapField[name]
	return f, ok
}

// GetFieldByLine gets MessageField by provided line.
// This ensures thread safety.
func (m *message) GetFieldByLine(line int) (*MessageField, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.lineToField[line]
	return f, ok
}

// GetFieldByLine gets oneof by provided line.
// This ensures thread safety.
func (m *message) GetOneofFieldByLine(line int) (Oneof, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.lineToOneofField[line]
	return f, ok
}

// GetMapFieldByLine gets MapField by provided line.
// This ensures thread safety.
func (m *message) GetMapFieldByLine(line int) (*MapField, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	f, ok := m.lineToMapField[line]
	return f, ok
}

type MessageField struct {
	ProtoField *protobuf.NormalField
}

func NewMessageField(protoMessage *protobuf.NormalField) *MessageField {
	return &MessageField{
		ProtoField: protoMessage,
	}
}
