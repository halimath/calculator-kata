from io import StringIO
from scanner import Scanner, Token, TokenType


def test_scanner():
    s = Scanner(StringIO('( 2 +   3.4   - 5 ) * 6 /  7'))
    want = [
        Token(type=TokenType.L_PAREN),
        Token(type=TokenType.NUMBER, value='2'),
        Token(type=TokenType.ADD),
        Token(type=TokenType.NUMBER, value='3.4'),
        Token(type=TokenType.SUB),
        Token(type=TokenType.NUMBER, value='5'),
        Token(type=TokenType.R_PAREN),
        Token(type=TokenType.MUL),
        Token(type=TokenType.NUMBER, value='6'),
        Token(type=TokenType.DIV),
        Token(type=TokenType.NUMBER, value='7'),
    ]
    got = [t for t in s]
    assert want == got
