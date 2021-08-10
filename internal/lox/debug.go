package lox

import (
	"errors"
	"github.com/christiannicola/glox/internal/lox/ast"
	"strconv"
)

var errValueNotSerializable = errors.New("unable to serialize value")

// DebugPrinter is used to convert an Expression to a human-readable string.
// Useful for debugging Scanner / Parser passes.
type DebugPrinter struct{}

// NewDebugPrinter returns a pointer to a new DebugPrinter
func NewDebugPrinter() *DebugPrinter {
	return &DebugPrinter{}
}

func (p DebugPrinter) stringifyValue(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64), nil
	case bool:
		return strconv.FormatBool(v.(bool)), nil
	case nil:
		return "nil", nil
	default:
		return "", errValueNotSerializable
	}
}

// Print returns the Expression in its string representation.
func (p DebugPrinter) Print(expr ast.Expression) (string, error) {
	v, err := expr.Accept(p)
	if err != nil {
		return "", err
	}

	return p.stringifyValue(v)
}

// VisitBinaryExpr converts a binary expression into a human-readable string.
func (p DebugPrinter) VisitBinaryExpr(expr *ast.Binary) (interface{}, error) {
	return p.parenthesize(expr.Operator().Lexeme(), expr.Left(), expr.Right())
}

// VisitGroupingExpr converts a group expression into a human-readable string.
func (p DebugPrinter) VisitGroupingExpr(expr *ast.Grouping) (interface{}, error) {
	return p.parenthesize("group", expr.Expression())
}

// VisitLiteralExpr converts a literal expression into a human-readable string.
func (p DebugPrinter) VisitLiteralExpr(expr *ast.Literal) (interface{}, error) {
	if expr.Value() == nil {
		return "nil", nil
	}

	return expr.Value(), nil
}

// VisitUnaryExpr converts a literal expression into a human-readable string.
func (p DebugPrinter) VisitUnaryExpr(expr *ast.Unary) (interface{}, error) {
	return p.parenthesize(expr.Operator().Lexeme(), expr.Right())
}

func (p DebugPrinter) parenthesize(name string, expressions ...ast.Expression) (string, error) {
	rv := "(" + name

	for i := range expressions {
		rv += " "

		v, err := expressions[i].Accept(p)
		if err != nil {
			return "", err
		}

		s, err := p.stringifyValue(v)
		if err != nil {
			return "", err
		}

		rv += s
	}

	rv += ")"

	return rv, nil
}
