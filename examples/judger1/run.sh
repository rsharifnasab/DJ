#!/bin/bash

set -e
set -o pipefail
set -u

COMMAND="$1"  #"count" or "test"
if [[ "count" == "$COMMAND" ]]; then
    echo "5"
fi
