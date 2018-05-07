package expr

import (
	"github.com/skillian/errors"
)

type unaryBool [1]BoolExpr

func (u unaryBool) copy(transformations ...Mapper) Expr {
	return u.Operand().Copy(transformations...)
}

func (u unaryBool) Operand() Expr {
	return u[0]
}

// Not inverts its operand's boolean result
type Not unaryBool

// Copy the expression.
func (n Not) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Not{unaryBool(n).copy(transformations...).(BoolExpr)}, transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (n Not) EvalBool() (bool, error) {
	op, err := n.Operand().Eval()
	if err != nil {
		return false, errors.ErrorfWithCause(err, "failed to evaluate Not expression")
	}
	return !Truthy(op), nil
}

// Eval the expression.
func (n Not) Eval() (interface{}, error) {
	return n.EvalBool()
}

// Operand gets the unary expression's operand.
func (n Not) Operand() Expr {
	return unaryBool(n).Operand()
}

type binary [2]Expr

// copy copies the binary's operands but the returned binary needs to
// be wrapped by the real Binary implementation's Copy function.
func (b binary) copy(transformations ...Mapper) binary {
	return binary{
		b.Left().Copy(transformations...),
		b.Right().Copy(transformations...),
	}
}

// Left binary operand.
func (b binary) Left() Expr {
	return b[0]
}

// Right binary operand.
func (b binary) Right() Expr {
	return b[1]
}

// Equal defines the equality operator.
type Equal binary

// Eq is an alias for the Equal operator.
type Eq = Equal

// Copy the expression.
func (eq Eq) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Eq(binary(eq).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a boolean result.
func (eq Eq) EvalBool() (bool, error) {
	result, err := Cmp(eq.Left(), eq.Right())
	if err != nil {
		return false, err
	}
	return result == 0, nil
}

// EvalValue evaluates the expression.
func (eq Eq) EvalValue() (Value, error) {
	result, err := eq.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (eq Eq) Eval() (interface{}, error) {
	return eq.EvalBool()
}

// Left side of the binary expression
func (eq Eq) Left() Expr {
	return binary(eq).Left()
}

// Right side of the binary expression
func (eq Eq) Right() Expr {
	return binary(eq).Right()
}

func (eq Eq) String() string {
	return stringifyBinaryInfixHelper(eq, "==")
}

// NotEqual represents the inequality operator.
type NotEqual binary

// Ne is a shorthand for the NotEqual operator.
type Ne = NotEqual

// Copy the expression.
func (ne Ne) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Ne(binary(ne).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (ne Ne) EvalBool() (bool, error) {
	result, err := Cmp(ne.Left(), ne.Right())
	if err != nil {
		return false, err
	}
	return result != 0, nil
}

// EvalValue evaluates the expression.
func (ne Ne) EvalValue() (Value, error) {
	result, err := ne.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (ne Ne) Eval() (interface{}, error) {
	return ne.EvalBool()
}

// Left gets the left side of the binary expression.
func (ne Ne) Left() Expr {
	return binary(ne).Left()
}

// Right gets the left side of the binary expression.
func (ne Ne) Right() Expr {
	return binary(ne).Right()
}

func (ne Ne) String() string {
	return stringifyBinaryInfixHelper(ne, "!=")
}

// GreaterThan represents the > operator.
type GreaterThan binary

// Gt is a shorthand for the GreaterThan operator.
type Gt = GreaterThan

// Copy the expression.
func (gt Gt) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Gt(binary(gt).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (gt Gt) EvalBool() (bool, error) {
	result, err := Cmp(gt.Left(), gt.Right())
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

// EvalValue evaluates the expression.
func (gt Gt) EvalValue() (Value, error) {
	result, err := gt.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (gt Gt) Eval() (interface{}, error) {
	return gt.EvalBool()
}

// Left gets the left side of the binary expression.
func (gt Gt) Left() Expr {
	return binary(gt).Left()
}

// Right gets the left side of the binary expression.
func (gt Gt) Right() Expr {
	return binary(gt).Right()
}

func (gt Gt) String() string {
	return stringifyBinaryInfixHelper(gt, ">")
}

// GreaterThanOrEqual represents the >= operator.
type GreaterThanOrEqual binary

// Ge is a shorthand for the GreaterThan operator.
type Ge = GreaterThanOrEqual

// Copy the expression.
func (ge Ge) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Ge(binary(ge).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (ge Ge) EvalBool() (bool, error) {
	result, err := Cmp(ge.Left(), ge.Right())
	if err != nil {
		return false, err
	}
	return result >= 0, nil
}

// EvalValue evaluates the expression.
func (ge Ge) EvalValue() (Value, error) {
	result, err := ge.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (ge Ge) Eval() (interface{}, error) {
	return ge.EvalBool()
}

// Left gets the left side of the binary expression.
func (ge Ge) Left() Expr {
	return binary(ge).Left()
}

// Right gets the left side of the binary expression.
func (ge Ge) Right() Expr {
	return binary(ge).Right()
}

func (ge Ge) String() string {
	return stringifyBinaryInfixHelper(ge, ">=")
}

// LessThan represents the < operator.
type LessThan binary

// Lt is a shorthand for the LessThan operator.
type Lt = LessThan

// Copy the expression.
func (lt Lt) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Lt(binary(lt).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (lt Lt) EvalBool() (bool, error) {
	result, err := Cmp(lt.Left(), lt.Right())
	if err != nil {
		return false, err
	}
	return result < 0, nil
}

// EvalValue evaluates the expression.
func (lt Lt) EvalValue() (Value, error) {
	result, err := lt.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (lt Lt) Eval() (interface{}, error) {
	return lt.EvalBool()
}

// Left gets the left side of the binary expression.
func (lt Lt) Left() Expr {
	return binary(lt).Left()
}

// Right gets the left side of the binary expression.
func (lt Lt) Right() Expr {
	return binary(lt).Right()
}

func (lt Lt) String() string {
	return stringifyBinaryInfixHelper(lt, "<")
}

// LessThanOrEqual represents the <= operator.
type LessThanOrEqual binary

// Le is a shorthand for the GreaterThan operator.
type Le = LessThanOrEqual

// Copy the expression.
func (le Le) Copy(transformations ...Mapper) Expr {
	return ApplyMappers(Le(binary(le).copy(transformations...)), transformations...)
}

// EvalBool evaluates the expression to a bool result.
func (le Le) EvalBool() (bool, error) {
	result, err := Cmp(le.Left(), le.Right())
	if err != nil {
		return false, err
	}
	return result <= 0, nil
}

// EvalValue evaluates the expression.
func (le Le) EvalValue() (Value, error) {
	result, err := le.EvalBool()
	return Bool(result), err
}

// Eval evaluates the expression.
func (le Le) Eval() (interface{}, error) {
	return le.EvalBool()
}

// Left gets the left side of the binary expression.
func (le Le) Left() Expr {
	return binary(le).Left()
}

// Right gets the left side of the binary expression.
func (le Le) Right() Expr {
	return binary(le).Right()
}

func (le Le) String() string {
	return stringifyBinaryInfixHelper(le, "<=")
}
