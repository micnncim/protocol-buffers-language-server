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

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "PROTOBUF_LSP"

// Config represents a configuration for server.
type Config struct {
	Env    Env
	Server Server
	Log    Log
}

// Env represents a environment variables for server.
type Env struct {
}

// Server represents a configuration for server.
type Server struct {
	Address string
	Port    int
}

// Log represents a configuration for zap.Logger.
type Log struct {
	File  string
	Level string
}

func newEnv() (Env, error) {
	env := Env{}
	if err := envconfig.Process(envPrefix, env); err != nil {
		return Env{}, nil
	}
	return env, nil
}
