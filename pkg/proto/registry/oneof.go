package registry

import (
	"sync"

	"github.com/emicklei/proto"
)

type Oneof struct {
	ProtoOneofField *proto.Oneof

	FieldNameToField map[string]*OneofField

	LineToField map[int]*OneofField

	mu *sync.RWMutex
}

type OneofField struct {
	ProtoOneOfField *proto.OneOfField
}

func NewOneof(protoOneofField *proto.Oneof) *Oneof {
	oneof := &Oneof{
		ProtoOneofField: protoOneofField,

		FieldNameToField: make(map[string]*OneofField),

		LineToField: make(map[int]*OneofField),
	}

	for _, e := range protoOneofField.Elements {
		v, ok := e.(*proto.OneOfField)
		if !ok {
			continue
		}
		f := NewOneofField(v)
		oneof.FieldNameToField[v.Name] = f
		oneof.LineToField[v.Position.Line] = f
	}

	return oneof
}

func NewOneofField(protoOneOfField *proto.OneOfField) *OneofField {
	return &OneofField{
		ProtoOneOfField: protoOneOfField,
	}
}
