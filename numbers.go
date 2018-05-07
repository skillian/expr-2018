package expr

import (
	"math/big"

	"github.com/skillian/errors"
)

// Number is a specialization of Value for any rational number.
type Number interface {
	Value
	Rat(r *big.Rat)
}

// CmpNumbers is a specialization of the Cmp function that can handle any
// value that implements Number.
func CmpNumbers(left, right Number) (result int) {
	var rats [2]big.Rat
	left.Rat(&rats[0])
	right.Rat(&rats[1])
	return (&rats[0]).Cmp(&rats[1])
}

func numberArithmeticHelper(left, right Number, bigRatFunc func(*big.Rat, *big.Rat, *big.Rat) *big.Rat) Number {
	var rats [2]big.Rat
	left.Rat(&rats[0])
	right.Rat(&rats[1])
	return (*Rational)(bigRatFunc(new(big.Rat), (&rats[0]), (&rats[1])))
}

// Rational wraps any rational numeric value.
type Rational big.Rat

type rationaltype struct{}

// RationalType is the expression Type of a Rational value.
var RationalType Type = rationaltype{}

// RationalZero creates a new Rational zero value.
// This function is called by (new(Rational)).Type().Zero so it's safe to
// overwrite that zero value with subsequent operations.
func RationalZero() *Rational {
	return (*Rational)(big.NewRat(0, 1))
}

// Zero creates a new Zero-value Rational.
func (t rationaltype) Zero() Value {
	return RationalZero()
}

// Var creates a new Rational variable
func (t rationaltype) Var() Var {
	v := new(RationalVar)
	v.value = RationalZero()
	return v
}

// copy is a helper member function to duplicate the underlying *big.Rat and
// convert it to a Rational.
func (r *Rational) copy() *Rational {
	return (*Rational)(new(big.Rat).Set((*big.Rat)(r)))
}

// Cmp implements Cmper
func (r *Rational) Cmp(v Value) (int, error) {
	if n, ok := v.(Number); ok {
		return CmpNumbers(r, n), nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", r, v)
}

// Copy the expression.
func (r *Rational) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(r.copy(), transformations...)
}

// Eval the expression.
func (r *Rational) Eval() (interface{}, error) {
	value, err := r.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the expression to a Value
func (r *Rational) EvalValue() (Value, error) {
	return r.Value(), nil
}

// EvalAdd adds v to r.
func (r *Rational) EvalAdd(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(r, n, (*big.Rat).Add), nil
	}
	return nil, errors.Errorf("cannot add %v to %v", r, v)
}

// EvalSubtract subtracts v from r.
func (r *Rational) EvalSubtract(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(r, n, (*big.Rat).Sub), nil
	}
	return nil, errors.Errorf("cannot subtract %v from %v", r, v)
}

// EvalMultiply multiplies r by v.
func (r *Rational) EvalMultiply(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(r, n, (*big.Rat).Mul), nil
	}
	return nil, errors.Errorf("cannot multiply %v by %v", r, v)
}

// EvalDivide divides r by v.
func (r *Rational) EvalDivide(v Value) (Value, error) {
	if n, ok := v.(Number); ok {
		return numberArithmeticHelper(r, n, (*big.Rat).Quo), nil
	}
	return nil, errors.Errorf("cannot divide %v by %v", r, v)
}

// Interface can return an int64, uint64, float64, or *big.Rat.
func (r *Rational) Interface() interface{} {
	return r.Value().Interface()
}

// Type gets the Rational value's Type: RationalType.
func (r *Rational) Type() Type {
	return RationalType
}

// Value attempts to simplify the Rational number into an Int, Int64, or
// Float64.
//
// While the expr package is performing arithmetic on the objects, they stay as
// Rationals in order to reduce the number of conversions. Once the last
// operation completes, then it makes sense to reduce the result to a simpler
// data type.
func (r *Rational) Value() Value {
	br := (*big.Rat)(r)
	if br.IsInt() {
		n := br.Num()
		if n.IsInt64() {
			return Int64(n.Int64())
		}
		if n.IsUint64() {
			return Uint64(n.Uint64())
		}
	}
	if f, exact := br.Float64(); exact {
		return Float64(f)
	}
	return r
}

// RationalVar is a variable to hold a rational value.
type RationalVar struct {
	value *Rational
}

// Copy copies the RationalVar if the copy of its inner Rational value is still
// a Rational value.  If not, that new inner value is returned directly.
func (r *RationalVar) Copy(transformations ...Mapper) Expr {
	copied := r.value.Copy(transformations...)
	if rat, ok := copied.(*Rational); ok {
		return ApplyMappers(&RationalVar{rat}, transformations...)
	}
	return copied
}

// Eval evaluates the RationalVar
func (r *RationalVar) Eval() (interface{}, error) {
	if r.value == nil {
		return nil, errNilPtr
	}
	return r.value.Eval()
}

// EvalValue evaluates the RationalVar to its Rational value.
func (r *RationalVar) EvalValue() (Value, error) {
	if r.value == nil {
		return nil, errNilPtr
	}
	return r.value.EvalValue()
}

// Value tries to get the Rational as a simpler value (such as an Int64 or
// Uint64) but if it cannot be represented without any loss of precision, it
// simply returns itself.
func (r *RationalVar) Value() Value {
	return r.value
}

// SetValue sets r's value to v.
func (r *RationalVar) SetValue(v Value) error {
	if n, ok := v.(Number); ok {
		var b big.Rat
		n.Rat(&b)
		((*big.Rat)(r.value)).Set(&b)
		return nil
	}
	return errors.Errorf("cannot set %T value to %v", r, v)
}

// Rat stores r's rational value into the math/big.Rat rat.
func (r *Rational) Rat(rat *big.Rat) {
	rat.Set((*big.Rat)(r))
}

// String represents r as a string.
func (r *Rational) String() string {
	return ((*big.Rat)(r)).FloatString(10)
}
