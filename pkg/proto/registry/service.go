package registry

import "github.com/emicklei/proto"

type Service struct {
	protoService *proto.Service

	rpcNameToRPC map[string]*RPC

	lineToRPC map[int]*RPC
}

func newService(protoService *proto.Service) *Service {
	s := &Service{
		protoService: protoService,

		rpcNameToRPC: make(map[string]*RPC),

		lineToRPC: make(map[int]*RPC),
	}

	for _, e := range protoService.Elements {
		v, ok := e.(*proto.RPC)
		if !ok {
			continue
		}
		r := newRPC(v)
		s.rpcNameToRPC[v.Name] = r
		s.lineToRPC[v.Position.Line] = r
	}

	return s
}
