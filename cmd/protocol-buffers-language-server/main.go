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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/go-language-server/jsonrpc2"

	"github.com/micnncim/protocol-buffers-language-server/pkg/config"
	"github.com/micnncim/protocol-buffers-language-server/pkg/logging"
	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/server"
	"github.com/micnncim/protocol-buffers-language-server/pkg/lsp/source"
)

var (
	logfile = kingpin.Flag("logfile", "Filename to log.").String()
	address = kingpin.Flag("address", "Address on run server. Use for debugging purposes.").String()
	port    = kingpin.Flag("port", "Port on run server. Use for debugging purposes.").Int()
	debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
)

var (
	stdout = os.Stdout
	stderr = os.Stderr
)

func main() {
	cfg, err := config.New()
	if err != nil {
		exit(err)
	}

	logger, err := logging.NewLogger(cfg.Env.LogLevel)
	if err != nil {
		exit(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	session := source.NewSession()

	if err := runServer(ctx, session, server.WithLogger(logger)); err != nil {
		exit(err)
	}
}

func runServer(ctx context.Context, session source.Session, opts ...server.Option) error {
	run := func(srv *server.Server) {
		go srv.Run(ctx)
	}

	if *address != "" {
		return server.RunServerOnAddress(ctx, session, *address, run)
	}

	if *port != 0 {
		return server.RunServerOnPort(ctx, session, *port, run)
	}

	stream := jsonrpc2.NewStream(stdout, stderr)
	srv := server.New(ctx, session, stream)

	return srv.Run(ctx)
}

func exit(err error) {
	if err == nil {
		return
	}
	fmt.Fprint(stderr, err)
	os.Exit(1)
}
