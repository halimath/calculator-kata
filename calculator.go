// Package calc implements an arithmetic calculator for very long mathematical expressions.
package calc

import (
	"errors"
	"fmt"
	"io"
	"strconv"

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

		if val, ok := tok.(token.Number); ok {
			val, err := strconv.ParseFloat(string(val), 64)
			if err != nil {
				return 0, fmt.Errorf("%w: %v", ErrInvalidInput, err)
			}

			operands.Push(float64(val))
			continue
		}

		if _, ok := tok.(token.Operator); ok {
			if len(operands) < 2 {
				return 0, ErrEmptyStack
			}

			l := operands.Pop()
			r := operands.Pop()

			switch tok {
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