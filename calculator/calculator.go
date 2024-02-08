package calculator

import (
	"errors"
	"fmt"
	"io"

	"github.com/halimath/calc/calculator/rpn"
	"github.com/halimath/calc/lexer"
	"github.com/halimath/calc/stack"
	"github.com/halimath/calc/token"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrEmptyStack     = errors.New("empty stack")
	ErrDivisionByZero = errors.New("division by zero")
)

func Eval(r io.Reader) (float64, error) {
	operands := make(stack.Stack[float64], 0, 64)

	rpnConverver := rpn.New(lexer.New(r))

	for {
		tok, err := rpnConverver.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return 0, fmt.Errorf("%w: %v", ErrInvalidInput, err)
		}

		if val, ok := tok.(token.Number); ok {
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
