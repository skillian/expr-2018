package expr

import (
	"github.com/skillian/errors"
)

type binaryValue [2]ValueExpr

func (bv binaryValue) copy(transformations ...Mapper) binaryValue {
	return binaryValue{
		bv.Left().Copy(transformations...).(ValueExpr),
		bv.Right().Copy(transformations...).(ValueExpr),
	}
}

// Left side of binary expression
func (bv binaryValue) Left() Expr {
	return bv[0]
}

// Right side of binary expression
func (bv binaryValue) Right() Expr {
	return bv[1]
}

// Addition binary expression
type Addition binaryValue

// Add is short for Addition.
type Add = Addition

// EvalAdder is the interface that must be implemented in order to be "addable."
type EvalAdder interface {
	EvalAdd(v Value) (Value, error)
}

// Copy the expression.
func (a Add) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Add(binaryValue(a).copy(transformations...)), transformations...)
}

// Eval the expression.
func (a Add) Eval() (interface{}, error) {
	value, err := a.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the expression to a value.
func (a Add) EvalValue() (Value, error) {
	var operands [2]Value
	err := EvalValueInto(a[:], operands[:])
	if err != nil {
		return nil, err
	}
	adder, ok := operands[0].(EvalAdder)
	if !ok {
		return nil, errors.Errorf("cannot add %v to %v", operands[0], operands[1])
	}
	return adder.EvalAdd(operands[1])
}

// Left gets the left side of the binary expression.
func (a Add) Left() Expr {
	return a[0]
}

// Right gets the right side of the binary expression.
func (a Add) Right() Expr {
	return a[1]
}

// String represents the expression as a string.
func (a Add) String() string {
	return stringifyBinaryInfixHelper(a, "+")
}

// Subtraction binary expression
type Subtraction binaryValue

// Sub is short for Subtraction.
type Sub = Subtraction

// EvalSubtracter is the interface that must be implemented in order to be "subtractable."
type EvalSubtracter interface {
	EvalSubtract(v Value) (Value, error)
}

// Copy the expression.
func (s Sub) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Sub(binaryValue(s).copy(transformations...)), transformations...)
}

// Eval the expression.
func (s Sub) Eval() (interface{}, error) {
	value, err := s.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the expression to a value.
func (s Sub) EvalValue() (Value, error) {
	var operands [2]Value
	err := EvalValueInto(s[:], operands[:])
	if err != nil {
		return nil, err
	}
	subber, ok := operands[0].(EvalSubtracter)
	if !ok {
		return nil, errors.Errorf("cannot subtract %v from %v", operands[0], operands[1])
	}
	return subber.EvalSubtract(operands[1])
}

// Left side of the binary expression.
func (s Sub) Left() Expr {
	return s[0]
}

// Right side of the binary expression.
func (s Sub) Right() Expr {
	return s[1]
}

// String represents the expression as a string.
func (s Sub) String() string {
	return stringifyBinaryInfixHelper(s, "+")
}

// Multiplication binary expression
type Multiplication binaryValue

// Mul is short for Multiplication.
type Mul = Multiplication

// EvalMultiplier is the interface that must be implemented in order to be "multipliable."
type EvalMultiplier interface {
	EvalMultiply(v Value) (Value, error)
}

// Copy the expression.
func (m Mul) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Mul(binaryValue(m).copy(transformations...)), transformations...)
}

// Eval the expression.
func (m Mul) Eval() (interface{}, error) {
	value, err := m.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the expression to a value.
func (m Mul) EvalValue() (Value, error) {
	var operands [2]Value
	err := EvalValueInto(m[:], operands[:])
	if err != nil {
		return nil, err
	}
	multiplier, ok := operands[0].(EvalMultiplier)
	if !ok {
		return nil, errors.Errorf("cannot multiply %v to %v", operands[0], operands[1])
	}
	return multiplier.EvalMultiply(operands[1])
}

// Left side of the binary expression.
func (m Mul) Left() Expr {
	return m[0]
}

// Right side of the binary expression.
func (m Mul) Right() Expr {
	return m[1]
}

// String represents the expression as a string.
func (m Mul) String() string {
	return stringifyBinaryInfixHelper(m, "*")
}

// Division binary expression
type Division binaryValue

// Div is short for Subtraction.
type Div = Division

// EvalDivider is the interface that must be implemented in order to be divisable.
type EvalDivider interface {
	EvalDivide(v Value) (Value, error)
}

// Copy the expression.
func (d Div) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Div(binaryValue(d).copy(transformations...)), transformations...)
}

// Eval the expression.
func (d Div) Eval() (interface{}, error) {
	value, err := d.EvalValue()
	if err != nil {
		return nil, err
	}
	return value.Interface(), nil
}

// EvalValue evaluates the expression to a value.
func (d Div) EvalValue() (Value, error) {
	left, err := d[0].EvalValue()
	if err != nil {
		return nil, err
	}
	right, err := d[1].EvalValue()
	if err != nil {
		return nil, err
	}
	diver, ok := left.(EvalDivider)
	if !ok {
		return nil, errors.Errorf("cannot divide %v by %v", left, right)
	}
	return diver.EvalDivide(right)
}

// Left side of the binary expression.
func (d Div) Left() Expr {
	return d[0]
}

// Right side of the binary expression.
func (d Div) Right() Expr {
	return d[1]
}

// String represents the expression as a string.
func (d Div) String() string {
	return stringifyBinaryInfixHelper(d, "/")
}
