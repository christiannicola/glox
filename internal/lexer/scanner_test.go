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
		{description: "TestLeftParenthesis", source: "(", expectedSliceLen: 2, tokenType: lexer.LeftParen},
		{description: "TestRightParenthesis", source: ")", expectedSliceLen: 2, tokenType: lexer.RightParen},
		{description: "TestLeftBrace", source: "{", expectedSliceLen: 2, tokenType: lexer.LeftBrace},
		{description: "TestRightBrace", source: "}", expectedSliceLen: 2, tokenType: lexer.RightBrace},
		{description: "TestDot", source: ".", expectedSliceLen: 2, tokenType: lexer.Dot},
		{description: "TestComma", source: ",", expectedSliceLen: 2, tokenType: lexer.Comma},
		{description: "TestMinus", source: "-", expectedSliceLen: 2, tokenType: lexer.Minus},
		{description: "TestPlus", source: "+", expectedSliceLen: 2, tokenType: lexer.Plus},
		{description: "TestSemicolon", source: ";", expectedSliceLen: 2, tokenType: lexer.Semicolon},
		{description: "TestStar", source: "*", expectedSliceLen: 2, tokenType: lexer.Star},
		{description: "TestBang", source: "!", expectedSliceLen: 2, tokenType: lexer.Bang},
		{description: "TestEqual", source: "=", expectedSliceLen: 2, tokenType: lexer.Equal},
		{description: "TestLess", source: "<", expectedSliceLen: 2, tokenType: lexer.Less},
		{description: "TestGreater", source: ">", expectedSliceLen: 2, tokenType: lexer.Greater},
		{description: "TestSlash", source: "/", expectedSliceLen: 2, tokenType: lexer.Slash},
		{description: "TestBangEqual", source: "!=", expectedSliceLen: 2, tokenType: lexer.BangEqual},
		{description: "TestEqualEqual", source: "==", expectedSliceLen: 2, tokenType: lexer.EqualEqual},
		{description: "TestLessEqual", source: "<=", expectedSliceLen: 2, tokenType: lexer.LessEqual},
		{description: "TestGreaterEqual", source: ">=", expectedSliceLen: 2, tokenType: lexer.GreaterEqual},
		{description: "TestComment", source: "// this is a comment", expectedSliceLen: 1, tokenType: lexer.EOF},
		{description: "TestEmptyFile", source: "", expectedSliceLen: 1, tokenType: lexer.EOF},
		{description: "TestWhitespace", source: "\t\n\r\f ", expectedSliceLen: 1, tokenType: lexer.EOF},
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
