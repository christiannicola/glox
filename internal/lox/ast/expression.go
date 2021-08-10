package ast

type (
	// Expression is the base of the expression grammar of the lox programming language.
	// accept takes in a ExprVisitor which will in turn do something meaningful with said expression.
	Expression interface {
		accept(visitor ExprVisitor) (interface{}, error)
	}

	// ExprVisitor is a bundle of abstractions which need to be implemented to fully parse lox syntax.
	ExprVisitor interface {
		BinaryVisitor
		GroupingVisitor
		LiteralVisitor
		UnaryVisitor
	}

	// Binary represents a binary expression in the syntax tree
	Binary struct {
		left     Expression
		operator *Token
		right    Expression
	}

	// BinaryVisitor accepts a binary expression and can return contextual data from it.
	BinaryVisitor interface {
		visitBinaryExpr(expr *Binary) (interface{}, error)
	}

	// Grouping represents a group expression in the syntax tree
	Grouping struct {
		expression Expression
	}

	// GroupingVisitor accepts a grouping expression and can return contextual data from it.
	GroupingVisitor interface {
		visitGroupingExpr(expr *Grouping) (interface{}, error)
	}

	// Literal represents a literal expression in the syntax tree
	Literal struct {
		value interface{}
	}

	// LiteralVisitor accepts a literal expression and can return contextual data from it.
	LiteralVisitor interface {
		visitLiteralExpr(expr *Literal) (interface{}, error)
	}

	// Unary represents a unary expression in the syntax tree
	Unary struct {
		operator *Token
		right    Expression
	}

	// UnaryVisitor accepts a unary expression and can return contextual data from it
	UnaryVisitor interface {
		visitUnaryExpr(expr *Unary) (interface{}, error)
	}
)

// NewBinary returns a pointer to a Binary expression
func NewBinary(left Expression, operator *Token, right Expression) *Binary {
	return &Binary{left, operator, right}
}

func (b *Binary) accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.visitBinaryExpr(b)
}

// NewGrouping returns a pointer to a Grouping expression
func NewGrouping(expression Expression) *Grouping {
	return &Grouping{expression}
}

func (g *Grouping) accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.visitGroupingExpr(g)
}

// NewLiteral returns a pointer to a Literal expression
func NewLiteral(value interface{}) *Literal {
	return &Literal{value}
}

func (l *Literal) accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.visitLiteralExpr(l)
}

// NewUnary returns a pointer to a Unary expression
func NewUnary(operator *Token, right Expression) *Unary {
	return &Unary{operator, right}
}

func (u *Unary) accept(visitor ExprVisitor) (interface{}, error) {
	return visitor.visitUnaryExpr(u)
}
