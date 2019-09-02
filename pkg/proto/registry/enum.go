package registry

import "github.com/emicklei/proto"

type Enum struct {
	protoEnum *proto.Enum

	fullyQualifiedName string

	fieldNameToValue map[string]*EnumField

	LineToEnumField map[int]*EnumField
}

type EnumField struct {
	protoEnumField *proto.EnumField
}

func newEnum(protoEnum *proto.Enum) *Enum {
	enum := &Enum{
		protoEnum: protoEnum,

		fullyQualifiedName: "",

		fieldNameToValue: make(map[string]*EnumField),

		LineToEnumField: make(map[int]*EnumField),
	}

	for _, e := range protoEnum.Elements {
		v, ok := e.(*proto.EnumField)
		if !ok {
			continue
		}
		f := newEnumField(v)
		enum.fieldNameToValue[v.Name] = f
		enum.LineToEnumField[v.Position.Line] = f
	}

	return enum
}

func newEnumField(protoMessage *proto.EnumField) *EnumField {
	return &EnumField{
		protoEnumField: protoMessage,
	}
}
