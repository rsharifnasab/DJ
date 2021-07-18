#!/usr/bin/env bash

set -o errexit
set -o nounset

command=$1

folderPath=$2

if [ "$command" = "compile" ]; then
    echo "compile"

    find "$folderPath" -name "*.java" > compiled/sources.txt

    cp ./*.jar compiled
    javac --release 11 -cp .:junit-platform-console.jar:hamcrest.jar @compiled/sources.txt -d "compiled"


elif [ "$command" = "run" ]; then
    echo "run"

    java -jar junit-platform-console.jar --cp "compiled" --scan-class-path --fail-if-no-tests  --reports-dir="reportdir" --disable-banner

else 
    echo "error"
fi


