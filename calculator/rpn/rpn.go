package rpn

import (
	"errors"
	"fmt"
	"io"

	"github.com/halimath/calc/lexer"
	"github.com/halimath/calc/stack"
	"github.com/halimath/calc/token"
)

type RPN struct {
	l         *lexer.Lexer
	out       stack.Stack[token.Token]
	operators stack.Stack[token.Token]
}

func New(l *lexer.Lexer) *RPN {
	return &RPN{
		l:         l,
		out:       make(stack.Stack[token.Token], 0, 64),
		operators: make(stack.Stack[token.Token], 0, 64),
	}
}

func (rpn *RPN) Next() (token.Token, error) {
	if !rpn.out.Empty() {
		return rpn.out.Shift(), nil
	}

	tok, err := rpn.l.Next()
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return tok, err
		}

		for !rpn.operators.Empty() {
			rpn.out.Push(rpn.operators.Pop())
		}

		if !rpn.out.Empty() {
			return rpn.out.Shift(), nil
		}

		return nil, io.EOF
	}

	if _, ok := tok.(token.Number); ok {
		return tok, nil
	}

	if tok == token.LParen {
		rpn.operators.Push(tok)
		return rpn.Next()
	}

	if tok == token.RParen {
		for {
			if rpn.operators.Empty() {
				return nil, fmt.Errorf("unbalanced parenthesis")
			}

			tok = rpn.operators.Pop()
			if tok == token.LParen {
				break
			}

			rpn.out.Push(tok)
		}

		return rpn.Next()
	}

	if _, ok := tok.(token.Operator); ok {
		for !rpn.operators.Empty() {
			top := rpn.operators.Peek()
			if precedence(top) < precedence(tok) || top == token.LParen {
				break
			}
			rpn.out.Push(rpn.operators.Pop())
		}

		rpn.operators.Push(tok)
	}

	return rpn.Next()
}

func precedence(t token.Token) int {
	switch t {
	case token.Add, token.Sub:
		return 1
	case token.Mul, token.Div:
		return 2
	case token.LParen, token.RParen:
		return 3
	default:
		return 0
	}
}