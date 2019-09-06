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

.PHONY: test
test:
	GO111MODULE=on go test -race ./...

.PHONY: dep
dep:
	GO111MODULE=on go mod tidy
	bazel run //:gazelle
	bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=bazel/deps.bzl%go_repositories

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
