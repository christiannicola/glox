package lox

import "fmt"

// These constants define the types of tokens that are of interest to the lexer and parser
const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	Identifier
	String
	Number
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
	EOF
)

type (
	// Token represents a token of scanned source code.
	Token struct {
		tokenType TokenType
		lexeme    string
		literal   interface{}
		line      int64
	}

	// TokenType describes the type of Token
	TokenType int
)

// NewToken returns a new Token
func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int64) *Token {
	return &Token{tokenType, lexeme, literal, line}
}

// String returns token information for debugging purposes
func (t Token) String() string {
	return fmt.Sprintf("%d %s %v", t.tokenType, t.Lexeme(), t.literal)
}

// Type returns the TokenType of t
func (t Token) Type() TokenType {
	return t.tokenType
}

// Lexeme returns the lexeme of t
func (t Token) Lexeme() string {
	return t.lexeme
}
