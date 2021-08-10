package lox_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	star       = ast.NewToken(ast.Star, "*", nil, 1)
	slash      = ast.NewToken(ast.Slash, "/", nil, 1)
	plus       = ast.NewToken(ast.Plus, "+", nil, 1)
	minus      = ast.NewToken(ast.Minus, "-", nil, 1)
	less       = ast.NewToken(ast.Less, "<", nil, 1)
	lessEqual  = ast.NewToken(ast.LessEqual, "<=", nil, 1)
	gt         = ast.NewToken(ast.Greater, ">", nil, 1)
	gtEqual    = ast.NewToken(ast.GreaterEqual, ">=", nil, 1)
	bangEqual  = ast.NewToken(ast.BangEqual, "!=", nil, 1)
	equalEqual = ast.NewToken(ast.EqualEqual, "==", nil, 1)

	str    = ast.NewLiteral("this is a string")
	num    = ast.NewLiteral(float64(123))
	truthy = ast.NewLiteral(true)
	falsy  = ast.NewLiteral(false)
	null   = ast.NewLiteral(nil)
)

type (
	evalTestCase struct {
		source        string
		expectedValue interface{}
	}
)

func TestInterpreter_Evaluate(t *testing.T) {
	cases := []evalTestCase{
		{source: "2 + 3", expectedValue: float64(5)},
		{source: "false", expectedValue: false},
		{source: "true", expectedValue: true},
		{source: "nil", expectedValue: nil},
		{source: "-3", expectedValue: float64(-3)},
		{source: "(3 + 2 * (6 / 3)) + 2.54 - 1", expectedValue: float64(8.54)},
		{source: "\"a\" + \"b\"", expectedValue: "ab"},
		{source: "!nil", expectedValue: true},
		{source: "!false", expectedValue: true},
		{source: "!3", expectedValue: false},
		{source: "3 != 3", expectedValue: false},
		{source: "3 <= 3", expectedValue: true},
		{source: "4 < 5", expectedValue: true},
		{source: "2 >= 2", expectedValue: true},
		{source: "4 > 9", expectedValue: false},
		{source: "\"test\" == 3", expectedValue: false},
		{source: "false == false", expectedValue: true},
		{source: "3 == 3", expectedValue: true},
		{source: "nil == nil", expectedValue: true},
		{source: "\"a\" == \"a\"", expectedValue: true},
	}

	for i := range cases {
		scanner := lox.NewScanner(cases[i].source)
		tokens, err := scanner.ScanTokens()
		assert.NoError(t, err)
		assert.NotEmpty(t, tokens)

		parser := lox.NewParser(tokens)
		expr, err := parser.Parse()
		assert.NoError(t, err)
		assert.NotNil(t, expr)

		interpreter := lox.NewInterpreter()
		value, err := interpreter.Evaluate(expr)
		assert.NoError(t, err)
		assert.Equal(t, cases[i].expectedValue, value)
	}
}

func TestInterpreter_VisitBinaryExpr_Failure(t *testing.T) {
	interpreter := lox.NewInterpreter()

	cases := []*ast.Binary{
		ast.NewBinary(str, &star, num),
		ast.NewBinary(nil, &star, str),
		ast.NewBinary(str, &plus, nil),
		ast.NewBinary(null, &slash, str),
		ast.NewBinary(num, &minus, null),
		ast.NewBinary(str, &gt, num),
		ast.NewBinary(str, &gtEqual, num),
		ast.NewBinary(str, &less, num),
		ast.NewBinary(str, &lessEqual, num),
		ast.NewBinary(str, &plus, num),
		ast.NewBinary(num, &plus, str),
	}

	for i := range cases {
		v, err := interpreter.VisitBinaryExpr(cases[i])
		assert.Nil(t, v)
		assert.Error(t, err)
	}
}
