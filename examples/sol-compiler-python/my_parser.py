import logging
from lark import Lark, logger, __file__ as lark_file, ParseError

logger.setLevel(logging.DEBUG)

from pathlib import Path

grammer_path = Path(__file__).parent
grammer_file = grammer_path / 'grammer.lark'

parser = Lark.open(grammer_file, rel_to=__file__, parser="lalr")


def parse(code):
    try:
        parser.parse(code)
        return True
    except ParseError:
        return False


def debug_main(code):
    print("\n:::PARSER:::")
    print("~~~~~input:")
    print(code)
    tree = parser.parse(code)
    print(tree.pretty())
    print(type(tree))
