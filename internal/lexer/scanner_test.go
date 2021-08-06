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
