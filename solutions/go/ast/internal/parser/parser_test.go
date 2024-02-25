package parser

import (
	"strings"
	"testing"

	"github.com/halimath/calc/internal/ast"
	"github.com/halimath/calc/internal/scanner"
	"github.com/halimath/expect"
	"github.com/halimath/expect/is"
)

func TestParser(t *testing.T) {
	type testCase struct {
		in   string
		want ast.Node
		err  error
	}

	tests := []testCase{
		{in: "2", want: ast.Number{Value: "2"}},
		{
			in: "2+3", want: ast.Operator{
				L:  ast.Number{Value: "2"},
				R:  ast.Number{Value: "3"},
				Op: ast.Add,
			},
		},
		{
			in: "2-3", want: ast.Operator{
				L:  ast.Number{Value: "2"},
				R:  ast.Number{Value: "3"},
				Op: ast.Sub,
			},
		},
		{
			in: "2+3-4", want: ast.Operator{
				L: ast.Number{Value: "2"},
				R: ast.Operator{
					L:  ast.Number{Value: "3"},
					R:  ast.Number{Value: "4"},
					Op: ast.Sub,
				},
				Op: ast.Add,
			},
		},
		{
			in: "2*3", want: ast.Operator{
				L:  ast.Number{Value: "2"},
				R:  ast.Number{Value: "3"},
				Op: ast.Mul,
			},
		},
		{
			in: "2*3/4", want: ast.Operator{
				L: ast.Number{Value: "2"},
				R: ast.Operator{
					L:  ast.Number{Value: "3"},
					R:  ast.Number{Value: "4"},
					Op: ast.Div,
				},
				Op: ast.Mul,
			},
		},
		{
			in: "2+3*4", want: ast.Operator{
				L: ast.Number{Value: "2"},
				R: ast.Operator{
					L:  ast.Number{Value: "3"},
					R:  ast.Number{Value: "4"},
					Op: ast.Mul,
				},
				Op: ast.Add,
			},
		},
		{
			in: "2*3+4", want: ast.Operator{
				L: ast.Operator{
					L:  ast.Number{Value: "2"},
					R:  ast.Number{Value: "3"},
					Op: ast.Mul,
				},
				R:  ast.Number{Value: "4"},
				Op: ast.Add,
			},
		},
		{
			in: "(2+3)*4", want: ast.Operator{
				L: ast.Operator{
					L:  ast.Number{Value: "2"},
					R:  ast.Number{Value: "3"},
					Op: ast.Add,
				},
				R:  ast.Number{Value: "4"},
				Op: ast.Mul,
			},
		},
	}

	for _, test := range tests {
		p := New(scanner.New(strings.NewReader(test.in)))
		got, err := p.Expr()

		expect.WithMessage(t, "in: %q", test.in).That(
			is.Error(err, test.err),
			is.DeepEqualTo(got, test.want),
		)
	}
}
