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
	Env *Env
}

// New returns a new Config.
func New() (*Config, error) {
	cnf := &Config{}
	env, err := newEnv()
	if err != nil {
		return nil, err
	}
	cnf.Env = env
	return cnf, nil
}

// Env represents a environment variables for server.
type Env struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

func newEnv() (e *Env, err error) {
	if err = envconfig.Process(envPrefix, e); err != nil {
		return
	}
	return
}
