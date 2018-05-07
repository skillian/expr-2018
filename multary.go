package expr

type multaryBool []BoolExpr

func (m multaryBool) copy(transformations ...Mapper) multaryBool {
	m2 := make(multaryBool, len(m))
	for i, e := range m {
		m2[i] = e.Copy(transformations...).(BoolExpr)
	}
	return m2
}

func (m multaryBool) Operands() []Expr {
	operands := make([]Expr, len(m))
	for i, operand := range m {
		operands[i] = operand
	}
	return operands
}

// All evaluates to true if all of its inner BoolExprs evaluate to true.
type All multaryBool

// Copy the expression
func (a All) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(All(multaryBool(a).copy()), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (a All) EvalBool() (bool, error) {
	for _, e := range a {
		ok, err := e.EvalBool()
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

// EvalValue evaluates the expression to an expr.Value
func (a All) EvalValue() (Value, error) {
	ok, err := a.EvalBool()
	return Bool(ok), err
}

// Eval evaluates the expression.
func (a All) Eval() (interface{}, error) {
	return a.EvalBool()
}

// Operands gets all of the operands within the All expression.
func (a All) Operands() []Expr {
	return multaryBool(a).Operands()
}

// Any returns true if any of its inner BoolExprs evaluate to true.
type Any []BoolExpr

// Copy the expression.
func (a Any) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Any(multaryBool(a).copy()), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (a Any) EvalBool() (bool, error) {
	for _, e := range a {
		ok, err := e.EvalBool()
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}

// EvalValue just calls EvalBool.
func (a Any) EvalValue() (Value, error) {
	ok, err := a.EvalBool()
	return Bool(ok), err
}

// Eval just calls EvalBool.
func (a Any) Eval() (interface{}, error) {
	return a.EvalBool()
}

// Operands gets all of the operands within the Any expression.
func (a Any) Operands() []Expr {
	return multaryBool(a).Operands()
}
