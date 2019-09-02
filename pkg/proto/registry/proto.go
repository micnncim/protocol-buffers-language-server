package registry

import "github.com/emicklei/proto"

type Proto struct {
	protoProto *proto.Proto

	messageNameToMessage map[string]*Message
	enumNameToEnum       map[string]*Enum
	serviceNameToService map[string]*Service

	lineToMessage map[int]*Message
	lineToEnum    map[int]*Enum
	lineToService map[int]*Service
}

func NewProto(protoProto *proto.Proto) *Proto {
	p := &Proto{
		protoProto: protoProto,

		messageNameToMessage: make(map[string]*Message),
		enumNameToEnum:       make(map[string]*Enum),
		serviceNameToService: make(map[string]*Service),

		lineToMessage: make(map[int]*Message),
		lineToEnum:    make(map[int]*Enum),
		lineToService: make(map[int]*Service),
	}

	for _, el := range protoProto.Elements {
		switch v := el.(type) {

		case *proto.Message:
			m := newMessage(v)
			p.messageNameToMessage[v.Name] = m
			p.lineToMessage[v.Position.Line] = m

		case *proto.Enum:
			e := newEnum(v)
			p.enumNameToEnum[v.Name] = e
			p.lineToEnum[v.Position.Line] = e

		case *proto.Service:
			s := newService(v)
			p.serviceNameToService[v.Name] = s
			p.lineToService[v.Position.Line] = s

		default:

		}
	}

	return p
}
