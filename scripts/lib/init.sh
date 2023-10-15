#!/usr/bin/env bash


set -o errexit
set +o nounset
set -o pipefail

unset CDPATH

# Default use go modules
export GO111MODULE=on

# The root of the build/dist directory
PACKER_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"


