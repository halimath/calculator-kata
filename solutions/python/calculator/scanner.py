
from dataclasses import dataclass
from enum import Enum
from typing import TextIO


class TokenType (Enum):
    """
    Defines the types of a Token.
    """
    NUMBER = 1
    ADD = 2
    SUB = 3
    MUL = 4
    DIV = 5
    L_PAREN = 6
    R_PAREN = 7


@dataclass
class Token:
    """
    A single token extracted from the input source.
    """
    type: TokenType
    value: str | None = None


_number_chars = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.']


class Scanner:
    """
    Scanner implements a lexical scanner (or lexer) that converts a raw character
    stream to a stream of Tokens.

    Scanner implements the Iterator interface yielding one Token at a time.

    If an error occures during i/o handling, the error is raised during iteration.
    """
    def __init__(self, src: TextIO):
        self._src = src
        self._c = None

    def __iter__(self):
        return self

    def __next__(self):
        number_buf = ''

        while True:
            if self._c is None:
                self._c = self._src.read(1)

            if self._c == '':
                if len(number_buf) > 0:
                    return Token(type=TokenType.NUMBER, value=number_buf)

                raise StopIteration()
            
            if self._c.isspace():
                self._c = None
                continue

            if self._c in _number_chars:
                number_buf += self._c
                self._c = None
                continue

            if self._c in ('+', '-', '*', '/', '(', ')'):
                if len(number_buf) > 0:
                    return Token(type=TokenType.NUMBER, value=number_buf)

                c = self._c
                self._c = None

                match c:
                    case '+':
                        return Token(type=TokenType.ADD)
                    case '-':
                        return Token(type=TokenType.SUB)
                    case '*':
                        return Token(type=TokenType.MUL)
                    case '/':
                        return Token(type=TokenType.DIV)
                    case '(':
                        return Token(type=TokenType.L_PAREN)
                    case ')':
                        return Token(type=TokenType.R_PAREN)

            raise ValueError(f"Unexpected character: {self._c}")
