.PHONY: test
test:
	GO111MODULE=on go test -race ./...

.PHONY: dep
dep:
	GO111MODULE=on go mod tidy
	bazel run //:gazelle -- update-repos -from_file=go.mod
	bazel run //:gazelle

.PHONY: bazel-build
bazel-build:
	bazel build //...

.PHONY: bazel-test
bazel-test:
	bazel test //...

.PHONY: buildifier
buildifier:
	bazel run //:buildifier

.PHONY: bazel-clean
bazel-clean:
	bazel clean

.PHONY: coverage
coverage:
	go test -v -race -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: reviewdog
reviewdog:
	@reviewdog -reporter=github-pr-review

.PHONY: expose-generated-go
expose-generated-go:
	./hack/expose-generated-go.sh micnncim protocol-buffers-language-server