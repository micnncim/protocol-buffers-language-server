package registry

import protobuf "github.com/emicklei/proto"

type MapField struct {
	ProtoMapField *protobuf.MapField
}

func NewMapField(protoMapField *protobuf.MapField) *MapField {
	return &MapField{
		ProtoMapField: protoMapField,
	}
}
