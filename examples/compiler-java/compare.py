#!/usr/bin/env python

import argparse
from re import compile as re_compile

parser = argparse.ArgumentParser()
parser.add_argument(
    "-a", "--actual",
    dest="actual",
    default="./actual.txt",
    type=str,
    help="actual file from the program",
)
parser.add_argument(
    "-e", "--expected",
    dest="expected",
    required=True,
    type=str,
    help="expected file from the TA",
)


def equal(a, b):
    regex = re_compile(r'[^\w]')
    return regex.sub('', a) == regex.sub('', b)


def main():
    args = parser.parse_args()
    actual = open(args.actual).read()
    expected = open(args.expected).read()
    if equal(expected, actual):
        print("pass")
    else:
        print("fail")


if __name__ == "__main__":
    main()
