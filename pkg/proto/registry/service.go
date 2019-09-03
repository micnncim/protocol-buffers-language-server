package registry

import "github.com/emicklei/proto"

type Service struct {
	ProtoService *proto.Service

	RPCNameToRPC map[string]*RPC

	LineToRPC map[int]*RPC
}

func NewService(protoService *proto.Service) *Service {
	s := &Service{
		ProtoService: protoService,

		RPCNameToRPC: make(map[string]*RPC),

		LineToRPC: make(map[int]*RPC),
	}

	for _, e := range protoService.Elements {
		v, ok := e.(*proto.RPC)
		if !ok {
			continue
		}
		r := NewRPC(v)
		s.RPCNameToRPC[v.Name] = r
		s.LineToRPC[v.Position.Line] = r
	}

	return s
}
