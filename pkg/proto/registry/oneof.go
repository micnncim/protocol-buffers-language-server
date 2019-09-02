package registry

import "github.com/emicklei/proto"

type Oneof struct {
	protoOneofField *proto.Oneof

	fieldNameToField map[string]*OneofField
	lineToField      map[int]*OneofField
}

type OneofField struct {
	protoOneOfField *proto.OneOfField
}

func newOneof(protoOneofField *proto.Oneof) *Oneof {
	oneof := &Oneof{
		protoOneofField: protoOneofField,

		fieldNameToField: make(map[string]*OneofField),
		lineToField:      make(map[int]*OneofField),
	}

	for _, e := range protoOneofField.Elements {
		v, ok := e.(*proto.OneOfField)
		if !ok {
			continue
		}
		f := newOneofField(v)
		oneof.fieldNameToField[v.Name] = f
		oneof.lineToField[v.Position.Line] = f
	}

	return oneof
}

func newOneofField(protoOneOfField *proto.OneOfField) *OneofField {
	return &OneofField{
		protoOneOfField: protoOneOfField,
	}
}
