package lox_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrinter_Print(t *testing.T) {
	p := lox.NewDebugPrinter()

	minus := ast.NewToken(ast.Minus, "-", nil, 1)
	star := ast.NewToken(ast.Star, "*", nil, 1)

	expression := ast.NewBinary(
		ast.NewUnary(
			&minus,
			ast.NewLiteral("123"),
		),
		&star,
		ast.NewGrouping(ast.NewLiteral("45.67")),
	)

	rv, err := p.Print(expression)

	assert.NoError(t, err)
	assert.Equal(t, "(* (- 123) (group 45.67))", rv)
}
