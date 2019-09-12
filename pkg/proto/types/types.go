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

package types

type ProtoType string

// https://developers.google.com/protocol-buffers/docs/proto3
const (
	Double   ProtoType = "double"
	Float    ProtoType = "float"
	Int32    ProtoType = "int32"
	Int64    ProtoType = "int64"
	Uint32   ProtoType = "uint32"
	Uint64   ProtoType = "uint64"
	Sint32   ProtoType = "sint32"
	Sint64   ProtoType = "sint64"
	Fixed32  ProtoType = "fixed32"
	Fixed64  ProtoType = "fixed64"
	Sfixed32 ProtoType = "sfixed32"
	Sfixed64 ProtoType = "sfixed64"
	Bool     ProtoType = "bool"
	String   ProtoType = "string"
	Bytes    ProtoType = "bytes"
)

var BuildInProtoTypes = []ProtoType{
	Double,
	Float,
	Int32,
	Int64,
	Uint32,
	Uint64,
	Sint32,
	Sint64,
	Fixed32,
	Fixed64,
	Sfixed32,
	Sfixed64,
	Bool,
	String,
	Bytes,
}
