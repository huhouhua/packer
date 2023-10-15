#!/usr/bin/env bash

# 所有构建脚本的通用实用程序、变量和检查。
set -o errexit
set +o nounset
set -o pipefail


# 构建器目录的根目录
PACKER_ROOT=$(dirname "${BASH_SOURCE[0]}")/../..
source "${PACKER_ROOT}/scripts/install/environment.sh"

