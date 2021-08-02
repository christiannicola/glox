package lexer

import "fmt"

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	IDENTIFIER
	STRING
	NUMBER
	AND
	CLASS
	ELSE
	FLASE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF
)

type (
	Token struct {
		tokenType TokenType
		lexeme string
		literal interface{}
		line int64
	}

	TokenType int
)

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int64) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t Token) String() string {
	return fmt.Sprintf("%d %s %v", t.tokenType, t.lexeme, t.literal)
}
