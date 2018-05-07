package expr

import (
	"math/big"

	"github.com/skillian/errors"
)

// Int wraps a Go int value.
type Int int

type inttype struct{}

// IntType is the expression Type of Ints.
var IntType Type = inttype{}

func (t inttype) Zero() Value {
	return Int(0)
}

func (t inttype) Var() Var {
	return new(Int)
}

// Copy the expression.
func (i Int) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(i, transformations...)
}

// Cmp compares the Int to another Number. If v is not a Number, an error
// is returned.
func (i Int) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(i, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", i, v)
}

// Eval the expression.
func (i Int) Eval() (interface{}, error) {
	return i.Interface(), nil
}

// EvalAdd adds the given number to the current Int.  If v is not a number,
// an error is returned.
func (i Int) EvalAdd(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Add), nil
	}
	return nil, errors.Errorf("cannot add %v to %v", i, v)
}

// EvalSubtract subtracts v from i.
func (i Int) EvalSubtract(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Sub), nil
	}
	return nil, errors.Errorf("cannot subtract %v from %v", i, v)
}

// EvalMultiply multiplies i by v.
func (i Int) EvalMultiply(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Mul), nil
	}
	return nil, errors.Errorf("cannot multiply %v by %v", i, v)
}

// EvalDivide divides i by v.
func (i Int) EvalDivide(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Quo), nil
	}
	return nil, errors.Errorf("cannot divide %v by %v", i, v)
}

// Interface of the expression.
func (i Int) Interface() interface{} {
	return int(i)
}

// EvalValue evaluates the expression to a a value.
func (i Int) EvalValue() (Value, error) {
	return i, nil
}

// Rat stores the value of i into r.
func (i Int) Rat(r *big.Rat) {
	r.SetInt64(int64(int(i)))
}

// Type gets the type of Int i: IntType.
func (i Int) Type() Type {
	return IntType
}

// Value of the expression.
func (i *Int) Value() Value {
	return *i
}

// SetValue of the expression.
func (i *Int) SetValue(v Value) error {
	if v, ok := v.(Int); ok {
		*i = v
		return nil
	}
	return TypeErrorFromExpectedAndActual(i, v)
}

// Int64 wraps a Go int64 value.
type Int64 int64

type int64type struct{}

// Int64Type is the expr.Type of an Int64.
var Int64Type Type = int64type{}

func (t int64type) Zero() Value {
	return Int64(0)
}

func (t int64type) Var() Var {
	return new(Int64)
}

// Copy the expression.
func (i Int64) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(i, transformations...)
}

// Cmp compares i to v.
func (i Int64) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(i, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", i, v)
}

// Eval the expression.
func (i Int64) Eval() (interface{}, error) {
	return i.Interface(), nil
}

// Interface of the expression.
func (i Int64) Interface() interface{} {
	return int(i)
}

// EvalValue evaluates the expression to a a value.
func (i Int64) EvalValue() (Value, error) {
	return i, nil
}

// EvalAdd adds v to i.
func (i Int64) EvalAdd(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Add), nil
	}
	return nil, errors.Errorf("cannot add %v to %v", i, v)
}

// EvalSubtract subtracts v from i.
func (i Int64) EvalSubtract(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Sub), nil
	}
	return nil, errors.Errorf("cannot subtract %v from %v", i, v)
}

// EvalMultiply multiplies i by v.
func (i Int64) EvalMultiply(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Mul), nil
	}
	return nil, errors.Errorf("cannot multiply %v by %v", i, v)
}

// EvalDivide divides i by v.
func (i Int64) EvalDivide(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Quo), nil
	}
	return nil, errors.Errorf("cannot divide %v by %v", i, v)
}

// Rat stores the value of i into r.
func (i Int64) Rat(r *big.Rat) {
	r.SetInt64(int64(i))
}

// Type gets the expression type of i: Int64Type.
func (i Int64) Type() Type {
	return Int64Type
}

// Value of the expression.
func (i *Int64) Value() Value {
	return *i
}

// SetValue of the expression.
func (i *Int64) SetValue(v Value) error {
	switch v := v.(type) {
	case Int64:
		*i = v
	case Int:
		*i = Int64(int64(int(v)))
	default:
		return TypeErrorFromExpectedAndActual(i, v)
	}
	return nil
}

// Uint64 wraps a Go uint64 value.
type Uint64 uint64

type uint64type struct{}

// Uint64Type defines the expression type of a Uint64.
var Uint64Type Type = uint64type{}

func (t uint64type) Zero() Value {
	return Uint64(0)
}

func (t uint64type) Var() Var {
	return new(Uint64)
}

// Copy the expression.
func (i Uint64) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(i, transformations...)
}

// Cmp compares i to v.
func (i Uint64) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(i, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", i, v)
}

// Eval the expression.
func (i Uint64) Eval() (interface{}, error) {
	return i.Interface(), nil
}

// Interface of the expression.
func (i Uint64) Interface() interface{} {
	return int(i)
}

// EvalValue evaluates the expression to a a value.
func (i Uint64) EvalValue() (Value, error) {
	return i, nil
}

// EvalAdd adds v to i.
func (i Uint64) EvalAdd(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Add), nil
	}
	return nil, errors.Errorf("cannot add %v to %v", i, v)
}

// EvalSubtract subtracts v from i.
func (i Uint64) EvalSubtract(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Sub), nil
	}
	return nil, errors.Errorf("cannot subtract %v from %v", i, v)
}

// EvalMultiply multiplies i by v.
func (i Uint64) EvalMultiply(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Mul), nil
	}
	return nil, errors.Errorf("cannot multiply %v by %v", i, v)
}

// EvalDivide divides i by v.
func (i Uint64) EvalDivide(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(i, n, (*big.Rat).Quo), nil
	}
	return nil, errors.Errorf("cannot divide %v by %v", i, v)
}

// Rat stores i's value into r.
func (i Uint64) Rat(r *big.Rat) {
	var bi big.Int
	bi.SetUint64(uint64(i))
	r.SetInt(&bi)
}

// Type gets the expression type of i: Uint64Type.
func (i Uint64) Type() Type {
	return Uint64Type
}

// Value of the expression.
func (i *Uint64) Value() Value {
	return *i
}

// SetValue of the expression.
func (i *Uint64) SetValue(v Value) error {
	switch v := v.(type) {
	case Uint64:
		*i = v
	case Int:
		if v < 0 {
			return errors.Errorf("cannot set Uint64 from negative integer: %v", v)
		}
		*i = Uint64(uint64(int(v)))
	case Int64:

	default:
		return TypeErrorFromExpectedAndActual(i, v)
	}
	return nil
}
