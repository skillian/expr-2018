package expr

import (
	"fmt"

	"github.com/skillian/errors"
)

// EvalValue evaluates the full expression into a Value
func EvalValue(e ValueExpr) (Value, error) {
	result := Simplify(e, EvaluateVariable)
	value, ok := result.(Value)
	if ok {
		return value, nil
	}
	return nil, errors.Errorf(
		"Evaluating expression %v did not produce a value (result: %v)",
		e, result)
}

// EvalValueInto evaluates the given expressions and puts their results into the
// slice of values
func EvalValueInto(exprs []ValueExpr, values []Value) (err error) {
	if len(exprs) > len(values) {
		return errors.Errorf("Expr slice is bigger than target value slice")
	}
	for i, expr := range exprs {
		values[i], err = EvalValue(expr)
		if err != nil {
			return err
		}
	}
	return nil
}

// Simplify performs some compile-time-like simlifications on a copy of its
// expression and returns it.  Arithmetic and logical expressions on constants
// will be evaluated but not operations on Vars (unless explicitly included
// in the additional simplifications).
func Simplify(e Expr, simplifications ...Mapper) Expr {
	builtins := []Mapper{}
	simplifications = append(builtins, simplifications...)
	return e.Copy(simplifications...)
}

// SimplifyAll simplifies the slice of expressions into a new slice of
// expressions.
func SimplifyAll(source []Expr, simplifications ...Mapper) (target []Expr) {
	target = make([]Expr, len(source))
	SimplifySlice(source, target, simplifications...)
	return target
}

// SimplifySlice simplifies the expressions in source into the target slice
// according to the same rules as Simplify.
func SimplifySlice(source []Expr, target []Expr, simplifications ...Mapper) {
	for i, expr := range source {
		target[i] = Simplify(expr, simplifications...)
	}
}

// SimplifyAllInPlace simplifies every expression in exprs in place and returns
// the same slice.
func SimplifyAllInPlace(exprs []Expr, simplifications ...Mapper) {
	SimplifySlice(exprs, exprs, simplifications...)
}

// EvaluateVariable is a Mapper that, when it encounters a Var, will get its
// Value and return that.  This is useful when you intend to evaluate the
// expression locally.
func EvaluateVariable(e Expr) Expr {
	logger.Error1("EvaluatingVariable: %v", e)
	vari, ok := e.(Var)
	if ok {
		value, ok := vari.Value().(Expr)
		if ok {
			fmt.Printf("EvaluateVariable %v -> %v", e, value)
			return value
		}
	}
	return e
}

// SimplifyLeftRight is a helper function to simplify two expressions at once.
func SimplifyLeftRight(left, right Expr, simplifications ...Mapper) [2]Expr {
	return [2]Expr{
		Simplify(left, simplifications...),
		Simplify(right, simplifications...),
	}
}
