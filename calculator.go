// Package calc implements an arithmetic calculator for very long mathematical expressions.
package calc

import (
	"errors"
	"fmt"
	"strconv"

	"io"

	"github.com/halimath/calc/internal/ast"
	"github.com/halimath/calc/internal/parser"
	"github.com/halimath/calc/internal/scanner"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrDivisionByZero = errors.New("division by zero")
)

// Eval evaluates the expression read from r and returns the result as well as any error.
func Eval(r io.Reader) (float64, error) {
	node, err := parser.New(scanner.New(r)).Expr()
	if err != nil {
		return 0, fmt.Errorf("%w: parsing error: %v", ErrInvalidInput, err)
	}

	return eval(node)
}

func eval(node ast.Node) (float64, error) {
	switch n := node.(type) {
	case ast.Number:
		v, err := strconv.ParseFloat(n.Value, 64)
		if err != nil {
			return 0, fmt.Errorf("%w: %v", ErrInvalidInput, err)
		}
		return v, nil

	case ast.Operator:
		l, err := eval(n.L)
		if err != nil {
			return 0, err
		}

		r, err := eval(n.R)
		if err != nil {
			return 0, err
		}

		switch n.Op {
		case ast.Add:
			return l + r, nil
		case ast.Sub:
			return l - r, nil
		case ast.Mul:
			return l * r, nil
		case ast.Div:
			if r == 0 {
				return 0, ErrDivisionByZero
			}
			return l / r, nil
		default:
			return 0, fmt.Errorf("%w: unexpected operator: %v", ErrInvalidInput, n.Op)
		}
	default:
		return 0, fmt.Errorf("%w: unexpected ast node: %v", ErrInvalidInput, node)
	}

}
