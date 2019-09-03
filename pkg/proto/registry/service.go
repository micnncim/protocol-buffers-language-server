package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

// Service is a registry for protobuf service.
type Service interface {
	Protobuf() *protobuf.Service

	GetRPCByName(name string) *RPC

	GetRPCByLine(line int) *RPC
}

type service struct {
	protoService *protobuf.Service

	rpcNameToRPC map[string]*RPC

	lineToRPC map[int]*RPC

	mu *sync.RWMutex
}

var _ Service = (*service)(nil)

// NewService returns Service initialized by provided []*protobuf.Service.
func NewService(protoService *protobuf.Service) Service {
	s := &service{
		protoService: protoService,

		rpcNameToRPC: make(map[string]*RPC),

		lineToRPC: make(map[int]*RPC),
	}

	for _, e := range protoService.Elements {
		v, ok := e.(*protobuf.RPC)
		if !ok {
			continue
		}
		r := NewRPC(v)
		s.rpcNameToRPC[v.Name] = r
		s.lineToRPC[v.Position.Line] = r
	}

	return s
}

// Protobuf returns *protobuf.Service.
func (s *service) Protobuf() *protobuf.Service {
	return s.protoService
}

// GetRPCByName gets RPC by provided name.
// This ensures thread safety.
func (s *service) GetRPCByName(name string) *RPC {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.rpcNameToRPC[name]
}

// GetRPCByLine gets RPC by provided line.
// This ensures thread safety.
func (s *service) GetRPCByLine(line int) *RPC {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.lineToRPC[line]
}

// RPC is a registry for protobuf rpc.
type RPC struct {
	ProtoRPC *protobuf.RPC
}

// NewRPC returns RPC initialized by provided []*protobuf.RPC.
func NewRPC(protoRPC *protobuf.RPC) *RPC {
	return &RPC{
		ProtoRPC: protoRPC,
	}
}
