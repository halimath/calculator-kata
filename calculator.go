// Package calc implements an arithmetic calculator for very long mathematical expressions.
package calc

import (
	"errors"
	"fmt"
	"io"

	"github.com/halimath/calc/internal/rpn"
	"github.com/halimath/calc/internal/scanner"
	"github.com/halimath/calc/internal/stack"
	"github.com/halimath/calc/internal/token"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrEmptyStack     = errors.New("empty stack")
	ErrDivisionByZero = errors.New("division by zero")
)

// Eval evaluates the expression read from r and returns the result as well as any error.
func Eval(r io.Reader) (float64, error) {
	operands := make(stack.Stack[float64], 0, 64)

	rpnConverver := rpn.New(scanner.New(r))

	for {
		tok, err := rpnConverver.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, fmt.Errorf("%w: %v", ErrInvalidInput, err)
		}

		if tok.Type == token.Number {
			operands.Push(tok.Value)
			continue
		}

		if token.IsOperator(tok) {
			if len(operands) < 2 {
				return 0, ErrEmptyStack
			}

			l := operands.Pop()
			r := operands.Pop()

			switch tok.Type {
			case token.Add:
				operands.Push(r + l)
			case token.Sub:
				operands.Push(r - l)
			case token.Mul:
				operands.Push(r * l)
			case token.Div:
				if l == 0 {
					return 0, ErrDivisionByZero
				}
				operands.Push(r / l)
			}

			continue
		}

		return 0, fmt.Errorf("%w: unexpected token: %v", ErrInvalidInput, tok)
	}

	if operands.Empty() {
		return 0, ErrEmptyStack
	}

	return operands.Pop(), nil
}
