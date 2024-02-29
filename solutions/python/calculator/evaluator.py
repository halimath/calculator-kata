
from array import array
from collections.abc import Iterable
from typing import TextIO

from .scanner import Token, TokenType, Scanner


def eval_expr(src: TextIO) -> float:
    operands = array('d')
    transformer = _RPNTransformer(Scanner(src))
    for t in transformer:
        match t.type:
            case TokenType.NUMBER:
                operands.append(float(t.value))
            case TokenType.ADD:
                r = operands.pop()
                l = operands.pop()
                operands.append(l+r)
            case TokenType.SUB:
                r = operands.pop()
                l = operands.pop()
                operands.append(l-r)
            case TokenType.MUL:
                r = operands.pop()
                l = operands.pop()
                operands.append(l*r)
            case TokenType.DIV:
                r = operands.pop()
                l = operands.pop()
                operands.append(l/r)
    return operands.pop()


class _RPNTransformer:
    def __init__(self, src: Iterable[Token]):
        self._src = src
        self._out = []
        self._operators = []

    def __iter__(self):
        return self

    def __next__(self):
        if len(self._out) > 0:
            return self._out.pop(0)

        try:
            t = next(self._src)

            if t.type == TokenType.NUMBER:
                return t
            
            if t.type == TokenType.L_PAREN:
                self._operators.append(t)
                return next(self)
            
            if t.type == TokenType.R_PAREN:
                while len(self._operators) > 0 and self._operators[-1].type != TokenType.L_PAREN:
                    self._out.append(self._operators.pop())

                if len(self._operators) == 0:
                    raise ValueError('missing (')
                
                # Pop the left paren
                self._operators.pop()

                return next(self)
            
            while len(self._operators) > 0 and _precedence(self._operators[-1]) >= _precedence(t):
                self._out.append(self._operators.pop())
                
            self._operators.append(t)

        except StopIteration:
            if len(self._operators) == 0:
                raise StopIteration()
            
            while len(self._operators) > 0:
                self._out.append(self._operators.pop())
            
        return next(self)


def _precedence(t: Token) -> int:
    match t.type:
        case TokenType.ADD | TokenType.SUB:
            return 1
        case TokenType.MUL | TokenType.DIV:
            return 2
        case _:
            return 0
