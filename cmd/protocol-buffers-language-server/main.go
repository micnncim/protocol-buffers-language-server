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
	stdout = os.Stdout
	stderr = os.Stderr
)

var cfg config.Config

func init() {
	kingpin.Flag("logfile", "Filename to log.").StringVar(&cfg.Log.File)
	kingpin.Flag("loglevel", "Level of logging.").Default("info").StringVar(&cfg.Log.Level)
	kingpin.Flag("address", "Address on run server. Use for debugging purposes.").StringVar(&cfg.Server.Address)
	kingpin.Flag("port", "Port on run server. Use for debugging purposes.").IntVar(&cfg.Server.Port)
}

func main() {
	kingpin.Parse()

	logger, err := logging.NewLogger(cfg.Log)
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
	run := func(ctx context.Context, srv *server.Server) {
		go srv.Run(ctx)
	}

	if address := cfg.Server.Address; address != "" {
		return server.RunServerOnAddress(ctx, session, address, run, opts...)
	}

	if port := cfg.Server.Port; port != 0 {
		return server.RunServerOnPort(ctx, session, port, run, opts...)
	}

	stream := jsonrpc2.NewStream(stdout, stderr)
	ctx, srv := server.NewServer(ctx, session, stream, opts...)

	return srv.Run(ctx)
}

func exit(err error) {
	if err == nil {
		return
	}
	fmt.Fprint(stderr, err)
	os.Exit(1)
}
