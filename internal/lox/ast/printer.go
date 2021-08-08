package ast

import (
	"errors"
	"strconv"
)

var errValueNotSerializable = errors.New("unable to serialize value")

// Printer is used to print an expression. Said expression will be returned in its string representation.
type Printer struct{}

// NewPrinter returns a pointer to a new Printer
func NewPrinter() *Printer {
	return &Printer{}
}

// Print returns the Expression in its string representation.
func (p Printer) Print(expr Expression) (string, error) {
	v, err := expr.accept(p)
	if err != nil {
		return "", err
	}

	return p.convertToString(v)
}

func (p Printer) convertToString(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	case float64:
		return strconv.FormatFloat(v.(float64), 'E', -1, 64), nil
	case bool:
		return strconv.FormatBool(v.(bool)), nil
	default:
		return "", errValueNotSerializable
	}
}

func (p Printer) visitBinaryExpr(expr *Binary) (interface{}, error) {
	return p.parenthesize(expr.operator.Lexeme(), expr.left, expr.right)
}

func (p Printer) visitGroupingExpr(expr *Grouping) (interface{}, error) {
	return p.parenthesize("group", expr.expression)
}

func (p Printer) visitLiteralExpr(expr *Literal) (interface{}, error) {
	if expr.value == nil {
		return "nil", nil
	}

	return expr.value, nil
}

func (p Printer) visitUnaryExpr(expr *Unary) (interface{}, error) {
	return p.parenthesize(expr.operator.Lexeme(), expr.right)
}

func (p Printer) parenthesize(name string, expressions ...Expression) (string, error) {
	rv := "(" + name

	for i := range expressions {
		rv += " "

		v, err := expressions[i].accept(p)

		if err != nil {
			return "", err
		}

		s, err := p.convertToString(v)

		if err != nil {
			return "", err
		}

		rv += s
	}

	rv += ")"

	return rv, nil
}
