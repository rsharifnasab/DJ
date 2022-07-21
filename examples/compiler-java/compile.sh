#!/usr/bin/env bash

set -euo pipefail

rm -rf out
mkdir -p out
find . -name "*.java" > out/sources.txt
javac -d out --release 11 -cp ".:lib/*" @out/sources.txt
