package expr

import (
	"reflect"

	"github.com/skillian/errors"
)

// A Set is a collection of ValueExprs that can be treated as a single
// ValueExpr.
type Set []ValueExpr

type settype struct {
	types []Type
}

// Zero actually creates a ValueSet with the Zero values of each of its
// expressions. Set itself holds arbitray expressions, so a Zero value doesn't
// really make sense to me.
func (t settype) Zero() Value {
	logger.Warn(
		"Unexpected use of %T. If you get an error from using this Zero "+
			"value, let skillian know.", t)
	s := make(ValueSet, len(t.types))
	for i, typ := range t.types {
		s[i] = typ.Zero()
	}
	return s
}

// Var is supposed to create a Set variable.  I can't think of the right way for
// this to work, so this just panics.  Let me know if you have a use case.
func (t settype) Var() Var {
	panic(
		errors.Errorf(
			"Set cannot yet be made into a Var.  Let skillian know that you " +
				"got this error message and how. Hopefully after I see some " +
				"use cases, I can come up with an expected way for this to " +
				"operate."))
}

// Copy the Set and the inner expressions.
func (s Set) Copy(transformations ...Mapper) Expr {
	s2 := make(Set, len(s))
	for i, e := range s {
		s2[i] = e.Copy(transformations...).(ValueExpr)
	}
	return ApplyMappers(s2, transformations...)
}

// EvalValue evaluates each of the expressions in the set and stores the results
// in a ValueSet and returns it.
func (s Set) EvalValue() (result Value, err error) {
	v := make(ValueSet, len(s))
	for i, valueExpr := range s {
		v[i], err = valueExpr.EvalValue()
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}

// Eval evaluates the Set into a ValueSet and then evaluates that ValueSet
func (s Set) Eval() (result interface{}, err error) {
	value, err := s.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Eval()
}

// Type gets the expression's Type.
func (s Set) Type() Type {
	return settype{types: TypesOf(s)}
}

// ValueSet can be treated as an actual Value (all of its inner expressions)
// are Values).
type ValueSet []Value

// Copy the ValueSet
func (s ValueSet) Copy(transformations ...Mapper) Expr {
	s2 := make(ValueSet, len(s))
	for i, e := range s {
		s2[i] = e.Copy(transformations...).(Value)
	}
	return ApplyMappers(s2, transformations...)
}

// EvalValue returns the same set.
func (s ValueSet) EvalValue() (Value, error) {
	return s, nil
}

// Eval evaluates the ValueSet into a []interface{}
func (s ValueSet) Eval() (result interface{}, err error) {
	results := make([]interface{}, len(s))
	for i, value := range s {
		results[i], err = value.Eval()
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

// Interface gets the set of values as a []interface{}
func (s ValueSet) Interface() interface{} {
	results := make([]interface{}, len(s))
	for i, value := range s {
		results[i] = value.Interface()
	}
	return results
}

// Cmp allows sets to be compared to one another.
func (s ValueSet) Cmp(v Value) (int, error) {
	if s2, ok := v.(ValueSet); ok {
		len1 := len(s)
		len2 := len(s2)
		minlen := Min(len1, len2)
		for i, value := range s[:minlen] {
			cmp, err := Cmp(value, s2[i])
			if err != nil {
				return 0, errors.ErrorfWithCause(
					err,
					"Failed to compare element %v to %v at index %d",
					value, s2[i], i)
			}
			if cmp != 0 {
				return cmp, nil
			}
		}
		// Every element compared equal
		return 0, nil
	}
	return 0, errors.Errorf("cannot compare %v to %v", s, v)
}

// Type gets the type of this set.
func (s ValueSet) Type() Type {
	values := make([]ValueExpr, len(s))
	for i, value := range s {
		values[i] = value
	}
	return settype{types: TypesOf(values)}
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// TypesOf Takes a slice of values and gets their expression Types.
func TypesOf(values []ValueExpr) []Type {
	types := make([]Type, len(values))
	for i, value := range values {
		if tve, ok := value.(TypedValueExpr); ok {
			types[i] = tve.Type()
		} else {
			types[i] = dynamictype{Type: reflect.TypeOf(value)}
		}
	}
	return types
}
