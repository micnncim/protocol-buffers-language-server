// Copyright 2019 The Protocol Buffers Language Server Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"sync"

	protobuf "github.com/emicklei/proto"
)

// Service is a registry for protobuf service.
type Service interface {
	Protobuf() *protobuf.Service

	GetRPCByName(bool string) (*RPC, bool)

	GetRPCByLine(line int) (*RPC, bool)
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
func (s *service) GetRPCByName(name string) (*RPC, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	r, ok := s.rpcNameToRPC[name]
	return r, ok
}

// GetRPCByLine gets RPC by provided line.
// This ensures thread safety.
func (s *service) GetRPCByLine(line int) (*RPC, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	r, ok := s.lineToRPC[line]
	return r, ok
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
