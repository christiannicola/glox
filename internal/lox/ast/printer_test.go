package ast_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrinter_Print(t *testing.T) {
	p := ast.NewPrinter()

	expression := ast.NewBinary(
		ast.NewUnary(
			lox.NewToken(lox.Minus, "-", nil, 1),
			ast.NewLiteral("123"),
		),
		lox.NewToken(lox.Star, "*", nil, 1),
		ast.NewGrouping(ast.NewLiteral("45.67")),
	)

	rv, err := p.Print(expression)

	assert.NoError(t, err)
	assert.Equal(t, "(* (- 123) (group 45.67))", rv)
}
