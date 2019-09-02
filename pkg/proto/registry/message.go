package registry

import "github.com/emicklei/proto"

type Message struct {
	ProtoMessage *proto.Message

	fullyQualifiedName string

	nestedEnumNameToEnum       map[string]*Enum
	nestedMessageNameToMessage map[string]*Message

	fieldNameToField           map[string]*MessageField
	oneofFieldNameToOneofField map[string]*Oneof
	mapFieldNameToMapField     map[string]*MapField

	LineToField      map[int]*MessageField
	LineToOneofField map[int]*Oneof
	LineToMapField   map[int]*MapField
}

type MessageField struct {
	ProtoField *proto.NormalField
}

func newMessage(protoMessage *proto.Message) *Message {
	m := &Message{
		ProtoMessage: protoMessage,

		fullyQualifiedName: "",

		nestedEnumNameToEnum:       make(map[string]*Enum),
		nestedMessageNameToMessage: make(map[string]*Message),

		fieldNameToField:           make(map[string]*MessageField),
		oneofFieldNameToOneofField: make(map[string]*Oneof),
		mapFieldNameToMapField:     make(map[string]*MapField),

		LineToField:      make(map[int]*MessageField),
		LineToOneofField: make(map[int]*Oneof),
		LineToMapField:   make(map[int]*MapField),
	}

	for _, e := range protoMessage.Elements {
		switch v := e.(type) {

		case *proto.NormalField:
			f := newMessageField(v)

			m.fieldNameToField[v.Name] = f
			m.LineToField[v.Position.Line] = f

		case *proto.Oneof:
			f := &Oneof{protoOneofField: v}

			m.oneofFieldNameToOneofField[v.Name] = f
			m.LineToOneofField[v.Position.Line] = f

		case *proto.MapField:
			f := &MapField{protoMapField: v}

			m.mapFieldNameToMapField[v.Name] = f
			m.LineToMapField[v.Position.Line] = f

		default:
		}
	}

	return m
}

func newMessageField(protoMessage *proto.NormalField) *MessageField {
	return &MessageField{
		ProtoField: protoMessage,
	}
}
