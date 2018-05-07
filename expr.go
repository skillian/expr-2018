package expr

// Expr represents any expression.
type Expr interface {
	// Copy creates a copy of the expression and its children.
	// A set of transformations can be applied to the expression(s)
	Copy(transformations ...Mapper) Expr

	// Eval evaluates the expression into a Go value.
	Eval() (interface{}, error)
}

// Mapper transforms an expression into a new expression.
type Mapper func(e Expr) Expr

// Filter can be used to filter out expressions.
type Filter func(e Expr) bool

// Unary represents a unary expression
type Unary interface {
	Expr
	// Operand gets the Unary expression's operand
	Operand() Expr
}

// Binary is a binary expression.
type Binary interface {
	Expr
	// Left gets the first operand of the Binary expression
	Left() Expr
	// Right gets the second operand of the Binary expression.
	Right() Expr
}

// Multary is a variadic/polyadic/multary expression
type Multary interface {
	Expr
	// Operands() gets all of the operands within the Multary expression.
	Operands() []Expr
}

// ApplyMappers applies a collection of mappers to an expression, taking the
// result of each mapper and feeding it into the next and returns the final
// expression.  ApplyMappers does not apply the mappers to e's children.
func ApplyMappers(e Expr, mappers ...Mapper) Expr {
	for _, mapper := range mappers {
		e = mapper(e)
	}
	return e
}

// Pipeline turns a slice of Mappers into a single mapper pipeline.
func Pipeline(mappers ...Mapper) Mapper {
	return func(e Expr) Expr {
		return ApplyMappers(e, mappers...)
	}
}

// IsTerminal returns true if the expression is a terminal expression in the
// tree.
func IsTerminal(e Expr) bool {
	switch e := e.(type) {
	case Value, Var:
		return true
		_ = e
	}
	return false
}
