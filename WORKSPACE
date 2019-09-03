load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

## Go

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/0.19.3/rules_go-0.19.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.3/rules_go-0.19.3.tar.gz",
    ],
    sha256 = "313f2c7a23fecc33023563f082f381a32b9b7254f727a7dd2d6380ccc6dfe09b",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

## gazelle

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "co_honnef_go_tools",
    importpath = "honnef.co/go/tools",
    sum = "h1:/8zB6iBfHCl1qAnEAWwGPNrUvapuy6CPla1VM0k8hQw=",
    version = "v0.0.0-20190106161140-3f1c8253044a",
)

go_repository(
    name = "com_github_anmitsu_go_shlex",
    importpath = "github.com/anmitsu/go-shlex",
    sum = "h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
    version = "v0.0.0-20161002113705-648efa622239",
)

go_repository(
    name = "com_github_beorn7_perks",
    importpath = "github.com/beorn7/perks",
    sum = "h1:xJ4a3vCFaGF/jqvzLMYoU8P317H5OQ+Via4RmuPwCS0=",
    version = "v0.0.0-20180321164747-3a771d992973",
)

go_repository(
    name = "com_github_bradfitz_go_smtpd",
    importpath = "github.com/bradfitz/go-smtpd",
    sum = "h1:ckJgFhFWywOx+YLEMIJsTb+NV6NexWICk5+AMSuz3ss=",
    version = "v0.0.0-20170404230938-deb6d6237625",
)

go_repository(
    name = "com_github_buger_jsonparser",
    importpath = "github.com/buger/jsonparser",
    sum = "h1:D21IyuvjDCshj1/qq+pCNd3VZOAEI9jy6Bi131YlXgI=",
    version = "v0.0.0-20181115193947-bf1c66bbce23",
)

go_repository(
    name = "com_github_burntsushi_toml",
    importpath = "github.com/BurntSushi/toml",
    sum = "h1:WXkYYl6Yr3qBf1K79EBnL4mak0OimBfB0XUf9Vl28OQ=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_client9_misspell",
    importpath = "github.com/client9/misspell",
    sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
    version = "v0.3.4",
)

go_repository(
    name = "com_github_coreos_go_systemd",
    importpath = "github.com/coreos/go-systemd",
    sum = "h1:t5Wuyh53qYyg9eqn4BbnlIT+vmhyww0TatL+zT3uWgI=",
    version = "v0.0.0-20181012123002-c6f51f82210d",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    importpath = "github.com/davecgh/go-spew",
    sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_dustin_go_humanize",
    importpath = "github.com/dustin/go-humanize",
    sum = "h1:VSnTsYCnlFHaM2/igO1h6X3HA71jcobQuxemgkq4zYo=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_flynn_go_shlex",
    importpath = "github.com/flynn/go-shlex",
    sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
    version = "v0.0.0-20150515145356-3f9db97f8568",
)

go_repository(
    name = "com_github_francoispqt_gojay",
    importpath = "github.com/francoispqt/gojay",
    sum = "h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=",
    version = "v1.2.13",
)

go_repository(
    name = "com_github_fsnotify_fsnotify",
    importpath = "github.com/fsnotify/fsnotify",
    sum = "h1:IXs+QLmnXW2CcXuY+8Mzv/fWEsPGWxqefPtCP5CnV9I=",
    version = "v1.4.7",
)

go_repository(
    name = "com_github_ghodss_yaml",
    importpath = "github.com/ghodss/yaml",
    sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_gliderlabs_ssh",
    importpath = "github.com/gliderlabs/ssh",
    sum = "h1:j3L6gSLQalDETeEg/Jg0mGY0/y/N6zI2xX1978P0Uqw=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_go_errors_errors",
    importpath = "github.com/go-errors/errors",
    sum = "h1:LUHzmkK3GUKUrL/1gfBUxAHzcev3apQlezX/+O7ma6w=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_go_language_server_jsonrpc2",
    importpath = "github.com/go-language-server/jsonrpc2",
    sum = "h1:wzfLvaeOS6v1SIkn8jdUq0ba7zaLlqE3Elc8OKeTyT8=",
    version = "v0.2.5",
)

go_repository(
    name = "com_github_go_language_server_protocol",
    importpath = "github.com/go-language-server/protocol",
    sum = "h1:95f98pTF6rnjHIOXuAUloPXvu9+vl/ocdHNhsDQt9ks=",
    version = "v0.4.2",
)

go_repository(
    name = "com_github_go_language_server_uri",
    importpath = "github.com/go-language-server/uri",
    sum = "h1:2p3pYf79r7NSs6Hs7SoyDgWQwQvmQoq4GW3f6Rqooq0=",
    version = "v0.1.2",
)

go_repository(
    name = "com_github_gogo_protobuf",
    importpath = "github.com/gogo/protobuf",
    sum = "h1:72R+M5VuhED/KujmZVcIquuo8mBgX4oVda//DQb3PXo=",
    version = "v1.1.1",
)

go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "com_github_golang_mock",
    importpath = "github.com/golang/mock",
    sum = "h1:28o5sBqPkBsMGnC6b4MvE2TzSr5/AT4c/1fLqVGIwlk=",
    version = "v1.2.0",
)

go_repository(
    name = "com_github_golang_protobuf",
    importpath = "github.com/golang/protobuf",
    sum = "h1:YF8+flBXS5eO826T4nzqPrxfhQThhXl0YzfuUPu4SBg=",
    version = "v1.3.1",
)

go_repository(
    name = "com_github_google_btree",
    importpath = "github.com/google/btree",
    sum = "h1:964Od4U6p2jUkFxvCydnIczKteheJEzHRToSGK3Bnlw=",
    version = "v0.0.0-20180813153112-4030bb1f1f0c",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:Xye71clBPdm5HgqGwUkwhbynsUJZhDbS20FvLhQ2izg=",
    version = "v0.3.1",
)

go_repository(
    name = "com_github_google_go_github",
    importpath = "github.com/google/go-github",
    sum = "h1:N0LgJ1j65A7kfXrZnUDaYCs/Sf4rEjNlfyDHW9dolSY=",
    version = "v17.0.0+incompatible",
)

go_repository(
    name = "com_github_google_go_querystring",
    importpath = "github.com/google/go-querystring",
    sum = "h1:Xkwi/a1rcvNg1PPYe5vI8GbeBY/jrVuDX5ASuANWTrk=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_google_martian",
    importpath = "github.com/google/martian",
    sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
    version = "v2.1.0+incompatible",
)

go_repository(
    name = "com_github_google_pprof",
    importpath = "github.com/google/pprof",
    sum = "h1:eqyIo2HjKhKe/mJzTG8n4VqvLXIOEG+SLdDqX7xGtkY=",
    version = "v0.0.0-20181206194817-3ea8567a2e57",
)

go_repository(
    name = "com_github_googleapis_gax_go",
    importpath = "github.com/googleapis/gax-go",
    sum = "h1:j0GKcs05QVmm7yesiZq2+9cxHkNK9YM6zKx4D2qucQU=",
    version = "v2.0.0+incompatible",
)

go_repository(
    name = "com_github_googleapis_gax_go_v2",
    importpath = "github.com/googleapis/gax-go/v2",
    sum = "h1:siORttZ36U2R/WjiJuDz8znElWBiAlO9rVt+mqJt0Cc=",
    version = "v2.0.3",
)

go_repository(
    name = "com_github_gopherjs_gopherjs",
    importpath = "github.com/gopherjs/gopherjs",
    sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
    version = "v0.0.0-20181017120253-0766667cb4d1",
)

go_repository(
    name = "com_github_gregjones_httpcache",
    importpath = "github.com/gregjones/httpcache",
    sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
    version = "v0.0.0-20180305231024-9cad4c3443a7",
)

go_repository(
    name = "com_github_grpc_ecosystem_grpc_gateway",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
    sum = "h1:WcmKMm43DR7RdtlkEXQJyo5ws8iTp98CyhCCbOHMvNI=",
    version = "v1.5.0",
)

go_repository(
    name = "com_github_jellevandenhooff_dkim",
    importpath = "github.com/jellevandenhooff/dkim",
    sum = "h1:ujPKutqRlJtcfWk6toYVYagwra7HQHbXOaS171b4Tg8=",
    version = "v0.0.0-20150330215556-f50fe3d243e1",
)

go_repository(
    name = "com_github_json_iterator_go",
    importpath = "github.com/json-iterator/go",
    sum = "h1:MrUvLMLTMxbqFJ9kzlvat/rYZqZnW3u4wkLzWTaFwKs=",
    version = "v1.1.6",
)

go_repository(
    name = "com_github_jstemmer_go_junit_report",
    importpath = "github.com/jstemmer/go-junit-report",
    sum = "h1:rBMNdlhTLzJjJSDIjNEXX1Pz3Hmwmz91v+zycvx9PJc=",
    version = "v0.0.0-20190106144839-af01ea7f8024",
)

go_repository(
    name = "com_github_kr_pretty",
    importpath = "github.com/kr/pretty",
    sum = "h1:L/CwN0zerZDmRFUapSPitk6f+Q3+0za1rQkzVuMiMFI=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_kr_pty",
    importpath = "github.com/kr/pty",
    sum = "h1:/Um6a/ZmD5tF7peoOJ5oN5KMQ0DrGVQSXLNwyckutPk=",
    version = "v1.1.3",
)

go_repository(
    name = "com_github_kr_text",
    importpath = "github.com/kr/text",
    sum = "h1:45sCR5RtlFHMR4UwH9sdQ5TC8v0qDQCHnXt+kaKSTVE=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_lunixbochs_vtclean",
    importpath = "github.com/lunixbochs/vtclean",
    sum = "h1:xu2sLAri4lGiovBDQKxl5mrXyESr3gUr5m5SM5+LVb8=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_mailru_easyjson",
    importpath = "github.com/mailru/easyjson",
    sum = "h1:W/GaMY0y69G4cFlmsC6B9sbuo2fP8OFP1ABjt4kPz+w=",
    version = "v0.0.0-20190312143242-1de009706dbe",
)

go_repository(
    name = "com_github_matttproud_golang_protobuf_extensions",
    importpath = "github.com/matttproud/golang_protobuf_extensions",
    sum = "h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_microcosm_cc_bluemonday",
    importpath = "github.com/microcosm-cc/bluemonday",
    sum = "h1:SIYunPjnlXcW+gVfvm0IlSeR5U3WZUOLfVmqg85Go44=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_modern_go_concurrent",
    importpath = "github.com/modern-go/concurrent",
    sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_repository(
    name = "com_github_modern_go_reflect2",
    importpath = "github.com/modern-go/reflect2",
    sum = "h1:9f412s+6RmYXLWZSEzVVgPGK7C2PphHj5RJrvfx9AWI=",
    version = "v1.0.1",
)

go_repository(
    name = "com_github_neelance_astrewrite",
    importpath = "github.com/neelance/astrewrite",
    sum = "h1:D6paGObi5Wud7xg83MaEFyjxQB1W5bz5d0IFppr+ymk=",
    version = "v0.0.0-20160511093645-99348263ae86",
)

go_repository(
    name = "com_github_neelance_sourcemap",
    importpath = "github.com/neelance/sourcemap",
    sum = "h1:eFXv9Nu1lGbrNbj619aWwZfVF5HBrm9Plte8aNptuTI=",
    version = "v0.0.0-20151028013722-8c68805598ab",
)

go_repository(
    name = "com_github_openzipkin_zipkin_go",
    importpath = "github.com/openzipkin/zipkin-go",
    sum = "h1:A/ADD6HaPnAKj3yS7HjGHRK77qi41Hi0DirOOIQAeIw=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:PCj9X21C4pet4sEcElTfAi6LSl5ShkjE8doieLc+cbU=",
    version = "v0.8.2-0.20190227000051-27936f6d90f9",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    importpath = "github.com/pmezard/go-difflib",
    sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_prometheus_client_golang",
    importpath = "github.com/prometheus/client_golang",
    sum = "h1:1921Yw9Gc3iSc4VQh3PIoOqgPCZS7G/4xQNVUp8Mda8=",
    version = "v0.8.0",
)

go_repository(
    name = "com_github_prometheus_client_model",
    importpath = "github.com/prometheus/client_model",
    sum = "h1:idejC8f05m9MGOsuEi1ATq9shN03HrxNkD/luQvxCv8=",
    version = "v0.0.0-20180712105110-5c3871d89910",
)

go_repository(
    name = "com_github_prometheus_common",
    importpath = "github.com/prometheus/common",
    sum = "h1:n/3MEhJQjQxrOUCzh1Y3Re6aJUUWRp2M9+Oc3eVn/54=",
    version = "v0.0.0-20180801064454-c7de2306084e",
)

go_repository(
    name = "com_github_prometheus_procfs",
    importpath = "github.com/prometheus/procfs",
    sum = "h1:agujYaXJSxSo18YNX3jzl+4G6Bstwt+kqv47GS12uL0=",
    version = "v0.0.0-20180725123919-05ee40e3a273",
)

go_repository(
    name = "com_github_russross_blackfriday",
    importpath = "github.com/russross/blackfriday",
    sum = "h1:HyvC0ARfnZBqnXwABFeSZHpKvJHJJfPz81GNueLj0oo=",
    version = "v1.5.2",
)

go_repository(
    name = "com_github_sergi_go_diff",
    importpath = "github.com/sergi/go-diff",
    sum = "h1:Kpca3qRNrduNnOQeazBd0ysaKrUJiIuISHxogkT9RPQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_shurcool_component",
    importpath = "github.com/shurcooL/component",
    sum = "h1:Fth6mevc5rX7glNLpbAMJnqKlfIkcTjZCSHEeqvKbcI=",
    version = "v0.0.0-20170202220835-f88ec8f54cc4",
)

go_repository(
    name = "com_github_shurcool_events",
    importpath = "github.com/shurcooL/events",
    sum = "h1:vabduItPAIz9px5iryD5peyx7O3Ya8TBThapgXim98o=",
    version = "v0.0.0-20181021180414-410e4ca65f48",
)

go_repository(
    name = "com_github_shurcool_github_flavored_markdown",
    importpath = "github.com/shurcooL/github_flavored_markdown",
    sum = "h1:qb9IthCFBmROJ6YBS31BEMeSYjOscSiG+EO+JVNTz64=",
    version = "v0.0.0-20181002035957-2122de532470",
)

go_repository(
    name = "com_github_shurcool_go",
    importpath = "github.com/shurcooL/go",
    sum = "h1:MZM7FHLqUHYI0Y/mQAt3d2aYa0SiNms/hFqC9qJYolM=",
    version = "v0.0.0-20180423040247-9e1955d9fb6e",
)

go_repository(
    name = "com_github_shurcool_go_goon",
    importpath = "github.com/shurcooL/go-goon",
    sum = "h1:llrF3Fs4018ePo4+G/HV/uQUqEI1HMDjCeOf2V6puPc=",
    version = "v0.0.0-20170922171312-37c2f522c041",
)

go_repository(
    name = "com_github_shurcool_gofontwoff",
    importpath = "github.com/shurcooL/gofontwoff",
    sum = "h1:Yoy/IzG4lULT6qZg62sVC+qyBL8DQkmD2zv6i7OImrc=",
    version = "v0.0.0-20180329035133-29b52fc0a18d",
)

go_repository(
    name = "com_github_shurcool_gopherjslib",
    importpath = "github.com/shurcooL/gopherjslib",
    sum = "h1:UOk+nlt1BJtTcH15CT7iNO7YVWTfTv/DNwEAQHLIaDQ=",
    version = "v0.0.0-20160914041154-feb6d3990c2c",
)

go_repository(
    name = "com_github_shurcool_highlight_diff",
    importpath = "github.com/shurcooL/highlight_diff",
    sum = "h1:vYEG87HxbU6dXj5npkeulCS96Dtz5xg3jcfCgpcvbIw=",
    version = "v0.0.0-20170515013008-09bb4053de1b",
)

go_repository(
    name = "com_github_shurcool_highlight_go",
    importpath = "github.com/shurcooL/highlight_go",
    sum = "h1:7pDq9pAMCQgRohFmd25X8hIH8VxmT3TaDm+r9LHxgBk=",
    version = "v0.0.0-20181028180052-98c3abbbae20",
)

go_repository(
    name = "com_github_shurcool_home",
    importpath = "github.com/shurcooL/home",
    sum = "h1:MPblCbqA5+z6XARjScMfz1TqtJC7TuTRj0U9VqIBs6k=",
    version = "v0.0.0-20181020052607-80b7ffcb30f9",
)

go_repository(
    name = "com_github_shurcool_htmlg",
    importpath = "github.com/shurcooL/htmlg",
    sum = "h1:crYRwvwjdVh1biHzzciFHe8DrZcYrVcZFlJtykhRctg=",
    version = "v0.0.0-20170918183704-d01228ac9e50",
)

go_repository(
    name = "com_github_shurcool_httperror",
    importpath = "github.com/shurcooL/httperror",
    sum = "h1:eHRtZoIi6n9Wo1uR+RU44C247msLWwyA89hVKwRLkMk=",
    version = "v0.0.0-20170206035902-86b7830d14cc",
)

go_repository(
    name = "com_github_shurcool_httpfs",
    importpath = "github.com/shurcooL/httpfs",
    sum = "h1:SWV2fHctRpRrp49VXJ6UZja7gU9QLHwRpIPBN89SKEo=",
    version = "v0.0.0-20171119174359-809beceb2371",
)

go_repository(
    name = "com_github_shurcool_httpgzip",
    importpath = "github.com/shurcooL/httpgzip",
    sum = "h1:fxoFD0in0/CBzXoyNhMTjvBZYW6ilSnTw7N7y/8vkmM=",
    version = "v0.0.0-20180522190206-b1c53ac65af9",
)

go_repository(
    name = "com_github_shurcool_issues",
    importpath = "github.com/shurcooL/issues",
    sum = "h1:T4wuULTrzCKMFlg3HmKHgXAF8oStFb/+lOIupLV2v+o=",
    version = "v0.0.0-20181008053335-6292fdc1e191",
)

go_repository(
    name = "com_github_shurcool_issuesapp",
    importpath = "github.com/shurcooL/issuesapp",
    sum = "h1:Y+TeIabU8sJD10Qwd/zMty2/LEaT9GNDaA6nyZf+jgo=",
    version = "v0.0.0-20180602232740-048589ce2241",
)

go_repository(
    name = "com_github_shurcool_notifications",
    importpath = "github.com/shurcooL/notifications",
    sum = "h1:TQVQrsyNaimGwF7bIhzoVC9QkKm4KsWd8cECGzFx8gI=",
    version = "v0.0.0-20181007000457-627ab5aea122",
)

go_repository(
    name = "com_github_shurcool_octicon",
    importpath = "github.com/shurcooL/octicon",
    sum = "h1:bu666BQci+y4S0tVRVjsHUeRon6vUXmsGBwdowgMrg4=",
    version = "v0.0.0-20181028054416-fa4f57f9efb2",
)

go_repository(
    name = "com_github_shurcool_reactions",
    importpath = "github.com/shurcooL/reactions",
    sum = "h1:LneqU9PHDsg/AkPDU3AkqMxnMYL+imaqkpflHu73us8=",
    version = "v0.0.0-20181006231557-f2e0b4ca5b82",
)

go_repository(
    name = "com_github_shurcool_sanitized_anchor_name",
    importpath = "github.com/shurcooL/sanitized_anchor_name",
    sum = "h1:/vdW8Cb7EXrkqWGufVMES1OH2sU9gKVb2n9/1y5NMBY=",
    version = "v0.0.0-20170918181015-86672fcb3f95",
)

go_repository(
    name = "com_github_shurcool_users",
    importpath = "github.com/shurcooL/users",
    sum = "h1:YGaxtkYjb8mnTvtufv2LKLwCQu2/C7qFB7UtrOlTWOY=",
    version = "v0.0.0-20180125191416-49c67e49c537",
)

go_repository(
    name = "com_github_shurcool_webdavfs",
    importpath = "github.com/shurcooL/webdavfs",
    sum = "h1:JtcyT0rk/9PKOdnKQzuDR+FSjh7SGtJwpgVpfZBRKlQ=",
    version = "v0.0.0-20170829043945-18c3829fa133",
)

go_repository(
    name = "com_github_sourcegraph_annotate",
    importpath = "github.com/sourcegraph/annotate",
    sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
    version = "v0.0.0-20160123013949-f4cad6c6324d",
)

go_repository(
    name = "com_github_sourcegraph_syntaxhighlight",
    importpath = "github.com/sourcegraph/syntaxhighlight",
    sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
    version = "v0.0.0-20170531221838-bd320f5d308e",
)

go_repository(
    name = "com_github_stretchr_objx",
    importpath = "github.com/stretchr/objx",
    sum = "h1:4G4v2dO3VZwixGIRoQ5Lfboy6nUhCyYzaqnIAPPhYs4=",
    version = "v0.1.0",
)

go_repository(
    name = "com_github_stretchr_testify",
    importpath = "github.com/stretchr/testify",
    sum = "h1:TivCn/peBQ7UY8ooIcPgZFpTNSz0Q2U6UrFlUfqbe0Q=",
    version = "v1.3.0",
)

go_repository(
    name = "com_github_tarm_serial",
    importpath = "github.com/tarm/serial",
    sum = "h1:UyzmZLoiDWMRywV4DUYb9Fbt8uiOSooupjTq10vpvnU=",
    version = "v0.0.0-20180830185346-98f6abe2eb07",
)

go_repository(
    name = "com_github_viant_assertly",
    importpath = "github.com/viant/assertly",
    sum = "h1:5x1GzBaRteIwTr5RAGFVG14uNeRFxVNbXPWrK2qAgpc=",
    version = "v0.4.8",
)

go_repository(
    name = "com_github_viant_toolbox",
    importpath = "github.com/viant/toolbox",
    sum = "h1:6TteTDQ68CjgcCe8wH3D3ZhUQQOJXMTbj/D9rkk2a1k=",
    version = "v0.24.0",
)

go_repository(
    name = "com_google_cloud_go",
    importpath = "cloud.google.com/go",
    sum = "h1:69FNAINiZfsEuwH3fKq8QrAAnHz+2m4XL4kVYi5BX0Q=",
    version = "v0.37.0",
)

go_repository(
    name = "com_shuralyov_dmitri_app_changes",
    importpath = "dmitri.shuralyov.com/app/changes",
    sum = "h1:hJiie5Bf3QucGRa4ymsAUOxyhYwGEz1xrsVk0P8erlw=",
    version = "v0.0.0-20180602232624-0a106ad413e3",
)

go_repository(
    name = "com_shuralyov_dmitri_html_belt",
    importpath = "dmitri.shuralyov.com/html/belt",
    sum = "h1:SPOUaucgtVls75mg+X7CXigS71EnsfVUK/2CgVrwqgw=",
    version = "v0.0.0-20180602232347-f7d459c86be0",
)

go_repository(
    name = "com_shuralyov_dmitri_service_change",
    importpath = "dmitri.shuralyov.com/service/change",
    sum = "h1:GvWw74lx5noHocd+f6HBMXK6DuggBB1dhVkuGZbv7qM=",
    version = "v0.0.0-20181023043359-a85b471d5412",
)

go_repository(
    name = "com_shuralyov_dmitri_state",
    importpath = "dmitri.shuralyov.com/state",
    sum = "h1:ivON6cwHK1OH26MZyWDCnbTRZZf0IhNsENoNAKFS1g4=",
    version = "v0.0.0-20180228185332-28bcc343414c",
)

go_repository(
    name = "com_sourcegraph_sourcegraph_go_diff",
    importpath = "sourcegraph.com/sourcegraph/go-diff",
    sum = "h1:eTiIR0CoWjGzJcnQ3OkhIl/b9GJovq4lSAVRt0ZFEG8=",
    version = "v0.5.0",
)

go_repository(
    name = "com_sourcegraph_sqs_pbtypes",
    importpath = "sourcegraph.com/sqs/pbtypes",
    sum = "h1:JPJh2pk3+X4lXAkZIk2RuE/7/FoK9maXw+TNPJhVS/c=",
    version = "v0.0.0-20180604144634-d3ebe8f20ae4",
)

go_repository(
    name = "in_gopkg_check_v1",
    importpath = "gopkg.in/check.v1",
    sum = "h1:yhCVgyC4o1eVCa2tZl7eS0r+SDo693bJlVdllGtEeKM=",
    version = "v0.0.0-20161208181325-20d25e280405",
)

go_repository(
    name = "in_gopkg_inf_v0",
    importpath = "gopkg.in/inf.v0",
    sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
    version = "v0.9.1",
)

go_repository(
    name = "in_gopkg_yaml_v2",
    importpath = "gopkg.in/yaml.v2",
    sum = "h1:ZCJp+EgiOT7lHqUV2J862kp8Qj64Jo6az82+3Td9dZw=",
    version = "v2.2.2",
)

go_repository(
    name = "io_opencensus_go",
    importpath = "go.opencensus.io",
    sum = "h1:Mk5rgZcggtbvtAun5aJzAtjKKN/t0R3jJPlWILlv938=",
    version = "v0.18.0",
)

go_repository(
    name = "org_apache_git_thrift_git",
    importpath = "git.apache.org/thrift.git",
    replace = "github.com/apache/thrift",
    sum = "h1:btlJ8oUpVIaJj+kl2yh/ACoII+MltLj1vukud/FwOvw=",
    version = "v0.0.0-20180902110319-2566ecd5d999",
)

go_repository(
    name = "org_go4",
    importpath = "go4.org",
    sum = "h1:+hE86LblG4AyDgwMCLTE6FOlM9+qjHSYS+rKqxUVdsM=",
    version = "v0.0.0-20180809161055-417644f6feb5",
)

go_repository(
    name = "org_go4_grpc",
    importpath = "grpc.go4.org",
    sum = "h1:tmXTu+dfa+d9Evp8NpJdgOy6+rt8/x4yG7qPBrtNfLY=",
    version = "v0.0.0-20170609214715-11d0a25b4919",
)

go_repository(
    name = "org_golang_google_api",
    importpath = "google.golang.org/api",
    sum = "h1:K6z2u68e86TPdSdefXdzvXgR1zEMa+459vBSfWYAZkI=",
    version = "v0.1.0",
)

go_repository(
    name = "org_golang_google_appengine",
    importpath = "google.golang.org/appengine",
    sum = "h1:/wp5JvzpHIxhs/dumFmF7BXTf3Z+dd4uXta4kVyO508=",
    version = "v1.4.0",
)

go_repository(
    name = "org_golang_google_genproto",
    importpath = "google.golang.org/genproto",
    sum = "h1:VOR2wHHZJgoALLvnlCN4JUaWACO1lOLXiSN2F3g/GXU=",
    version = "v0.0.0-20190306203927-b5d61aea6440",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:cfg4PD8YEdSFnm7qLV4++93WcmhH2nIUhMjhdCvl3j8=",
    version = "v1.19.0",
)

go_repository(
    name = "org_golang_x_build",
    importpath = "golang.org/x/build",
    sum = "h1:E2M5QgjZ/Jg+ObCQAudsXxuTsLj7Nl5RV/lZcQZmKSo=",
    version = "v0.0.0-20190111050920-041ab4dc3f9d",
)

go_repository(
    name = "org_golang_x_crypto",
    importpath = "golang.org/x/crypto",
    sum = "h1:YX8ljsm6wXlHZO+aRz9Exqr0evNhKRNe5K/gi+zKh4U=",
    version = "v0.0.0-20190313024323-a1f597ede03a",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:c2HOrn5iMezYjSlGPncknSEr/8x5LELb/ilJbXi9DEA=",
    version = "v0.0.0-20190121172915-509febef88a4",
)

go_repository(
    name = "org_golang_x_lint",
    importpath = "golang.org/x/lint",
    sum = "h1:GmgasJE571dBGXS7E282h2rIZj+KvCLV8z5I6QXbKNI=",
    version = "v0.0.0-20190227174305-5b3e6a55c961",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:actzWV6iWn3GLqN8dZjzsB+CLt+gaV2+wsxroxiQI8I=",
    version = "v0.0.0-20190313220215-9f648a60d977",
)

go_repository(
    name = "org_golang_x_oauth2",
    importpath = "golang.org/x/oauth2",
    sum = "h1:Wo7BWFiOk0QRFMLYMqJGFMd9CgUAcGx7V+qEg/h5IBI=",
    version = "v0.0.0-20190226205417-e64efc72b421",
)

go_repository(
    name = "org_golang_x_perf",
    importpath = "golang.org/x/perf",
    sum = "h1:xYq6+9AtI+xP3M4r0N1hCkHrInHDBohhquRgx9Kk6gI=",
    version = "v0.0.0-20180704124530-6e6d33e29852",
)

go_repository(
    name = "org_golang_x_sync",
    importpath = "golang.org/x/sync",
    sum = "h1:bjcUS9ztw9kFmmIxJInhon/0Is3p+EHBKNgquIzo1OI=",
    version = "v0.0.0-20190227155943-e225da77a7e6",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:yCrMx/EeIue0+Qca57bWZS7VX6ymEoypmhWyPhz0NHM=",
    version = "v0.0.0-20190316082340-a2f829d7f35f",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:z99zHgr7hKfrUcX/KsoJk5FJfjTceCKIp96+biqP4To=",
    version = "v0.3.1-0.20180807135948-17ff2d5776d2",
)

go_repository(
    name = "org_golang_x_time",
    importpath = "golang.org/x/time",
    sum = "h1:fqgJT0MGcGpPgpWU7VRdRjuArfcOvC4AoJmILihzhDg=",
    version = "v0.0.0-20181108054448-85acf8d2951c",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:vamGzbGri8IKo20MQncCuljcQ5uAO6kaCeawQPVblAI=",
    version = "v0.0.0-20190226205152-f727befe758c",
)

go_repository(
    name = "org_golang_x_xerrors",
    importpath = "golang.org/x/xerrors",
    sum = "h1:9zdDQZ7Thm29KFXgAX/+yaf3eVbP7djjWp/dXAppNCc=",
    version = "v0.0.0-20190717185122-a985d3407aa7",
)

go_repository(
    name = "org_uber_go_atomic",
    importpath = "go.uber.org/atomic",
    sum = "h1:cxzIVoETapQEqDhQu3QfnvXAV4AlzcvUCxkVUFw3+EU=",
    version = "v1.4.0",
)

go_repository(
    name = "org_uber_go_multierr",
    importpath = "go.uber.org/multierr",
    sum = "h1:Q0q3m6foxW9fUIiNtElisPPzFNmxZ6jxL00xve0A348=",
    version = "v1.1.1-0.20190429210458-bd075f90b08f",
)

go_repository(
    name = "org_uber_go_zap",
    importpath = "go.uber.org/zap",
    sum = "h1:2/OCbaWT0ZStdlrAU9heVF2FAVUoP+Mky/bPxP2Mqhs=",
    version = "v1.10.1-0.20190430155229-8a2ee5670ced",
)

go_repository(
    name = "com_github_emicklei_proto",
    importpath = "github.com/emicklei/proto",
    sum = "h1:XbpwxmuOPrdES97FrSfpyy67SSCV/wBIKXqgJzh6hNw=",
    version = "v1.6.15",
)
