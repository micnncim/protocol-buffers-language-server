package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

// Enum is a registry for protobuf enum.
type Enum interface {
	Protobuf() *protobuf.Enum

	GetFieldByName(name string) *EnumField

	GetFieldByLine(line int) *EnumField
}

type enum struct {
	protoEnum *protobuf.Enum

	fullyQualifiedName string

	fieldNameToValue map[string]*EnumField

	lineToEnumField map[int]*EnumField

	mu *sync.RWMutex
}

var _ Enum = (*enum)(nil)

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

// GetFieldByName gets EnumField  by provided name.
// This ensures thread safety.
func (e *enum) GetFieldByName(name string) *EnumField {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.fieldNameToValue[name]
}

// GetMapFieldByLine gets MapField by provided line.
// This ensures thread safety.
func (e *enum) GetFieldByLine(line int) *EnumField {
	e.mu.RLock()
	defer e.mu.RUnlock()

	return e.lineToEnumField[line]
}

type EnumField struct {
	ProtoEnumField *protobuf.EnumField
}

func NewEnumField(protoMessage *protobuf.EnumField) *EnumField {
	return &EnumField{
		ProtoEnumField: protoMessage,
	}
}
