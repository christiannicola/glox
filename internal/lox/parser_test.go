package lox_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	// NOTE (c.nicola): (-123 - 5) * 10
	tokens := []ast.Token{
		ast.NewToken(ast.LeftParen, "(", nil, 1),
		ast.NewToken(ast.Minus, "-", nil, 1),
		ast.NewToken(ast.Number, "123", 123, 1),
		ast.NewToken(ast.Minus, "-", nil, 1),
		ast.NewToken(ast.Number, "5", 5, 1),
		ast.NewToken(ast.RightParen, ")", nil, 1),
		ast.NewToken(ast.Star, "*", nil, 1),
		ast.NewToken(ast.Number, "10", 5, 1),
		ast.NewToken(ast.EOF, "", nil, 1),
	}

	parser := lox.NewParser(tokens)

	_, err := parser.Parse()
	assert.NoError(t, err)
}
