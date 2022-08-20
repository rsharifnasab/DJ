#!/usr/bin/env bash
set -e
set -o pipefail
set -o nounset

./scripts/build.sh

if false; then
    readonly MARS_URL="https://courses.missouristate.edu/KenVollmar/mars/MARS_4_5_Aug2014/Mars4_5.jar"
    readonly java_mars="./examples/judge-compiler-java/mars.jar"
    readonly python_mars="./examples/judge-compiler-python/mars.jar"

    [[ -e "$java_mars" ]] || curl -o "$java_mars" "$MARS_URL"
    [[ -e "$python_mars" ]] || cp "$java_mars" "$python_mars"
fi

cp -r ./examples/judge-compiler-* ./bin
cp -r ./examples/question-compiler ./bin
cp -r ./examples/COMPILER_README.* ./bin
(
    cd bin
    zip -q -r ./compiler-phase3.zip -- ./*

)
