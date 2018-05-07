package expr_test

import (
	"testing"

	"github.com/skillian/expr"
)

func TestEquality(t *testing.T) {
	t.Parallel()
	setupLogger(t)
	var v expr.Int
	if err := (&v).SetValue(expr.Int(1)); err != nil {
		t.Error(err)
	}
	tcs := []struct {
		expr.Expr
		expect interface{}
	}{
		{expr.Eq{expr.Int(1), expr.Int(1)}, true},
		{expr.Eq{expr.Int(1), expr.Int(-1)}, false},
		{expr.Gt{expr.Int(1), expr.Int(-1)}, true},
		{expr.Ge{expr.Int(1), expr.Int(1)}, true},
		{expr.Ge{expr.Int(-1), expr.Int(1)}, false},
		{expr.Eq{expr.Int(1), &v}, true},
	}
	for _, tc := range tcs {
		testHelper(t, tc.Expr, tc.expect)
	}
}

func testHelper(t *testing.T, e expr.Expr, expected interface{}) {
	t.Helper()
	result, err := e.Eval()
	if err != nil {
		t.Errorf("failed to evaluate %v: %v", e, err)
		return
	}
	if result != expected {
		t.Errorf("Result evaluating %v was %v (expected %v)", e, result, expected)
		return
	}
	t.Logf("%v -> %v", e, result)
}
