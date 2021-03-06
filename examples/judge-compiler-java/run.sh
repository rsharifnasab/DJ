#!/bin/bash
set -euo pipefail

readonly COMMAND="$1" #"count" or "test"
readonly TEST_NUMBER="${2:-0}"

echerr() {
    echo "$@" 1>&2
}

test_count() {
    find "testgroup" -name "*.out" | wc -l
    #java -cp "./out:./lib/*" Judger "count"
}

compare() {
    local ACTUAL_FILE="$1"
    local EXPECTED_FILE="$2"
    local actual
    local expected
    actual=$(xargs <"$ACTUAL_FILE" | tr '\n' ' ' |
        awk '{gsub(/^ +| +$/,"")} { print $0 }')
    expected=$(xargs <"$EXPECTED_FILE" | tr '\n' ' ' |
        awk '{gsub(/^ +| +$/,"")} { print $0 }')
    if [[ "$actual" == "$expected" ]]; then
        printf "pass"
    else
        printf "fail"
        echerr ""
        echerr "----------actual:---------"
        echerr "$actual"
        echerr "---------expected:--------"
        echerr "$expected"
        echerr "--------------------------"
    fi
    #python3 compare.py -e "$expected_file" -a "actual.txt"
}

find_file() {
    local NUMBER="$1"
    local POSTFIX="$2"
    local list
    list="$(find "testgroup" -name "*.$POSTFIX" | sort)"
    lines="$(wc -l <<<"$list")"
    if [[ lines -lt NUMBER ]]; then
        echerr "invalid test number [$NUMBER]"
        exit 1
    fi
    echo "$list" | head -n "$NUMBER" | tail -1
}

run_interpreter() {
    local asm_file="$1"
    local inp_file="$2"
    local out_file="$3"

    spim -a -f "$asm_file" <"$inp_file" >"$out_file" 1>&2

    # more info here: http://courses.missouristate.edu/KenVollmar/MARS/Help/MarsHelpCommand.html
    #java -jar ./lib/mars.jar "nc" "ic" "me" "se1" "ae2" "100000" \
    #    "$compiled_file" > "$actual_out" 2> mars_log.txt || true
}

run_code() {
    local src_file="$1"
    local compiled_file="$2"
    java -cp "out:lib/*" Main -i "$src_file" -o "$compiled_file"
}

run_test() {
    local expected_file
    local inp_file
    local src_file

    local compiled_file="compiled.asm"
    local actual_out="actual.txt"

    rm -f "$compiled_file"
    rm -f "$actual_out"

    expected_file=$(find_file "$TEST_NUMBER" "out")
    inp_file=$(find_file "$TEST_NUMBER" "in")
    src_file=$(find_file "$TEST_NUMBER" "d")
    test_name="$TEST_NUMBER"

    #echerr "running test $test_name"

    run_log="$(run_code "$src_file" "$compiled_file")"
    echerr "$run_log" >&2

    if grep -qi "semantic error" "$compiled_file"; then
        cp "$compiled_file" "$actual_out"
    else
        run_interpreter "$compiled_file" "$inp_file" "$actual_out"
        true
        # TODO
    fi

    result="$(compare "$actual_out" "$expected_file")"
    printf "test[%s] - %s\n" "$test_name" "$result"
}

clean() {
    rm "compiled.asm"
    rm "actual.txt"
}

main() {
    cd -P -- "$(dirname -- "$0")"
    if [[ "compile" == "$COMMAND" ]]; then
        ./compile.sh
    elif [[ "count" == "$COMMAND" ]]; then
        test_count
    elif [[ "test" == "$COMMAND" ]]; then
        run_test "$TEST_NUMBER"
    elif [[ "clean" == "$COMMAND" ]]; then
        clean
    fi
}

main
