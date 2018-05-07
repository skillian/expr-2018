package expr_test

import (
	"testing"

	"github.com/skillian/logging"

	"github.com/skillian/errors"
	"github.com/skillian/expr"
)

type attrgetter struct{}

func (a attrgetter) Copy(mappers ...expr.Mapper) expr.Expr {
	return nil
}

func (a attrgetter) Eval() (interface{}, error) {
	return a.Interface(), nil
}

func (a attrgetter) Interface() interface{} {
	return a
}

func (a attrgetter) EvalValue() (expr.Value, error) {
	return a, nil
}

func (a attrgetter) Type() expr.Type {
	return nil
}

func (a attrgetter) GetAttr(name expr.String) (expr.Value, error) {
	if name == "Name" {
		return expr.String("Sean"), nil
	}
	return nil, errors.Errorf("No attribute %q", name)
}

type dynattrgetter struct {
	Name string
}

func (a dynattrgetter) Copy(mappers ...expr.Mapper) expr.Expr {
	return nil
}

func (a dynattrgetter) Eval() (interface{}, error) {
	return a.Interface(), nil
}

func (a dynattrgetter) Interface() interface{} {
	return a
}

func (a dynattrgetter) EvalValue() (expr.Value, error) {
	return a, nil
}

func (a dynattrgetter) Type() expr.Type {
	return nil
}

var logger = logging.GetLogger("github.com/skillian/expr")

func setupLogger(t *testing.T) {
	t.Helper()
	h := &testingTHandler{logging.BaseHandler{}, t}
	h.SetFormatter(logging.DefaultFormatter{})
	h.SetLevel(logging.DebugLevel)
	logger.AddHandler(h)
	logger.SetLevel(logging.DebugLevel)
}

type testingTHandler struct {
	logging.BaseHandler
	t *testing.T
}

func (t *testingTHandler) Emit(e *logging.Event) {
	msg := t.Formatter().Format(e)
	if e.Level >= logging.WarningLevel {
		t.t.Error(msg)
		return
	}
	t.t.Log(msg)
}

func TestAttr(t *testing.T) {
	t.Parallel()
	setupLogger(t)
	tcs := []struct {
		name string
		expr.ValueExpr
		checker func(v expr.Value) bool
	}{
		{"GetAttr implementation", expr.Attr{ValueExpr: attrgetter{}, Name: "Name"}, func(v expr.Value) bool {
			return v.(expr.String) == "Sean"
		}},
		{"Dynamic GetAttr", expr.Attr{ValueExpr: dynattrgetter{Name: "Ryan"}, Name: "Name"}, func(v expr.Value) bool {
			return v.Interface() == "Ryan"
		}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			value, err := tc.ValueExpr.EvalValue()
			if err != nil {
				t.Error(err)
			}
			if !tc.checker(value) {
				t.Error("result check failed")
			}
			t.Log(tc.ValueExpr, "->", value)
		})
	}
}
