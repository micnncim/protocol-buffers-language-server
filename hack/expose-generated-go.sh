#! /usr/bin/env bash

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

set -o errexit
set -o nounset
set -o pipefail

if [[ "$#" -ne 2 ]]; then
	echo "usage: $0 <organization> <repository>"
	exit 1
fi

ORGANIZATION=$1
REPOSITORY=$2

OS="$(go env GOHOSTOS)"
ARCH="$(go env GOARCH)"

printf "\e[32;1m>>> Exposing Go generated files\n\e[m"

expose_package () {
	local out_path=$1
	local package=$2
	local old_links=$(eval echo \$"$3")
	local generated_files=$(eval echo \$"$4")

    # Delete all old links
	for f in ${old_links}; do
		if [[ -f "${f}" ]]; then
			# shellcheck disable=SC2059
			printf "\e[32;1m>>> Deleting old link: ${f}\n\e[m"
			rm "${f}"
		fi
	done

	# Compute the relative_path from this package to the bazel-bin
	local count_paths="$(echo -n "${package}" | tr '/' '\n' | wc -l)"
	local relative_path=""
	for i in $(seq 0 "${count_paths}"); do
		relative_path="../${relative_path}"
	done

	local found=0
	for f in ${generated_files}; do
		if [[ -f "${f}" ]]; then
			found=1
            local base=${f##*/}
            printf "\e[32;1m>>> Adding a new link: ${package}/${base}\n\e[m"
			ln -nsf "${relative_path}${f}" "${package}/"
		fi
	done
	if [[ "${found}" == "0" ]]; then
		printf "\e[32;1m>>> Error: No generated file was found inside ${out_path} for the package ${package}\n\e[m"
		exit 1
	fi
}

# For proto go files

for label in $(bazel query 'kind(go_proto_library, //...)'); do
	bazel build "${label}"
done

for label in $(bazel query 'kind(go_proto_library, //...)'); do
	package="${label%%:*}"
	package="${package##//}"
	target="${label##*:}"
	[[ -d "${package}" ]] || continue

	# Compute the path where bazel put the files
	out_path="bazel-bin/${package}/${OS}_${ARCH}_stripped/${target}%/github.com/${ORGANIZATION}/${REPOSITORY}/${package}"

	old_links=$(eval echo "${package}"/*.pb.go)
	generated_files=$(eval echo "${out_path}"/*.pb.go)
	expose_package "${out_path}" "${package}" old_links generated_files
done

# For mock go files

# Build mock go files
for label in $(bazel query 'kind(gomock, //...)'); do
	bazel build "${label}"
done

# Link to the generated files and add them to excluding list in the root BUILD file.
for package in $(bazel query 'kind(gomock, //...)' --output package); do
	# Compute the path where Bazel puts the files.
	out_path="bazel-bin/${package}"

	# shellcheck disable=SC2125
	old_links=${package}/*.mock.go
	# shellcheck disable=SC2125
	generated_files=${out_path}/*.mock.go
	expose_package "${out_path}" "${package}" old_links generated_files
done

# Reset the root BUILD file
cat "${GENERATED_BUILD_FILE}" > "${BUILD_FILE}"
rm "${GENERATED_BUILD_FILE}"
