package lox

import (
	"errors"
	"fmt"
	"github.com/christiannicola/glox/internal/lox/ast"
	"strconv"
)

var (
	errOperatorNil       = errors.New("operator is nil")
	errMissingExpression = errors.New("missing expression")
	errInvalidFloat64    = errors.New("value is not of type float64")
	errInvalidValue      = errors.New("value type is invalid")
)

// Interpreter is the last part of our language pipeline. Its job is to interpret and evaluate expressions
// that Parser constructed from Tokens that were scanned from source code by Scanner.
type Interpreter struct{}

// NewInterpreter returns a new instance of Interpreter.
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// Interpret takes in a slice of statements and interprets each statement one after another.
// Returns an error if a statement cannot be interpreted correctly, otherwise nil.
func (i Interpreter) Interpret(statements []ast.Stmt) error {
	for j := range statements {
		if err := i.execute(statements[j]); err != nil {
			return nil
		}
	}

	return nil
}

func (i Interpreter) execute(statement ast.Stmt) error {
	return statement.Accept(i)
}

func (i Interpreter) evaluate(expr ast.Expr) (interface{}, error) {
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

func (i Interpreter) stringifyValue(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case float64:
		// return fmt.Sprintf("%f", v)
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v.(bool))
	case nil:
		return "nil"
	default:
		return ""
	}
}

// VisitBinaryExpr takes in expr and returns the interpreted value for expr.
// The return value is either float64, string, nil or bool.
// Also returns an error if an invalid value is encountered.
func (i Interpreter) VisitBinaryExpr(expr *ast.Binary) (interface{}, error) {
	if expr == nil || expr.Left() == nil || expr.Right() == nil {
		return nil, errMissingExpression
	}

	left, err := i.evaluate(expr.Left())
	if err != nil {
		return nil, err
	}

	right, err := i.evaluate(expr.Right())
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
// This is done by evaluating the Expr that is stored inside expr.
// See VisitUnaryExpr, VisitLiteralExpr & VisitBinaryExpr for possible return values.
func (i Interpreter) VisitGroupingExpr(expr *ast.Grouping) (interface{}, error) {
	return i.evaluate(expr.Expression())
}

// VisitLiteralExpr returns the literal value of expr without evaluation.
func (i Interpreter) VisitLiteralExpr(expr *ast.Literal) (interface{}, error) {
	return expr.Value(), nil
}

// VisitUnaryExpr takes in expr and returns the interpreted value for expr.
// The return value is either float64 or nil.
// Also returns an error if an invalid value is encountered.
func (i Interpreter) VisitUnaryExpr(expr *ast.Unary) (interface{}, error) {
	right, err := i.evaluate(expr.Right())
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

// VisitPrintStmt takes in a pointer to an ast.PrintStmt, evaluates the underlying expression
// and prints the result to the console.
func (i Interpreter) VisitPrintStmt(stmt *ast.PrintStmt) error {
	value, err := i.evaluate(stmt.Expression())
	if err != nil {
		return err
	}

	fmt.Println(i.stringifyValue(value))

	return err
}

// VisitExprStmt takes in a pointer to an ast.ExprStmt, evaluates the underlying expression and
// discards the resulting value.
func (i Interpreter) VisitExprStmt(stmt *ast.ExprStmt) error {
	_, err := i.evaluate(stmt.Expression())

	return err
}
