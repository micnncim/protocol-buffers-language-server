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

// Oneof is a registry for protobuf oneof.
type Oneof interface {
	Protobuf() *protobuf.Oneof

	GetFieldByName(name string) (*OneofField, bool)

	GetFieldByLine(line int) (*OneofField, bool)
}

type oneof struct {
	protoOneofField *protobuf.Oneof

	fieldNameToField map[string]*OneofField

	lineToField map[int]*OneofField

	mu *sync.RWMutex
}

var _ Oneof = (*oneof)(nil)

// NewOneof returns Oneof initialized by provided *protobuf.Oneof.
func NewOneof(protoOneofField *protobuf.Oneof) Oneof {
	oneof := &oneof{
		protoOneofField: protoOneofField,

		fieldNameToField: make(map[string]*OneofField),

		lineToField: make(map[int]*OneofField),
	}

	for _, e := range protoOneofField.Elements {
		v, ok := e.(*protobuf.OneOfField)
		if !ok {
			continue
		}
		f := NewOneofField(v)
		oneof.fieldNameToField[v.Name] = f
		oneof.lineToField[v.Position.Line] = f
	}

	return oneof
}

// Protobuf returns *protobuf.Oneof.
func (o *oneof) Protobuf() *protobuf.Oneof {
	return o.protoOneofField
}

// GetFieldByName gets EnumField  by provided name.
// This ensures thread safety.
func (o *oneof) GetFieldByName(name string) (*OneofField, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	f, ok := o.fieldNameToField[name]
	return f, ok
}

// GetFieldByName gets MapField by provided line.
// This ensures thread safety.
func (o *oneof) GetFieldByLine(line int) (*OneofField, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	f, ok := o.lineToField[line]
	return f, ok
}

// OneofField is a registry for protobuf oneof field.
type OneofField struct {
	ProtoOneOfField *protobuf.OneOfField
}

// NewOneofField returns OneofField initialized by provided *protobuf.OneofField.
func NewOneofField(protoOneOfField *protobuf.OneOfField) *OneofField {
	return &OneofField{
		ProtoOneOfField: protoOneOfField,
	}
}
