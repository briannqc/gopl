package main_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	expr "github.com/briannqc/gopl/ch7/sec7.9"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr       string
		env        expr.Env
		wantPretty string
		wantValue  string
	}{
		{
			expr:       "sqrt(A / pi)",
			env:        expr.Env{"A": 87616, "pi": math.Pi},
			wantPretty: "sqrt((A / pi))",
			wantValue:  "167",
		},
		{
			expr:       "pow(x, 3) + pow(y, 3)",
			env:        expr.Env{"x": 12, "y": 1},
			wantPretty: "(pow(x, 3) + pow(y, 3))",
			wantValue:  "1729",
		},
		{
			expr:       "pow(x, 3) + pow(y, 3)",
			env:        expr.Env{"x": 9, "y": 10},
			wantPretty: "(pow(x, 3) + pow(y, 3))",
			wantValue:  "1729",
		},
		{
			expr:       "5 / 9 * (F - 32)",
			env:        expr.Env{"F": -40},
			wantPretty: "((5 / 9) * (F - 32))",
			wantValue:  "-40",
		},
		{
			expr:       "5 / 9 * (F - 32)",
			env:        expr.Env{"F": 32},
			wantPretty: "((5 / 9) * (F - 32))",
			wantValue:  "0",
		},
		{
			expr:       "5 / 9 * (F - 32)",
			env:        expr.Env{"F": 212},
			wantPretty: "((5 / 9) * (F - 32))",
			wantValue:  "100",
		},
		{
			expr:       "5! * 2",
			env:        expr.Env{"F": 212},
			wantPretty: "(5! * 2)",
			wantValue:  "240",
		},
	}
	var prevExpr string
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			if test.expr != prevExpr {
				fmt.Printf("\n%s\n", test.expr)
				prevExpr = test.expr
			}

			expression, err := expr.Parse(test.expr)
			if err != nil {
				t.Error(err)
			}

			got := fmt.Sprintf("%.6g", expression.Eval(test.env))
			fmt.Printf("\t%v => %s\n", test.env, got)
			assert.Equal(t, test.wantValue, got)

			assert.Equal(t, test.wantPretty, expression.String())
			exprReParse, err := expr.Parse(test.wantPretty)
			if err != nil {
				t.Error(err)
			}
			assert.Equal(t, test.wantPretty, exprReParse.String())
		})
	}
}