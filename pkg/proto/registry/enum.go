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

// Enum is a registry for protobuf enum.
type Enum interface {
	Protobuf() *protobuf.Enum

	GetFieldByName(name string) (*EnumField, bool)

	GetFieldByLine(line int) (*EnumField, bool)
}

type enum struct {
	protoEnum *protobuf.Enum

	fullyQualifiedName string

	fieldNameToValue map[string]*EnumField

	lineToEnumField map[int]*EnumField

	mu *sync.RWMutex
}

var _ Enum = (*enum)(nil)

// NewEnum returns Enum initialized by provided *protobuf.Enum.
func NewEnum(protoEnum *protobuf.Enum) Enum {
	enum := &enum{
		protoEnum: protoEnum,

		fullyQualifiedName: "",

		fieldNameToValue: make(map[string]*EnumField),

		lineToEnumField: make(map[int]*EnumField),
	}

	for _, e := range protoEnum.Elements {
		v, ok := e.(*protobuf.EnumField)
		if !ok {
			continue
		}
		f := NewEnumField(v)
		enum.fieldNameToValue[v.Name] = f
		enum.lineToEnumField[v.Position.Line] = f
	}

	return enum
}

// Protobuf returns *protobuf.Enum.
func (e *enum) Protobuf() *protobuf.Enum {
	return e.protoEnum
}

// GetFieldByName gets EnumField by provided name.
// This ensures thread safety.
func (e *enum) GetFieldByName(name string) (f *EnumField, ok bool) {
	e.mu.RLock()
	f, ok = e.fieldNameToValue[name]
	e.mu.RUnlock()
	return
}

// GetMapFieldByLine gets MapField by provided line.
// This ensures thread safety.
func (e *enum) GetFieldByLine(line int) (f *EnumField, ok bool) {
	e.mu.RLock()
	f, ok = e.lineToEnumField[line]
	e.mu.RUnlock()
	return
}

// EnumField is a registry for protobuf enum field.
type EnumField struct {
	ProtoEnumField *protobuf.EnumField
}

// NewEnumField returns EnumField initialized by provided *protobuf.EnumField.
func NewEnumField(protoMessage *protobuf.EnumField) *EnumField {
	return &EnumField{
		ProtoEnumField: protoMessage,
	}
}
