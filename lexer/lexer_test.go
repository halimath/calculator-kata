package lexer

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/halimath/calc/token"
	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestLexer(t *testing.T) {
	type testCase struct {
		in   string
		want []token.Token
		err  error
	}

	tests := []testCase{
		{in: ""},
		{in: "abc", err: ErrLexer},
		{in: "  2.3", want: []token.Token{
			{Type: token.Number, Value: "2.3"},
		}},
		{in: "+", want: []token.Token{
			{Type: token.Add, Value: "+"},
		}},
		{in: "-", want: []token.Token{
			{Type: token.Sub, Value: "-"},
		}},
		{in: "*", want: []token.Token{
			{Type: token.Mul, Value: "*"},
		}},
		{in: "/", want: []token.Token{
			{Type: token.Div, Value: "/"},
		}},
		{in: "(", want: []token.Token{
			{Type: token.LParen, Value: "("},
		}},
		{in: ")", want: []token.Token{
			{Type: token.RParen, Value: ")"},
		}},
		{in: "2+3*4", want: []token.Token{
			{Type: token.Number, Value: "2"},
			{Type: token.Add, Value: "+"},
			{Type: token.Number, Value: "3"},
			{Type: token.Mul, Value: "*"},
			{Type: token.Number, Value: "4"},
		}},
	}

	for _, test := range tests {
		l := New(strings.NewReader(test.in))

		got, err := consumeAll(l)
		expect.WithMessage(t, "input: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}

func consumeAll(l *Lexer) (toks []token.Token, err error) {
	var t token.Token
	for {
		t, err = l.Next()
		if errors.Is(err, io.EOF) {
			err = nil
			return
		}

		if err != nil {
			return
		}

		toks = append(toks, t)
	}
}
