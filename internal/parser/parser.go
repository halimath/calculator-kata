package parser

import (
	"errors"
	"fmt"
	"io"

	"github.com/halimath/calc/internal/ast"
	"github.com/halimath/calc/internal/scanner"
	"github.com/halimath/calc/internal/token"
)

var ErrInvalidSyntax = errors.New("invalid syntax")

type Parser struct {
	s       *scanner.Scanner
	current token.Token
}

func New(s *scanner.Scanner) *Parser {
	p := Parser{s: s}
	p.advance()
	return &p
}

func (p *Parser) Expr() (ast.Node, error) {
	n, err := p.term()
	if err != nil {
		return n, err
	}

	if p.current == nil {
		return n, nil
	}

	if p.current != token.Add && p.current != token.Sub {
		return n, nil
	}
	op := ast.Add
	if p.current == token.Sub {
		op = ast.Sub
	}

	opNode := ast.Operator{
		L:  n,
		Op: op,
	}
	p.advance()
	opNode.R, err = p.Expr()
	if err != nil {
		return nil, err
	}

	return opNode, nil
}

func (p *Parser) term() (ast.Node, error) {
	n, err := p.atom()
	if err != nil {
		return nil, err
	}

	if p.current == nil {
		return n, nil
	}

	if p.current != token.Mul && p.current != token.Div {
		return n, nil
	}

	op := ast.Mul
	if p.current == token.Div {
		op = ast.Div
	}

	opNode := ast.Operator{
		L:  n,
		Op: op,
	}
	p.advance()
	opNode.R, err = p.term()
	if err != nil {
		return nil, err
	}

	return opNode, nil
}

func (p *Parser) atom() (ast.Node, error) {
	if v, ok := p.current.(token.Number); ok {
		p.advance()
		return ast.Number{Value: v.String()}, nil
	}

	if p.current == token.LParen {
		p.advance()
		n, err := p.Expr()
		if err != nil {
			return nil, err
		}

		if p.current != token.RParen {
			return nil, fmt.Errorf("%w: expected ) but got %q", ErrInvalidSyntax, p.current)
		}
		p.advance()

		return n, nil
	}

	return nil, fmt.Errorf("%w: unexpected %q", ErrInvalidSyntax, p.current)
}

func (p *Parser) advance() (err error) {
	p.current, err = p.s.Next()

	if errors.Is(err, io.EOF) {
		p.current = nil
		err = nil
	}

	return
}
