package expr

import (
	"reflect"

	"github.com/skillian/errors"
)

// Dynamic values can hold any Go value (it wraps reflect.Value)
type Dynamic reflect.Value

type dynamictype struct {
	reflect.Type
}

func (t dynamictype) Zero() Value {
	return Dynamic(reflect.Zero(t.Type))
}

func (t dynamictype) Var() Var {
	return new(Dynamic)
}

func (d Dynamic) copy() Dynamic {
	v := reflect.Value(d)
	var r reflect.Value
	r.Set(v)
	return Dynamic(r)
}

// Copy this Dynamic value
func (d Dynamic) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(d.copy(), transformations...)
}

// Eval unpacks the Dynamic value into an interface{}
func (d Dynamic) Eval() (interface{}, error) {
	v := reflect.Value(d)
	if !v.CanInterface() {
		return nil, errors.Errorf("Dynamic value %v cannot be evaluated", v)
	}
	return v.Interface(), nil
}

// EvalValue returns the same Dynamic value.
func (d Dynamic) EvalValue() (Value, error) {
	return d, nil
}

// Interface gets the unpacked Go value from this Dynamic value.
func (d Dynamic) Interface() interface{} {
	result, err := d.Eval()
	if err != nil {
		panic(err)
	}
	return result
}

// Type gets the dynamic type of the value.
func (d Dynamic) Type() Type {
	return dynamictype{reflect.Value(d).Type()}
}

// Value gets this Dynamic variable's Dynamic value.
func (d *Dynamic) Value() Value {
	return *d
}

// SetValue sets the value of this dynamic variable.
func (d *Dynamic) SetValue(v Value) error {
	*d = Dynamic(reflect.ValueOf(v.Interface()))
	return nil
}

// GetAttr gets an attribute from this Dynamic value
func (d Dynamic) GetAttr(name String) (Value, error) {
	n := string(name)
	v := reflect.Value(d)
	if v.Kind() == reflect.Struct {
		if f := v.FieldByName(n); f.IsValid() {
			return ValueOf(f.Interface()), nil
		}
	}
	if m := v.MethodByName(n); m.IsValid() {
		results := valuesInterfaces(m.Call(nil))
		if len(results) > 2 {
			return nil, errors.Errorf(
				"method %s of %v returned %d values: %v",
				n, v, len(results), results)
		}
		if len(results) == 1 {
			return ValueOf(results[0]), nil
		}
		err, ok := results[1].(error)
		if !ok {
			return nil, errors.Errorf(
				"second result of method %s of %v is of type %T, not error",
				name, v, results[1])
		}
		if err != nil {
			return nil, err
		}
		return ValueOf(results[0]), nil
	}
	return nil, errors.Errorf("%v has no attribute %q", v, n)
}

func valuesInterfaces(values []reflect.Value) []interface{} {
	results := make([]interface{}, len(values))
	for i, value := range values {
		results[i] = value.Interface()
	}
	return results
}
