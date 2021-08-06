package lexer_test

import (
	"github.com/christiannicola/glox/internal/lexer"
	"github.com/stretchr/testify/assert"
	"testing"
)

type singleTokenTestCase struct {
	description      string
	source           string
	expectedSliceLen int
	tokenType        lexer.TokenType
}

func TestScanner_ScanTokens(t *testing.T) {
	cases := []singleTokenTestCase{
		{description: "LeftParenthesis", source: "(", expectedSliceLen: 2, tokenType: lexer.LeftParen},
		{description: "RightParenthesis", source: ")", expectedSliceLen: 2, tokenType: lexer.RightParen},
		{description: "LeftBrace", source: "{", expectedSliceLen: 2, tokenType: lexer.LeftBrace},
		{description: "RightBrace", source: "}", expectedSliceLen: 2, tokenType: lexer.RightBrace},
		{description: "Dot", source: ".", expectedSliceLen: 2, tokenType: lexer.Dot},
		{description: "Comma", source: ",", expectedSliceLen: 2, tokenType: lexer.Comma},
		{description: "Minus", source: "-", expectedSliceLen: 2, tokenType: lexer.Minus},
		{description: "Plus", source: "+", expectedSliceLen: 2, tokenType: lexer.Plus},
		{description: "Semicolon", source: ";", expectedSliceLen: 2, tokenType: lexer.Semicolon},
		{description: "Star", source: "*", expectedSliceLen: 2, tokenType: lexer.Star},
		{description: "Bang", source: "!", expectedSliceLen: 2, tokenType: lexer.Bang},
		{description: "Equal", source: "=", expectedSliceLen: 2, tokenType: lexer.Equal},
		{description: "Less", source: "<", expectedSliceLen: 2, tokenType: lexer.Less},
		{description: "Greater", source: ">", expectedSliceLen: 2, tokenType: lexer.Greater},
		{description: "Slash", source: "/", expectedSliceLen: 2, tokenType: lexer.Slash},
		{description: "BangEqual", source: "!=", expectedSliceLen: 2, tokenType: lexer.BangEqual},
		{description: "EqualEqual", source: "==", expectedSliceLen: 2, tokenType: lexer.EqualEqual},
		{description: "LessEqual", source: "<=", expectedSliceLen: 2, tokenType: lexer.LessEqual},
		{description: "GreaterEqual", source: ">=", expectedSliceLen: 2, tokenType: lexer.GreaterEqual},
		{description: "Comment", source: "// this is a comment", expectedSliceLen: 1, tokenType: lexer.EOF},
		{description: "EmptyFile", source: "", expectedSliceLen: 1, tokenType: lexer.EOF},
		{description: "Whitespace", source: "\t\n\r\f ", expectedSliceLen: 1, tokenType: lexer.EOF},
		{description: "String", source: "\"test-string\"", expectedSliceLen: 2, tokenType: lexer.String},
		{description: "ThreeStrings", source: "\"test-string-1\" \"test-string-2\"\n   \"test-string-3\"", expectedSliceLen: 4, tokenType: lexer.String},
		{description: "Number", source: "123.456", expectedSliceLen: 2, tokenType: lexer.Number},
		{description: "ThreeNumbers", source: "123 456.12\n666.77", expectedSliceLen: 4, tokenType: lexer.Number},
		{description: "And", source: "and", expectedSliceLen: 2, tokenType: lexer.And},
		{description: "Class", source: "class", expectedSliceLen: 2, tokenType: lexer.Class},
		{description: "Else", source: "else", expectedSliceLen: 2, tokenType: lexer.Else},
		{description: "False", source: "false", expectedSliceLen: 2, tokenType: lexer.False},
		{description: "For", source: "for", expectedSliceLen: 2, tokenType: lexer.For},
		{description: "Fun", source: "fun", expectedSliceLen: 2, tokenType: lexer.Fun},
		{description: "If", source: "if", expectedSliceLen: 2, tokenType: lexer.If},
		{description: "Nil", source: "nil", expectedSliceLen: 2, tokenType: lexer.Nil},
		{description: "Or", source: "or", expectedSliceLen: 2, tokenType: lexer.Or},
		{description: "Print", source: "print", expectedSliceLen: 2, tokenType: lexer.Print},
		{description: "Return", source: "return", expectedSliceLen: 2, tokenType: lexer.Return},
		{description: "Super", source: "super", expectedSliceLen: 2, tokenType: lexer.Super},
		{description: "This", source: "this", expectedSliceLen: 2, tokenType: lexer.This},
		{description: "True", source: "true", expectedSliceLen: 2, tokenType: lexer.True},
		{description: "Var", source: "var", expectedSliceLen: 2, tokenType: lexer.Var},
		{description: "While", source: "while", expectedSliceLen: 2, tokenType: lexer.While},
	}

	for i := range cases {
		t.Run(cases[i].description, func(t *testing.T) {
			scanner := lexer.NewScanner(cases[i].source)
			tokens, err := scanner.ScanTokens()
			assert.NoError(t, err)
			assert.Len(t, tokens, cases[i].expectedSliceLen)
			assert.Equal(t, tokens[0].Type(), cases[i].tokenType)
			assert.Equal(t, tokens[len(tokens)-1].Type(), lexer.EOF)
		})
	}
}
