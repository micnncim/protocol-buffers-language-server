# Copyright 2019 The Protocol Buffers Language Server Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: dep
dep: bin/bazelisk
	go mod tidy
	bin/bazelisk run //:gazelle
	bin/bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=bazel/deps.bzl%go_repositories

.PHONY: run
run: bin/bazelisk
	bin/bazelisk run //cmd/protocol-buffers-language-server

.PHONY: build
build: bin/bazelisk
	bin/bazelisk build //...

.PHONY: test
test: bin/bazelisk
	bin/bazelisk test //...

.PHONY: buildifier
buildifier: bin/bazelisk
	bin/bazelisk run //:buildifier

.PHONY: clean
clean: bin/bazelisk
	bin/bazelisk clean

.PHONY: coverage
coverage:
	go test -v -race -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt ./...

.PHONY: expose-generated-go
 expose-generated-go:
	./hack/expose-generated-go.sh micnncim protocol-buffers-language-server

bin/bazelisk:
	@mkdir -p bin
	go build -o bin/bazelisk github.com/bazelbuild/bazelisk
