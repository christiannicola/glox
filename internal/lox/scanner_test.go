package lox_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/christiannicola/glox/internal/lox/ast"
	"github.com/stretchr/testify/assert"
	"testing"
)

type singleTokenTestCase struct {
	description      string
	source           string
	expectedSliceLen int
	tokenType        ast.TokenType
}

func TestScanner_ScanTokens(t *testing.T) {
	cases := []singleTokenTestCase{
		{description: "LeftParenthesis", source: "(", expectedSliceLen: 2, tokenType: ast.LeftParen},
		{description: "RightParenthesis", source: ")", expectedSliceLen: 2, tokenType: ast.RightParen},
		{description: "LeftBrace", source: "{", expectedSliceLen: 2, tokenType: ast.LeftBrace},
		{description: "RightBrace", source: "}", expectedSliceLen: 2, tokenType: ast.RightBrace},
		{description: "Dot", source: ".", expectedSliceLen: 2, tokenType: ast.Dot},
		{description: "Comma", source: ",", expectedSliceLen: 2, tokenType: ast.Comma},
		{description: "Minus", source: "-", expectedSliceLen: 2, tokenType: ast.Minus},
		{description: "Plus", source: "+", expectedSliceLen: 2, tokenType: ast.Plus},
		{description: "Semicolon", source: ";", expectedSliceLen: 2, tokenType: ast.Semicolon},
		{description: "Star", source: "*", expectedSliceLen: 2, tokenType: ast.Star},
		{description: "Bang", source: "!", expectedSliceLen: 2, tokenType: ast.Bang},
		{description: "Equal", source: "=", expectedSliceLen: 2, tokenType: ast.Equal},
		{description: "Less", source: "<", expectedSliceLen: 2, tokenType: ast.Less},
		{description: "Greater", source: ">", expectedSliceLen: 2, tokenType: ast.Greater},
		{description: "Slash", source: "/", expectedSliceLen: 2, tokenType: ast.Slash},
		{description: "BangEqual", source: "!=", expectedSliceLen: 2, tokenType: ast.BangEqual},
		{description: "EqualEqual", source: "==", expectedSliceLen: 2, tokenType: ast.EqualEqual},
		{description: "LessEqual", source: "<=", expectedSliceLen: 2, tokenType: ast.LessEqual},
		{description: "GreaterEqual", source: ">=", expectedSliceLen: 2, tokenType: ast.GreaterEqual},
		{description: "Comment", source: "// this is a comment", expectedSliceLen: 1, tokenType: ast.EOF},
		{description: "EmptyFile", source: "", expectedSliceLen: 1, tokenType: ast.EOF},
		{description: "Whitespace", source: "\t\n\r\f ", expectedSliceLen: 1, tokenType: ast.EOF},
		{description: "String", source: "\"test-string\"", expectedSliceLen: 2, tokenType: ast.String},
		{description: "ThreeStrings", source: "\"test-string-1\" \"test-string-2\"\n   \"test-string-3\"", expectedSliceLen: 4, tokenType: ast.String},
		{description: "Number", source: "123.456", expectedSliceLen: 2, tokenType: ast.Number},
		{description: "ThreeNumbers", source: "123 456.12\n666.77", expectedSliceLen: 4, tokenType: ast.Number},
		{description: "And", source: "and", expectedSliceLen: 2, tokenType: ast.And},
		{description: "Class", source: "class", expectedSliceLen: 2, tokenType: ast.Class},
		{description: "Else", source: "else", expectedSliceLen: 2, tokenType: ast.Else},
		{description: "False", source: "false", expectedSliceLen: 2, tokenType: ast.False},
		{description: "For", source: "for", expectedSliceLen: 2, tokenType: ast.For},
		{description: "Fun", source: "fun", expectedSliceLen: 2, tokenType: ast.Fun},
		{description: "If", source: "if", expectedSliceLen: 2, tokenType: ast.If},
		{description: "Nil", source: "nil", expectedSliceLen: 2, tokenType: ast.Nil},
		{description: "Or", source: "or", expectedSliceLen: 2, tokenType: ast.Or},
		{description: "Print", source: "print", expectedSliceLen: 2, tokenType: ast.Print},
		{description: "Return", source: "return", expectedSliceLen: 2, tokenType: ast.Return},
		{description: "Super", source: "super", expectedSliceLen: 2, tokenType: ast.Super},
		{description: "This", source: "this", expectedSliceLen: 2, tokenType: ast.This},
		{description: "True", source: "true", expectedSliceLen: 2, tokenType: ast.True},
		{description: "Var", source: "var", expectedSliceLen: 2, tokenType: ast.Var},
		{description: "While", source: "while", expectedSliceLen: 2, tokenType: ast.While},
	}

	for i := range cases {
		t.Run(cases[i].description, func(t *testing.T) {
			scanner := lox.NewScanner(cases[i].source)
			tokens, err := scanner.ScanTokens()
			assert.NoError(t, err)
			assert.Len(t, tokens, cases[i].expectedSliceLen)
			assert.Equal(t, tokens[0].Type(), cases[i].tokenType)
			assert.Equal(t, tokens[len(tokens)-1].Type(), ast.EOF)
		})
	}
}
