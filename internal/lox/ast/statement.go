package ast

type (
	// Stmt represents an expression statement in the Lox programming language.
	Stmt interface {
		Accept(visitor StmtVisitor) error
	}

	// StmtVisitor is a bundle of abstractions which need to be implemented to fully interpret and evaluate
	// expression statements
	StmtVisitor interface {
		PrintStmtVisitor
		ExprStmtVisitor
	}

	// ExprStmt represents a statement that you can use to place
	// an Expr where a Stmt is expected. They exist to evaluate expressions
	// that have side effects.
	ExprStmt struct {
		expr Expr
	}

	// ExprStmtVisitor accepts an expression statement and evaluates it.
	ExprStmtVisitor interface {
		VisitExprStmt(stmt *ExprStmt) error
	}

	// PrintStmt represents a statement which is evaluated; the resulting value
	// will then be printed to the console.
	PrintStmt struct {
		expr Expr
	}

	// PrintStmtVisitor accepts an expression statement and evaluates it.
	PrintStmtVisitor interface {
		VisitPrintStmt(stmt *PrintStmt) error
	}
)

// NewPrintStmt returns a new PrintStmt
func NewPrintStmt(value Expr) *PrintStmt {
	return &PrintStmt{value}
}

// Accept takes in an StmtVisitor and calls its implementation of VisitPrintStmt.
// Implements Stmt.
func (p *PrintStmt) Accept(visitor StmtVisitor) error {
	return visitor.VisitPrintStmt(p)
}

// Expression returns the Expr of p
func (p PrintStmt) Expression() Expr {
	return p.expr
}

// NewExprStmt returns a new ExprStmt
func NewExprStmt(expr Expr) *ExprStmt {
	return &ExprStmt{expr}
}

// Accept takes in an ExprStmt and calls its implementation of VisitExprStmt.
// Implements Stmt.
func (e *ExprStmt) Accept(visitor StmtVisitor) error {
	return visitor.VisitExprStmt(e)
}

// Expression returns the Expr of e
func (e ExprStmt) Expression() Expr {
	return e.expr
}
