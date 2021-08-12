package ast

type (
	// Expr is the base of the expression grammar of the lox programming language.
	// Accept takes in a ExprVisitor which will in turn do something meaningful with said expression.
	Expr interface {
		Accept(visitor ExprVisitor) (interface{}, error)
	}

	// ExprVisitor is a bundle of abstractions which need to be implemented to fully interpret and evaluate
	// the lox programming language and its grammar.
	ExprVisitor interface {
		BinaryVisitor
		GroupingVisitor
		LiteralVisitor
		UnaryVisitor
	}

	// Binary represents a binary expression in the syntax tree
	Binary struct {
		left     Expr
		operator *Token
		right    Expr
	}

	// BinaryVisitor accepts a binary expression and can return contextual data from it.
	BinaryVisitor interface {
		VisitBinaryExpr(expr *Binary) (interface{}, error)
	}

	// Grouping represents a group expression in the syntax tree
	Grouping struct {
		expression Expr
	}

	// GroupingVisitor accepts a grouping expression and can return contextual data from it.
	GroupingVisitor interface {
		VisitGroupingExpr(expr *Grouping) (interface{}, error)
	}

	// Literal represents a literal expression in the syntax tree
	Literal struct {
		value interface{}
	}

	// LiteralVisitor accepts a literal expression and can return contextual data from it.
	LiteralVisitor interface {
		VisitLiteralExpr(expr *Literal) (interface{}, error)
	}

	// Unary represents a unary expression in the syntax tree
	Unary struct {
		operator *Token
		right    Expr
	}

	// UnaryVisitor accepts a unary expression and can return contextual data from it
	UnaryVisitor interface {
		VisitUnaryExpr(expr *Unary) (interface{}, error)
	}
)

// NewBinary returns a pointer to a Binary expression
func NewBinary(left Expr, operator *Token, right Expr) *Binary {
	return &Binary{left, operator, right}
}

// Accept takes in an ExprVisitor and calls its implementation of VisitBinaryExpr.
// Implements Expr.
func (b *Binary) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitBinaryExpr(b)
}

// Left returns the left-hand side expression of b
func (b Binary) Left() Expr {
	return b.left
}

// Right returns the right-hand side expression of b
func (b Binary) Right() Expr {
	return b.right
}

// Operator returns the operator Token of b
func (b Binary) Operator() *Token {
	return b.operator
}

// NewGrouping returns a pointer to a Grouping expression
func NewGrouping(expression Expr) *Grouping {
	return &Grouping{expression}
}

// Accept takes in an ExprVisitor and calls its implementation of VisitGroupingExpr.
// Implements Expr.
func (g *Grouping) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitGroupingExpr(g)
}

// Expression returns the underlying Expr of g.
func (g Grouping) Expression() Expr {
	return g.expression
}

// NewLiteral returns a pointer to a Literal expression
func NewLiteral(value interface{}) *Literal {
	return &Literal{value}
}

// Accept takes in an ExprVisitor and calls its implementation of VisitLiteralExpr.
// Implements Expr.
func (l *Literal) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitLiteralExpr(l)
}

// Value returns the literal value of l
func (l *Literal) Value() interface{} {
	return l.value
}

// NewUnary returns a pointer to a Unary expression
func NewUnary(operator *Token, right Expr) *Unary {
	return &Unary{operator, right}
}

// Accept takes in an ExprVisitor and calls its implementation of VisitUnaryExpr.
// Implements Expr.
func (u *Unary) Accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.VisitUnaryExpr(u)
}

// Right returns the right-hand side expression of u
func (u Unary) Right() Expr {
	return u.right
}

// Operator returns the token for the unary expression u
func (u Unary) Operator() *Token {
	return u.operator
}
