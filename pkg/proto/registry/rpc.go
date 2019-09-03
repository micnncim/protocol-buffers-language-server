package registry

import "github.com/emicklei/proto"

type RPC struct {
	ProtoRPC *proto.RPC
}

func NewRPC(protoRPC *proto.RPC) *RPC {
	return &RPC{
		ProtoRPC: protoRPC,
	}
}
