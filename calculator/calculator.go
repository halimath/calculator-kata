package calculator

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/halimath/calc/lexer"
	"github.com/halimath/calc/rpn"
	"github.com/halimath/calc/stack"
	"github.com/halimath/calc/token"
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
			return 0, err
		}

		if tok.Type == token.Number {
			val, err := strconv.ParseFloat(tok.Value, 64)
			if err != nil {
				return 0, err
			}
			operands.Push(val)
			continue
		}

		if token.IsOperator(tok) {
			if len(operands) < 2 {
				return 0, fmt.Errorf("empty stack")
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
					return 0, fmt.Errorf("division by zero")
				}
				operands.Push(r / l)
			}

			continue
		}

		return 0, fmt.Errorf("unexpected token: %v", tok)
	}

	if operands.Empty() {
		return 0, fmt.Errorf("empty stack")
	}

	return operands.Pop(), nil
}
