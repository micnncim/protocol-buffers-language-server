package registry

import "github.com/emicklei/proto"

type MapField struct {
	protoMapField *proto.MapField
}

func NewMapField(protoMapField *proto.MapField) *MapField {
	return &MapField{
		protoMapField: protoMapField,
	}
}
