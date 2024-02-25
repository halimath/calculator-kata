package calc

import (
	"strings"
	"testing"

	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestEval(t *testing.T) {
	type testCase struct {
		in   string
		want float64
		err  error
	}

	tests := []testCase{
		{in: "2+3", want: 5},
		{in: "2/3", want: 2.0 / 3.0},
		{in: "2+3+4", want: 9},
		{in: "2+3*4", want: 14},
		{in: "2+3*4-5", want: 9},
		{in: "2+3*(4-5)", want: -1},
		{in: "8 / 2 * (2 + 2)", want: 16},

		{in: "", want: 0, err: ErrEmptyStack},
		{in: "2+", want: 0, err: ErrEmptyStack},
		{in: "2/0", want: 0, err: ErrDivisionByZero},
		{in: "abc", want: 0, err: ErrInvalidInput},
		{in: "2.3.", want: 0, err: ErrInvalidInput},
	}

	for _, test := range tests {
		got, err := Eval(strings.NewReader(test.in))

		expect.WithMessage(t, "in: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}
