package expr

import (
	"reflect"
)

// GetAttrValue gets an attribute from the given value.
//
// An attribute can either be an exported struct field, a function with one
// return value or a function with two return values where the second is an
// error.  If that error's result is not nil, GetAttrValue returns nil and that
// error.  If the value implements AttrGetter, GetAttrValue just uses the
// value's GetAttr function to get the value.
func GetAttrValue(value Value, name String) (Value, error) {
	if ag, ok := value.(AttrGetter); ok {
		return ag.GetAttr(name)
	}
	return Dynamic(reflect.ValueOf(value)).GetAttr(name)
}

// Attr gets an attribute of a value (e.g. value.name).
type Attr struct {
	// ValueExpr is the value from which the attribute is retrieved.
	ValueExpr
	// Name is the name of the attribute to get from the value.
	Name String
}

// AttrGetter is implemented by any type that can pull attributes from itself
// by name.
type AttrGetter interface {
	GetAttr(name String) (Value, error)
}

// AttrSetter is implemented by any type that can have its attributes assigned.
type AttrSetter interface {
	SetAttr(name String, value Value) error
}

// Copy the expression.
func (a Attr) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Attr{
		a.ValueExpr.Copy(transformations...).(ValueExpr),
		a.Name.Copy(transformations...).(String),
	}, transformations...)
}

// Eval evaluates the expression.
func (a Attr) Eval() (interface{}, error) {
	value, err := a.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the Attr expression into a Value
func (a Attr) EvalValue() (Value, error) {
	value, err := a.ValueExpr.EvalValue()
	if err != nil {
		return nil, err
	}
	return GetAttrValue(value, a.Name)
}
