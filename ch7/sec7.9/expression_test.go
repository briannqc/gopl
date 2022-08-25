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
		expr string
		env  expr.Env
		want string
	}{
		{"sqrt(A / pi)", expr.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", expr.Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", expr.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", expr.Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", expr.Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", expr.Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}

		expression, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}

		got := fmt.Sprintf("%.6g", expression.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		assert.Equal(t, test.want, got)
	}
}
