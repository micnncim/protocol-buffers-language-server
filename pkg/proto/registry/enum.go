package registry

import "github.com/emicklei/proto"

type Enum struct {
	ProtoEnum *proto.Enum

	FullyQualifiedName string

	FieldNameToValue map[string]*EnumField

	LineToEnumField map[int]*EnumField
}

type EnumField struct {
	ProtoEnumField *proto.EnumField
}

func NewEnum(protoEnum *proto.Enum) *Enum {
	enum := &Enum{
		ProtoEnum: protoEnum,

		FullyQualifiedName: "",

		FieldNameToValue: make(map[string]*EnumField),

		LineToEnumField: make(map[int]*EnumField),
	}

	for _, e := range protoEnum.Elements {
		v, ok := e.(*proto.EnumField)
		if !ok {
			continue
		}
		f := NewEnumField(v)
		enum.FieldNameToValue[v.Name] = f
		enum.LineToEnumField[v.Position.Line] = f
	}

	return enum
}

func NewEnumField(protoMessage *proto.EnumField) *EnumField {
	return &EnumField{
		ProtoEnumField: protoMessage,
	}
}
