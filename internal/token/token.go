// Package token contains types and support functions to represent tokens of
// the input language.
package token

import (
	"fmt"
)

type Type int

const (
	Number Type = iota + 1
	Add
	Sub
	Mul
	Div
	LParen
	RParen
)

type Token struct {
	Type  Type
	Value float64
}

func (t Token) String() string {
	switch t.Type {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	case LParen:
		return "("
	case RParen:
		return ")"
	case Number:
		return fmt.Sprintf("%.4f", t.Value)
	default:
		panic(fmt.Sprintf("unknown operator: %v", t.Type))
	}
}

func IsOperator(t Token) bool {
	return t.Type == Add || t.Type == Sub || t.Type == Mul || t.Type == Div
}
