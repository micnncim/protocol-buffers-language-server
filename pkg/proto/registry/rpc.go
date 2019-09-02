package registry

import "github.com/emicklei/proto"

type RPC struct {
	protoRPC *proto.RPC
}

func newRPC(protoRPC *proto.RPC) *RPC {
	return &RPC{
		protoRPC: protoRPC,
	}
}
