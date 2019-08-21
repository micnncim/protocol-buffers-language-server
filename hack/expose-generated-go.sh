#! /usr/bin/env bash

set -euo pipefail

if [[ "$#" -ne 2 ]]; then
	echo "usage: $0 <organization> <repository>"
	exit 1
fi

ORGANIZATION=$1
REPOSITORY=$2

OS="$(go env GOHOSTOS)"
ARCH="$(go env GOARCH)"

echo -e ">>> Exposing Go generated files"

expose_package () {
	local out_path=$1
	local package=$2
	local old_links=$(eval echo \$$3)
	local generated_files=$(eval echo \$$4)

    # Delete all old links
	for f in ${old_links}; do
		if [[ -f "${f}" ]]; then
			echo ">>> Deleting old link: ${f}"
			rm ${f}
		fi
	done

	# Compute the relative_path from this package to the bazel-bin
	local count_paths="$(echo -n "${package}" | tr '/' '\n' | wc -l)"
	local relative_path=""
	for i in $(seq 0 ${count_paths}); do
		relative_path="../${relative_path}"
	done

	local found=0
	for f in ${generated_files}; do
		if [[ -f "${f}" ]]; then
			found=1
            local base=${f##*/}
            echo ">>> Adding a new link: ${package}/${base}"
			ln -nsf "${relative_path}${f}" "${package}/"
		fi
	done
	if [[ "${found}" == "0" ]]; then
		echo ">>> Error: No generated file was found inside ${out_path} for the package ${package}"
		exit 1
	fi
}

####################
# For proto go files
####################

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

	old_links=$(eval echo ${package}/*.pb.go)
	generated_files=$(eval echo ${out_path}/*.pb.go)
	expose_package ${out_path} ${package} old_links generated_files
done
