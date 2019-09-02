package registry

import "github.com/emicklei/proto"

type Message struct {
	protoMessage *proto.Message

	fullyQualifiedName string

	nestedEnumNameToEnum       map[string]*Enum
	nestedMessageNameToMessage map[string]*Message

	fieldNameToField           map[string]*MessageField
	oneofFieldNameToOneofField map[string]*Oneof
	mapFieldNameToMapField     map[string]*MapField

	lineToField      map[int]*MessageField
	lineToOneofField map[int]*Oneof
	lineToMapField   map[int]*MapField
}

type MessageField struct {
	protoMessage *proto.NormalField
}

func newMessage(protoMessage *proto.Message) *Message {
	m := &Message{
		protoMessage: protoMessage,

		fullyQualifiedName: "",

		nestedEnumNameToEnum:       make(map[string]*Enum),
		nestedMessageNameToMessage: make(map[string]*Message),

		fieldNameToField:           make(map[string]*MessageField),
		oneofFieldNameToOneofField: make(map[string]*Oneof),
		mapFieldNameToMapField:     make(map[string]*MapField),

		lineToField:      make(map[int]*MessageField),
		lineToOneofField: make(map[int]*Oneof),
		lineToMapField:   make(map[int]*MapField),
	}

	for _, e := range protoMessage.Elements {
		switch v := e.(type) {

		case *proto.NormalField:
			f := newMessageField(v)

			m.fieldNameToField[v.Name] = f
			m.lineToField[v.Position.Line] = f

		case *proto.Oneof:
			f := &Oneof{protoOneofField: v}

			m.oneofFieldNameToOneofField[v.Name] = f
			m.lineToOneofField[v.Position.Line] = f

		case *proto.MapField:
			f := &MapField{protoMapField: v}

			m.mapFieldNameToMapField[v.Name] = f
			m.lineToMapField[v.Position.Line] = f

		default:
		}
	}

	return m
}

func newMessageField(protoMessage *proto.NormalField) *MessageField {
	return &MessageField{
		protoMessage: protoMessage,
	}
}
