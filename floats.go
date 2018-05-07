package expr

import (
	"math/big"

	"github.com/skillian/errors"
)

// Float32 wraps a Go float32 value.
type Float32 float32

type float32type struct{}

// Float32Type is the expr.Type object of Float32 values.
var Float32Type Type = float32type{}

func (t float32type) Zero() Value {
	return Float32(0)
}

func (t float32type) Var() Var {
	return new(Float32)
}

// Copy the expression.
func (f Float32) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(f, transformations...)
}

// Cmp compares the Float32 to another number. If v is not a number, an error
// is returned.
func (f Float32) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(f, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", f, v)
}

// Eval the expression.
func (f Float32) Eval() (interface{}, error) {
	return f.Interface(), nil
}

// Interface of the expression.
func (f Float32) Interface() interface{} {
	return float32(f)
}

// EvalValue evaluates the expression to a a value.
func (f Float32) EvalValue() (Value, error) {
	return f, nil
}

// Type gets the Float32's expr.Type: Float32Type.
func (f Float32) Type() Type {
	return Float32Type
}

// Value of the expression.
func (f *Float32) Value() Value {
	return *f
}

// SetValue of the expression.
func (f *Float32) SetValue(v Value) error {
	if v, ok := v.(Float32); ok {
		*f = v
		return nil
	}
	return TypeErrorFromExpectedAndActual(f, v)
}

// Rat sets the *big.Rat from this Float32.
func (f Float32) Rat(r *big.Rat) {
	r.SetFloat64(float64(float32(f)))
}

// Float64 wraps a Go float64 value.
type Float64 float64

type float64type struct{}

// Float64Type is the expr.Type object of Float64 values.
var Float64Type Type = float64type{}

func (t float64type) Zero() Value {
	return Float64(0)
}

func (t float64type) Var() Var {
	return new(Float64)
}

// Copy the expression.
func (f Float64) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(f, transformations...)
}

// Cmp compares the Float64 to another number.  If v is not a Number, the
// comparison fails.
func (f Float64) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(f, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", f, v)
}

// Eval the expression.
func (f Float64) Eval() (interface{}, error) {
	return f.Value(), nil
}

// Interface of the expression.
func (f Float64) Interface() interface{} {
	return float32(f)
}

// EvalValue evaluates the expression to a a value.
func (f Float64) EvalValue() (Value, error) {
	return f, nil
}

// Type gets the Float64's expr.Type: Float64Type.
func (f Float64) Type() Type {
	return Float64Type
}

// Value of the expression.
func (f *Float64) Value() Value {
	return *f
}

// SetValue of the expression.
func (f *Float64) SetValue(v Value) error {
	switch v := v.(type) {
	case Float64:
		*f = v
	case Float32:
		*f = Float64(float64(float32(v)))
	default:
		return TypeErrorFromExpectedAndActual(f, v)
	}
	return nil
}

// Rat sets the *big.Rat from this Float64.
func (f Float64) Rat(r *big.Rat) {
	r.SetFloat64(float64(f))
}
