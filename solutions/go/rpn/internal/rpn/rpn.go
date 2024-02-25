// Package rpn provides a type that converts mathematical expressions in infix notation to revese polish
// notation (RPN). It implements the [shunting yard algorithm] as defined by Edsger Dijkstra.
//
// [sunting yard algorithm]: https://en.wikipedia.org/wiki/Shunting_yard_algorithm
package rpn

import (
	"errors"
	"fmt"
	"io"

	"github.com/halimath/calc/internal/scanner"
	"github.com/halimath/calc/internal/stack"
	"github.com/halimath/calc/internal/token"
)

// RPN implements a type to consume token.Token from a scanner.Scanner assuming these to be in infix notation
// and transforms them to reverse plish notation.
type RPN struct {
	s         *scanner.Scanner
	out       stack.Stack[token.Token]
	operators stack.Stack[token.Token]
}

// New creates a new RPN consuming tokens from s.
func New(s *scanner.Scanner) *RPN {
	return &RPN{
		s:         s,
		out:       make(stack.Stack[token.Token], 0, 64),
		operators: make(stack.Stack[token.Token], 0, 64),
	}
}

// Next yields the next token in RPN or an error. If no more tokens are available it returns io.EOF.
func (rpn *RPN) Next() (token.Token, error) {
	if !rpn.out.Empty() {
		return rpn.out.Shift(), nil
	}

	tok, err := rpn.s.Next()
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

		return token.Token{}, io.EOF
	}

	if tok.Type == token.Number {
		return tok, nil
	}

	if tok.Type == token.LParen {
		rpn.operators.Push(tok)
		return rpn.Next()
	}

	if tok.Type == token.RParen {
		for {
			if rpn.operators.Empty() {
				return token.Token{}, fmt.Errorf("unbalanced parenthesis")
			}

			tok = rpn.operators.Pop()
			if tok.Type == token.LParen {
				break
			}

			rpn.out.Push(tok)
		}

		return rpn.Next()
	}

	if token.IsOperator(tok) {
		for !rpn.operators.Empty() {
			top := rpn.operators.Peek()
			if precedence(top) < precedence(tok) || top.Type == token.LParen {
				break
			}
			rpn.out.Push(rpn.operators.Pop())
		}

		rpn.operators.Push(tok)
	}

	return rpn.Next()
}

func precedence(t token.Token) int {
	switch t.Type {
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
