// Package token contains types and support functions to represent tokens of
// the input language.
package token

import (
	"fmt"
)

// Token defines the interface for tokens. Only types contained in this package
// may satisfy the Token interface.
type Token interface {
	String() string
	tok()
}

// Operator defines a type of Token that represents a binary operator in the
// input language.
type Operator int

const (
	Add Operator = iota + 1
	Sub
	Mul
	Div
)

func (Operator) tok() {}

func (o Operator) String() string {
	switch o {
	case Add:
		return "+"
	case Sub:
		return "-"
	case Mul:
		return "*"
	case Div:
		return "/"
	default:
		panic(fmt.Sprintf("unknown operator: %d", int(o)))
	}
}

// Paren defines a type of Token that represents either a left of right parenthesis.
type Paren int

const (
	LParen Paren = iota + 1
	RParen
)

func (Paren) tok() {}

func (p Paren) String() string {
	switch p {
	case LParen:
		return "("
	case RParen:
		return ")"
	default:
		panic(fmt.Sprintf("unknown paren: %d", int(p)))
	}
}

// Number defines a type of Token that represents a number literal.
type Number string

func (Number) tok() {}

func (n Number) String() string { return string(n) }
