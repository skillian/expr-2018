package expr_test

import (
	"math/big"
	"testing"

	"github.com/skillian/expr"
)

func TestArithmetic(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		expr.ValueExpr
		checker func(v expr.Value) bool
	}{
		{expr.Add{expr.Int(1), expr.Int(1)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				cmp, err := expr.Cmp(n, expr.Int(2))
				if err != nil {
					t.Error(err)
				}
				return cmp == 0
			}
			return false
		}},
		{expr.Add{expr.Int(1), expr.Int(2)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				cmp, err := expr.Cmp(n, expr.Int(2))
				if err != nil {
					t.Error(err)
				}
				return cmp != 0
			}
			t.Error(v, "is not a Number")
			return false
		}},
		{expr.Sub{expr.Int(1), expr.Int(1)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				cmp, err := expr.Cmp(n, expr.Int(0))
				if err != nil {
					t.Error(err)
				}
				return cmp == 0
			}
			return false
		}},
		{expr.Sub{expr.Int(1), expr.Int(2)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				cmp, err := expr.Cmp(n, expr.Int(-1))
				if err != nil {
					t.Error(err)
				}
				return cmp == 0
			}
			return false
		}},
		{expr.Mul{expr.Int(8), expr.Int(8)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				cmp, err := expr.Cmp(n, expr.Int(64))
				if err != nil {
					t.Error(err)
				}
				return cmp == 0
			}
			return false
		}},
		{expr.Div{expr.Int(1), expr.Int(15)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				expected := (*expr.Rational)(new(big.Rat).SetFrac64(1, 15))
				cmp, err := expr.Cmp(n, expected)
				if err != nil {
					t.Error(err)
				}
				return cmp == 0
			}
			return false
		}},
		{expr.Div{expr.Int(1), expr.Int(15)}, func(v expr.Value) bool {
			if n, ok := v.(expr.Number); ok {
				expected := (*expr.Rational)(new(big.Rat).SetFrac64(1, 25))
				cmp, err := expr.Cmp(n, expected)
				if err != nil {
					t.Error(err)
				}
				return cmp > 0
			}
			return false
		}},
	}
	for _, tc := range tcs {
		arithmeticTestHelper(t, tc.ValueExpr, tc.checker)
	}
}

func arithmeticTestHelper(t *testing.T, e expr.ValueExpr, resultChecker func(v expr.Value) bool) {
	t.Helper()
	result, err := e.EvalValue()
	if err != nil {
		t.Errorf("failed to evaluate %v: %v", e, err)
		return
	}
	if !resultChecker(result) {
		t.Errorf("Result evaluating %v was %v", e, result)
		return
	}
	t.Logf("%v -> %v", e, result)
}
