package calculator

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
		{in: "2+3+4", want: 9},
		{in: "2+3*4", want: 14},
		{in: "2+3*4-5", want: 9},
		{in: "2+3*(4-5)", want: -1},
	}

	for _, test := range tests {
		got, err := Eval(strings.NewReader(test.in))

		expect.WithMessage(t, "in: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}
