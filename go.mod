module github.com/micnncim/protocol-buffers-language-server

go 1.12

require (
	github.com/emicklei/proto v1.6.15
	github.com/go-language-server/jsonrpc2 v0.2.5
	github.com/go-language-server/protocol v0.4.2
	github.com/go-language-server/uri v0.1.2
	github.com/golang/mock v1.2.0
	go.uber.org/atomic v1.4.0
	go.uber.org/multierr v1.1.1-0.20190429210458-bd075f90b08f
	go.uber.org/zap v1.10.1-0.20190430155229-8a2ee5670ced
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7
)

// https://thrift.apache.org/lib/go suggests using github
replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
