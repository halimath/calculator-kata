from io import StringIO
from calculator.scanner import Scanner, Token, TokenType
from calculator.evaluator import _RPNTransformer, eval_expr

def test_rpn_transformer():    
    rpn = _RPNTransformer(Scanner(StringIO('2+3*(4 + 5)')))
    want = [
        Token(type=TokenType.NUMBER, value='2'),
        Token(type=TokenType.NUMBER, value='3'),
        Token(type=TokenType.NUMBER, value='4'),
        Token(type=TokenType.NUMBER, value='5'),
        Token(type=TokenType.ADD),
        Token(type=TokenType.MUL),
        Token(type=TokenType.ADD),
    ]
    got = [t for t in rpn]
    assert want == got


def test_eval():
    assert 16 == eval_expr(StringIO('8/2*(2+2)'))