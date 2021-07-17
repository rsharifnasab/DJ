#!/usr/bin/env bash

set -o errexit
set -o nounset

command=@1

filename=@2

if [ $command = "compile" ]; then
    echo "compile"

    javac --release 11 -d "compiled" $filename

elif [ "$command" = "run" ]; then
    echo "run"

    classFile=$(basename "$filename" .java)
    java -cp "compiled" "$classFile"

else 
    echo "error"
fi

