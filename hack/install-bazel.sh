#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

version=0.28.1
tmpdir=tmp

if [[ $(uname) == 'Darwin' ]]; then
    os='darwin'
elif [[ $(uname) =  'Linux' ]]; then
    os='linux'
else
    echo "OS should be Darwin or Linux"
    exit 1
fi

mkdir -p ${tmpdir}
printf "\e[32;1mInstalling installer script...\n\e[m"
curl -sSL \
  https://github.com/bazelbuild/bazel/releases/download/${version}/bazel-${version}-installer-${os}-x86_64.sh \
  -o ${tmpdir}/bazel-${version}-installer-${os}-x86_64.sh
chmod +x ${tmpdir}/bazel-${version}-installer-${os}-x86_64.sh
printf "\e[32;1mExecuting installer script...\n\e[m"
./${tmpdir}/bazel-${version}-installer-${os}-x86_64.sh
rm -rf ${tmpdir}
