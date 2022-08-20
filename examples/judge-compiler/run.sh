#!/bin/bash
set -euo pipefail

readonly COMMAND="$1" #"count" or "test"
readonly LANG="$2"    #
readonly TEST_NUMBER="${3:-0}"

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
    actual=$(sed 's/[[:space:]]/ /g' <"$ACTUAL_FILE")
    expected=$(sed 's/[[:space:]]/ /g' <"$EXPECTED_FILE")
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

    spim -a -quiet -f "$asm_file" <"$inp_file" >"$out_file" 2>"spim.log" || true
    cat "spim.log" 1>&2 || true
    if grep -qi "error" "spim.log"; then
        echerr "skipping spim"
        cp "$asm_file" "$out_file"
    fi
    if head -1 "$out_file" | grep  -q "SPIM Version" ; then
        echerr "old spim detected"
        sed -i "1,5d" "$out_file"
    fi

    # more info here:
    # http://courses.missouristate.edu/KenVollmar/MARS/Help/MarsHelpCommand.html
    #java -jar ./mars.jar "nc" "ic" "me" "se1" "ae2" "100000" \
    #    "$asm_file"  < "$inp_file"  > "$out_file" || {
    #    echerr "skipping spim"
    #    cp "$asm_file" "$out_file"
    #}

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
    echerr "----------<runlog>---------"
    echerr "$run_log"
    echerr "---------</runlog>---------"

    run_interpreter "$compiled_file" "$inp_file" "$actual_out"

    result="$(compare "$actual_out" "$expected_file")"
    printf "test[%s] - %s\n" "$test_name" "$result"
}

clean() {
    rm "compiled.asm"
    rm "actual.txt"
}

run_code() {
    src_file="$1"
    compiled_file="$2"
    "./lang/${LANG}.sh" run "$src_file" "$compiled_file"
}

compile_code() {
    "./lang/${LANG}.sh" compile
}

main() {
    cd -P -- "$(dirname -- "$0")"
    if [[ "compile" == "$COMMAND" ]]; then
        compile_code
    elif [[ "count" == "$COMMAND" ]]; then
        test_count
    elif [[ "test"  == "$COMMAND" ]]; then
        run_test "$TEST_NUMBER"
    elif [[ "clean" == "$COMMAND" ]]; then
        clean
    fi
}

main
