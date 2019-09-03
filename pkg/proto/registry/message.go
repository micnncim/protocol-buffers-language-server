package registry

import (
	"sync"

	"github.com/emicklei/proto"
)

type Message struct {
	ProtoMessage *proto.Message

	FullyQualifiedName string

	NestedEnumNameToEnum       map[string]*Enum
	NestedMessageNameToMessage map[string]*Message

	FieldNameToField           map[string]*MessageField
	OneofFieldNameToOneofField map[string]*Oneof
	MapFieldNameToMapField     map[string]*MapField

	LineToField      map[int]*MessageField
	LineToOneofField map[int]*Oneof
	LineToMapField   map[int]*MapField

	mu *sync.RWMutex
}

type MessageField struct {
	ProtoField *proto.NormalField
}

func NewMessage(protoMessage *proto.Message) *Message {
	m := &Message{
		ProtoMessage: protoMessage,

		FullyQualifiedName: "",

		NestedEnumNameToEnum:       make(map[string]*Enum),
		NestedMessageNameToMessage: make(map[string]*Message),

		FieldNameToField:           make(map[string]*MessageField),
		OneofFieldNameToOneofField: make(map[string]*Oneof),
		MapFieldNameToMapField:     make(map[string]*MapField),

		LineToField:      make(map[int]*MessageField),
		LineToOneofField: make(map[int]*Oneof),
		LineToMapField:   make(map[int]*MapField),
	}

	for _, e := range protoMessage.Elements {
		switch v := e.(type) {

		case *proto.NormalField:
			f := NewMessageField(v)

			m.FieldNameToField[v.Name] = f
			m.LineToField[v.Position.Line] = f

		case *proto.Oneof:
			f := &Oneof{ProtoOneofField: v}

			m.OneofFieldNameToOneofField[v.Name] = f
			m.LineToOneofField[v.Position.Line] = f

		case *proto.MapField:
			f := &MapField{ProtoMapField: v}

			m.MapFieldNameToMapField[v.Name] = f
			m.LineToMapField[v.Position.Line] = f

		default:
		}
	}

	return m
}

func NewMessageField(protoMessage *proto.NormalField) *MessageField {
	return &MessageField{
		ProtoField: protoMessage,
	}
}
