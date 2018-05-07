package expr

import "reflect"

const (
	// False wraps the Go bool false value into an expr.Bool
	False = Bool(false)

	// True wraps the Go bool true value into an expr.Bool
	True = Bool(true)
)

// Truthy returns true if the given value is not the zero value of its type
func Truthy(v interface{}) bool {
	if v == nil {
		return false
	}
	switch v := v.(type) {
	case bool:
		return v
	case float32:
		return v != 0
	case float64:
		return v != 0
	case int:
		return v != 0
	case string:
		return v != ""
	case Value:
		return IsConst(v) && Truthy(v.Interface())
	default:
		return reflect.Zero(reflect.TypeOf(v)).Interface() != v
	}
}

// Bool wraps a Go bool value.
type Bool bool

type booltype struct{}

// BoolType represents the Bool expression type.
var BoolType Type = booltype{}

// Zero value of the Bool type.
func (t booltype) Zero() Value {
	return False
}

// Var creates a Bool variable
func (t booltype) Var() Var {
	return new(Bool)
}

// Copy the expression.
func (b Bool) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(b, transformations...)
}

// Eval the expression.
func (b Bool) Eval() (interface{}, error) {
	return b.Interface(), nil
}

// EvalValue evaluates the expression to a value.
func (b Bool) EvalValue() (Value, error) {
	return b, nil
}

// Type of the expression
func (b Bool) Type() Type {
	return BoolType
}

// Interface of the expression.
func (b Bool) Interface() interface{} {
	return bool(b)
}

// Value gets the value of the expression.
func (b *Bool) Value() Value {
	return *b
}

// SetValue sets the value of the expression.
func (b *Bool) SetValue(v Value) error {
	if v, ok := v.(Bool); ok {
		*b = v
		return nil
	}
	return TypeErrorFromExpectedAndActual(b, v)
}

// String gets the string representation of this bool.
func (b Bool) String() string {
	if bool(b) {
		return "true"
	}
	return "false"
}

// GoString output can be used in a .go source file
func (b Bool) GoString() string {
	return "Bool(" + b.String() + ")"
}
