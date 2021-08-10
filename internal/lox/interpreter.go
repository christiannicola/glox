package lox

import (
	"errors"
	"github.com/christiannicola/glox/internal/lox/ast"
)

var (
	errOperatorNil       = errors.New("operator is nil")
	errMissingExpression = errors.New("missing expression")
	errInvalidFloat64    = errors.New("value is not of type float64")
	errInvalidString     = errors.New("value is not of type string")
	errInvalidValue      = errors.New("value type is invalid")
)

// Interpreter is the last part of our language pipeline. Its job is to interpret and evaluate expressions
// that Parser constructed from Tokens that were scanned from source code by Scanner.
type Interpreter struct{}

// NewInterpreter returns a new instance of Interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// Evaluate evaluates expr and returns the value for expr.
func (i Interpreter) Evaluate(expr ast.Expression) (interface{}, error) {
	return expr.Accept(i)
}

func (i Interpreter) isTruthy(v interface{}) bool {
	if v == nil {
		return false
	}

	if b, ok := v.(bool); ok {
		return b
	}

	return true
}

func (i Interpreter) isString(v interface{}) bool {
	_, ok := v.(string)

	return ok
}

func (i Interpreter) isFloat64(v interface{}) bool {
	_, ok := v.(float64)

	return ok
}

func (i Interpreter) isEqual(a interface{}, b interface{}) bool {
	return a == b
}

// VisitBinaryExpr takes in expr and returns the interpreted value for expr.
// The return value is either float64, string, nil or bool.
// Also returns an error if an invalid value is encountered.
func (i Interpreter) VisitBinaryExpr(expr *ast.Binary) (interface{}, error) {
	if expr == nil || expr.Left() == nil || expr.Right() == nil {
		return nil, errMissingExpression
	}

	left, err := i.Evaluate(expr.Left())
	if err != nil {
		return nil, err
	}

	right, err := i.Evaluate(expr.Right())
	if err != nil {
		return nil, err
	}

	operator := expr.Operator()
	if operator == nil {
		return nil, errOperatorNil
	}

	switch operator.Type() {
	case ast.Minus:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) - right.(float64), nil
	case ast.Slash:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) / right.(float64), nil
	case ast.Star:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) * right.(float64), nil
	case ast.Plus:
		if i.isFloat64(left) && i.isFloat64(right) {
			return left.(float64) + right.(float64), nil
		}

		if i.isString(left) && i.isString(right) {
			return left.(string) + right.(string), nil
		}

		return nil, errInvalidValue
	case ast.Greater:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) > right.(float64), nil
	case ast.GreaterEqual:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) >= right.(float64), nil
	case ast.Less:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) < right.(float64), nil
	case ast.LessEqual:
		if !i.isFloat64(left) || !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return left.(float64) <= right.(float64), nil
	case ast.BangEqual:
		return !i.isEqual(left, right), nil
	case ast.EqualEqual:
		return i.isEqual(left, right), nil
	default:
		return nil, nil
	}
}

// VisitGroupingExpr takes in expr and returns the interpreted value for expr.
// This is done by evaluating the Expression that is stored inside expr.
// See VisitUnaryExpr, VisitLiteralExpr & VisitBinaryExpr for possible return values.
func (i Interpreter) VisitGroupingExpr(expr *ast.Grouping) (interface{}, error) {
	return i.Evaluate(expr.Expression())
}

// VisitLiteralExpr returns the literal value of expr without evaluation.
func (i Interpreter) VisitLiteralExpr(expr *ast.Literal) (interface{}, error) {
	return expr.Value(), nil
}

// VisitUnaryExpr takes in expr and returns the interpreted value for expr.
// The return value is either float64 or nil.
// Also returns an error if an invalid value is encountered.
func (i Interpreter) VisitUnaryExpr(expr *ast.Unary) (interface{}, error) {
	right, err := i.Evaluate(expr.Right())
	if err != nil {
		return nil, err
	}

	operator := expr.Operator()
	if operator == nil {
		return nil, errOperatorNil
	}

	switch operator.Type() {
	case ast.Minus:
		if !i.isFloat64(right) {
			return nil, errInvalidFloat64
		}

		return -right.(float64), nil
	case ast.Bang:
		return !i.isTruthy(right), nil
	default:
		return nil, nil
	}
}
