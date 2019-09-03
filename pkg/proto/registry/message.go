package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

type Message interface {
	Protobuf() *protobuf.Message

	GetNestedEnumByName(name string) Enum
	GetNestedMessageByName(name string) Message

	GetFieldByName(name string) *MessageField
	GetOneofFieldByName(name string) Oneof
	GetMapFieldByName(name string) *MapField

	GetFieldByLine(line int) *MessageField
	GetOneofFieldByLine(line int) Oneof
	GetMapFieldByLine(line int) *MapField
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
	lineToOneofField map[int]*oneof
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
		lineToOneofField: make(map[int]*oneof),
		lineToMapField:   make(map[int]*MapField),
	}

	for _, e := range protoMessage.Elements {
		switch v := e.(type) {

		case *protobuf.NormalField:
			f := NewMessageField(v)

			m.fieldNameToField[v.Name] = f
			m.lineToField[v.Position.Line] = f

		case *protobuf.Oneof:
			f := &oneof{protoOneofField: v}

			m.oneofFieldNameToOneofField[v.Name] = f
			m.lineToOneofField[v.Position.Line] = f

		case *protobuf.MapField:
			f := &MapField{ProtoMapField: v}

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
func (m *message) GetNestedEnumByName(name string) Enum {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.nestedEnumNameToEnum[name]
}

// GetNestedMessageByName gets Message by provided name.
// This ensures thread safety.
func (m *message) GetNestedMessageByName(name string) Message {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.nestedMessageNameToMessage[name]
}

// GetFieldByName gets MessageField by provided name.
// This ensures thread safety.
func (m *message) GetFieldByName(name string) *MessageField {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.fieldNameToField[name]
}

// GetFieldByName gets oneof by provided name.
// This ensures thread safety.
func (m *message) GetOneofFieldByName(name string) Oneof {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.oneofFieldNameToOneofField[name]
}

// GetMapFieldByName gets MapField by provided name.
// This ensures thread safety.
func (m *message) GetMapFieldByName(name string) *MapField {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.mapFieldNameToMapField[name]
}

// GetFieldByLine gets MessageField by provided line.
// This ensures thread safety.
func (m *message) GetFieldByLine(line int) *MessageField {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.lineToField[line]
}

// GetFieldByLine gets oneof by provided line.
// This ensures thread safety.
func (m *message) GetOneofFieldByLine(line int) Oneof {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.lineToOneofField[line]
}

// GetMapFieldByLine gets MapField by provided line.
// This ensures thread safety.
func (m *message) GetMapFieldByLine(line int) *MapField {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.lineToMapField[line]
}

type MessageField struct {
	ProtoField *protobuf.NormalField
}

func NewMessageField(protoMessage *protobuf.NormalField) *MessageField {
	return &MessageField{
		ProtoField: protoMessage,
	}
}
