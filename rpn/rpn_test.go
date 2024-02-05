package rpn

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/halimath/calc/lexer"
	"github.com/halimath/calc/token"
	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestRPN(t *testing.T) {
	type testCase struct {
		in   string
		want []token.Token
		err  error
	}

	tests := []testCase{
		{in: "2+3", want: tokenize("2 3 +")},
		{in: "2+3+4", want: tokenize("2 3 + 4 +")},
		{in: "2+3*4", want: tokenize("2 3 4 * +")},
		{in: "2+3*4-5", want: tokenize("2 3 4 * + 5 -")},
		{in: "2+3*(4-5)", want: tokenize("2 3 4 5 - * +")},
	}

	for _, test := range tests {
		r := New(lexer.New(strings.NewReader(test.in)))
		got, err := consumeAll(r)

		expect.WithMessage(t, "in: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}

func consumeAll(l *RPN) (toks []token.Token, err error) {
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

func tokenize(s string) (toks []token.Token) {
	l := lexer.New(strings.NewReader(s))

	for {
		t, err := l.Next()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			panic(err)
		}

		toks = append(toks, t)
	}
}
