module github.com/micnncim/protocol-buffers-language-server

go 1.13

require (
	github.com/alecthomas/kingpin v2.2.6+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/bazelbuild/bazelisk v0.0.8
	github.com/emicklei/proto v1.6.15
	github.com/go-language-server/jsonrpc2 v0.3.0
	github.com/go-language-server/protocol v0.5.0
	github.com/go-language-server/uri v0.2.0
	github.com/golang/mock v1.3.1
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/atomic v1.6.0
	go.uber.org/zap v1.10.1-0.20190430155229-8a2ee5670ced
)

// https://thrift.apache.org/lib/go suggests using github
replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
