package expr

import "fmt"

// String wraps a Go string value.
type String string

type stringtype struct{}

var StringType Type = stringtype{}

func (t stringtype) Zero() Value {
	return String("")
}

func (t stringtype) Var() Var {
	return new(String)
}

// Copy the expression.
func (s String) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(s, transformations...)
}

// Eval the expression.
func (s String) Eval() (interface{}, error) {
	return s.Interface(), nil
}

// Interface of the expression.
func (s String) Interface() interface{} {
	return string(s)
}

func (s String) EvalValue() (Value, error) {
	return s, nil
}

func (s String) Type() Type {
	return StringType
}

// Value of the expression.
func (s *String) Value() Value {
	return *s
}

// SetValue of the expression.
func (s *String) SetValue(v Value) error {
	if v, ok := v.(String); ok {
		*s = v
		return nil
	}
	return TypeErrorFromExpectedAndActual(s, v)
}

// StringifyExpr represents the expression as a string. If the expression is
// not const or an inner expression (inner is true) then the stringified
// expression is wrapped in parentheses.
func StringifyExpr(e Expr, inner bool) (result string) {
	switch e := e.(type) {
	case fmt.Stringer:
		result = e.String()
	default:
		result = fmt.Sprint(e)
	}
	if !inner || IsConst(e) {
		return result
	}
	return "(" + result + ")"
}

func stringifyBinaryInfixHelper(b Binary, infix string) string {
	return fmt.Sprintf(
		"%v %v %v",
		StringifyExpr(b.Left(), true),
		infix,
		StringifyExpr(b.Right(), true))
}
