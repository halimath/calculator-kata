package lexer

import (
	"bytes"
	"errors"
	"io"
	"os"
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
			token.Number(2.3),
		}},
		{in: "+", want: []token.Token{
			token.Add,
		}},
		{in: "-", want: []token.Token{
			token.Sub,
		}},
		{in: "*", want: []token.Token{
			token.Mul,
		}},
		{in: "/", want: []token.Token{
			token.Div,
		}},
		{in: "(", want: []token.Token{
			token.LParen,
		}},
		{in: ")", want: []token.Token{
			token.RParen,
		}},
		{in: "2+3*4", want: []token.Token{
			token.Number(2),
			token.Add,
			token.Number(3),
			token.Mul,
			token.Number(4),
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

func BenchmarkLexer(b *testing.B) {
	content, err := os.ReadFile("../testdata/10m")
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := consumeAll(New(bytes.NewReader(content)))
		if err != nil {
			b.Fatal(err)
		}
	}
}
