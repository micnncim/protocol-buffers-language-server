package registry

import "github.com/emicklei/proto"

type MapField struct {
	ProtoMapField *proto.MapField
}

func NewMapField(protoMapField *proto.MapField) *MapField {
	return &MapField{
		ProtoMapField: protoMapField,
	}
}
