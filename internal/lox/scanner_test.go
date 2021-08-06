package lox_test

import (
	"github.com/christiannicola/glox/internal/lox"
	"github.com/stretchr/testify/assert"
	"testing"
)

type singleTokenTestCase struct {
	description      string
	source           string
	expectedSliceLen int
	tokenType        lox.TokenType
}

func TestScanner_ScanTokens(t *testing.T) {
	cases := []singleTokenTestCase{
		{description: "LeftParenthesis", source: "(", expectedSliceLen: 2, tokenType: lox.LeftParen},
		{description: "RightParenthesis", source: ")", expectedSliceLen: 2, tokenType: lox.RightParen},
		{description: "LeftBrace", source: "{", expectedSliceLen: 2, tokenType: lox.LeftBrace},
		{description: "RightBrace", source: "}", expectedSliceLen: 2, tokenType: lox.RightBrace},
		{description: "Dot", source: ".", expectedSliceLen: 2, tokenType: lox.Dot},
		{description: "Comma", source: ",", expectedSliceLen: 2, tokenType: lox.Comma},
		{description: "Minus", source: "-", expectedSliceLen: 2, tokenType: lox.Minus},
		{description: "Plus", source: "+", expectedSliceLen: 2, tokenType: lox.Plus},
		{description: "Semicolon", source: ";", expectedSliceLen: 2, tokenType: lox.Semicolon},
		{description: "Star", source: "*", expectedSliceLen: 2, tokenType: lox.Star},
		{description: "Bang", source: "!", expectedSliceLen: 2, tokenType: lox.Bang},
		{description: "Equal", source: "=", expectedSliceLen: 2, tokenType: lox.Equal},
		{description: "Less", source: "<", expectedSliceLen: 2, tokenType: lox.Less},
		{description: "Greater", source: ">", expectedSliceLen: 2, tokenType: lox.Greater},
		{description: "Slash", source: "/", expectedSliceLen: 2, tokenType: lox.Slash},
		{description: "BangEqual", source: "!=", expectedSliceLen: 2, tokenType: lox.BangEqual},
		{description: "EqualEqual", source: "==", expectedSliceLen: 2, tokenType: lox.EqualEqual},
		{description: "LessEqual", source: "<=", expectedSliceLen: 2, tokenType: lox.LessEqual},
		{description: "GreaterEqual", source: ">=", expectedSliceLen: 2, tokenType: lox.GreaterEqual},
		{description: "Comment", source: "// this is a comment", expectedSliceLen: 1, tokenType: lox.EOF},
		{description: "EmptyFile", source: "", expectedSliceLen: 1, tokenType: lox.EOF},
		{description: "Whitespace", source: "\t\n\r\f ", expectedSliceLen: 1, tokenType: lox.EOF},
		{description: "String", source: "\"test-string\"", expectedSliceLen: 2, tokenType: lox.String},
		{description: "ThreeStrings", source: "\"test-string-1\" \"test-string-2\"\n   \"test-string-3\"", expectedSliceLen: 4, tokenType: lox.String},
		{description: "Number", source: "123.456", expectedSliceLen: 2, tokenType: lox.Number},
		{description: "ThreeNumbers", source: "123 456.12\n666.77", expectedSliceLen: 4, tokenType: lox.Number},
		{description: "And", source: "and", expectedSliceLen: 2, tokenType: lox.And},
		{description: "Class", source: "class", expectedSliceLen: 2, tokenType: lox.Class},
		{description: "Else", source: "else", expectedSliceLen: 2, tokenType: lox.Else},
		{description: "False", source: "false", expectedSliceLen: 2, tokenType: lox.False},
		{description: "For", source: "for", expectedSliceLen: 2, tokenType: lox.For},
		{description: "Fun", source: "fun", expectedSliceLen: 2, tokenType: lox.Fun},
		{description: "If", source: "if", expectedSliceLen: 2, tokenType: lox.If},
		{description: "Nil", source: "nil", expectedSliceLen: 2, tokenType: lox.Nil},
		{description: "Or", source: "or", expectedSliceLen: 2, tokenType: lox.Or},
		{description: "Print", source: "print", expectedSliceLen: 2, tokenType: lox.Print},
		{description: "Return", source: "return", expectedSliceLen: 2, tokenType: lox.Return},
		{description: "Super", source: "super", expectedSliceLen: 2, tokenType: lox.Super},
		{description: "This", source: "this", expectedSliceLen: 2, tokenType: lox.This},
		{description: "True", source: "true", expectedSliceLen: 2, tokenType: lox.True},
		{description: "Var", source: "var", expectedSliceLen: 2, tokenType: lox.Var},
		{description: "While", source: "while", expectedSliceLen: 2, tokenType: lox.While},
	}

	for i := range cases {
		t.Run(cases[i].description, func(t *testing.T) {
			scanner := lox.NewScanner(cases[i].source)
			tokens, err := scanner.ScanTokens()
			assert.NoError(t, err)
			assert.Len(t, tokens, cases[i].expectedSliceLen)
			assert.Equal(t, tokens[0].Type(), cases[i].tokenType)
			assert.Equal(t, tokens[len(tokens)-1].Type(), lox.EOF)
		})
	}
}
