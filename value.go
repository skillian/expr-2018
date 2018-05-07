package expr

import (
	"reflect"

	"github.com/skillian/errors"
)

// ValueExpr represents an expression that evaluates to a value.
type ValueExpr interface {
	Expr
	EvalValue() (Value, error)
}

// TypedValueExpr is a ValueExpr that has a Type method to get its expression
// Type if it were evaluated.
type TypedValueExpr interface {
	ValueExpr
	Type() Type
}

// Value is implemented by any type that represents a value (as opposed to
// a nested expression).
type Value interface {
	ValueExpr
	// Interface "unpacks" the Value as an interface{} just like
	// reflect.Value.Interface.
	Interface() interface{}
}

// BoolExpr represents a value expression with a boolean result.
type BoolExpr interface {
	ValueExpr
	EvalBool() (bool, error)
}

// Var is a variable whose value can be set.
type Var interface {
	Value() Value
	SetValue(v Value) error
}

// VarExpr combines the Var and Expr interfaces.
type VarExpr interface {
	Var
	ValueExpr
}

// ValueOf works similiarly to reflect.ValueOf except there are Value
// specializations for some of Go's built in types.  If there is no
// specialization, a fallback dynamic value type is returned.
func ValueOf(v interface{}) Value {
	switch v := v.(type) {
	case Value:
		return v.Copy().(Value)

	case bool:
		return Bool(v)
	case *bool:
		b := new(Bool)
		*b = Bool(*v)
		return b

	case float32:
		return Float32(v)
	case *float32:
		f := new(Float32)
		*f = Float32(*v)
		return f

	case float64:
		return Float64(v)
	case *float64:
		f := new(Float64)
		*f = Float64(*v)
		return f

	case int:
		return Int(v)
	case *int:
		i := new(Int)
		*i = Int(*v)
		return i

	default:
		return Dynamic(reflect.ValueOf(v))
	}
}

// IsConst returns true if the expression implements the Value interface but
// does not implement the Var interface.
func IsConst(e Expr) bool {
	_, isval := e.(Value)
	_, isvar := e.(Var)
	return isval && !isvar
}

// ToValues simplifies expressions into values. If any expression doesn't
// simplify into an expression, an error is returned.
func ToValues(exprs ...Expr) ([]Value, error) {
	values := make([]Value, len(exprs))
	err := SimplifyIntoValues(exprs, values)
	return values, err
}

// SimplifyIntoValues attempts to simplify all of the expressions into Values.
// exprs are simplified in place with their variables evaluated.
// Any expression that cannot be simplified to a value results in an error
// being returned.
func SimplifyIntoValues(exprs []Expr, values []Value) (err error) {
	if len(values) < len(exprs) {
		return errors.Errorf("values slice smaller than exprs slice")
	}
	exprs = SimplifyAll(exprs, EvaluateVariable)
	for i, expr := range exprs {
		valueexpr, ok := expr.(ValueExpr)
		if !ok {
			return errors.Errorf(
				"Simplified expression must be a ValueExpr in order to be "+
					"simplified into a value (was %T)", expr)
		}
		values[i], err = valueexpr.EvalValue()
		if err != nil {
			return err
		}
	}
	return nil
}

// Type represents the type of an expression.
type Type interface {
	// Zero returns the type's Zero value.
	Zero() Value

	// Var creates a variable of the type.
	Var() Var
}
