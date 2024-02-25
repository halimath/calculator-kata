package scanner

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/halimath/calc/internal/token"
	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestScanner(t *testing.T) {
	type testCase struct {
		in   string
		want []token.Token
		err  error
	}

	tests := []testCase{
		{in: ""},
		{in: "abc", err: ErrScanFailed},
		{in: "  2.3", want: []token.Token{
			{Type: token.Number, Value: 2.3},
		}},
		{in: "+", want: []token.Token{
			{Type: token.Add},
		}},
		{in: "-", want: []token.Token{
			{Type: token.Sub},
		}},
		{in: "*", want: []token.Token{
			{Type: token.Mul},
		}},
		{in: "/", want: []token.Token{
			{Type: token.Div},
		}},
		{in: "(", want: []token.Token{
			{Type: token.LParen},
		}},
		{in: ")", want: []token.Token{
			{Type: token.RParen},
		}},
		{in: "2+3*4", want: []token.Token{
			{Type: token.Number, Value: 2},
			{Type: token.Add},
			{Type: token.Number, Value: 3},
			{Type: token.Mul},
			{Type: token.Number, Value: 4},
		}},
		{in: "2 + 3 * 4 ", want: []token.Token{
			{Type: token.Number, Value: 2},
			{Type: token.Add},
			{Type: token.Number, Value: 3},
			{Type: token.Mul},
			{Type: token.Number, Value: 4},
		}},
	}

	for _, test := range tests {
		s := New(strings.NewReader(test.in))

		got, err := consumeAll(s)
		expect.WithMessage(t, "input: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}

func consumeAll(s *Scanner) (toks []token.Token, err error) {
	var t token.Token
	for {
		t, err = s.Next()
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

func BenchmarkScanner(b *testing.B) {
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
