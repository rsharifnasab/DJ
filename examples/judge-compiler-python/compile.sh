#!/usr/bin/env bash

set -euo pipefail

python -m compileall -q .

rm -rf out
mkdir -p out
# TODO: what to do with out?
