package lox

import (
	"fmt"
	"github.com/christiannicola/glox/internal/lox/ast"
	"os"
)

type (
	// SyntaxError contains the relevant information to inform the user of a
	// syntax error in the source code.
	SyntaxError struct {
		line    int64
		where   string
		message string
	}

	// ParseError contains the relevant information to inform the user of an error
	// that occurred during the parsing of the tokens.
	ParseError struct {
		token   *ast.Token
		line    int64
		where   string
		message string
	}
)

// Error formats the syntax error info into an easy-to-read format.
func (s SyntaxError) Error() string {
	str := fmt.Sprintf("Error: %s\n", s.message)
	str += fmt.Sprintf("\n\t%d | %s", s.line, s.where)

	return str
}

// Error formats the parse error info into an easy-to-read format.
func (p ParseError) Error() string {
	str := fmt.Sprintf("Error: %s\n", p.message)

	if p.token.Type() == ast.EOF {
		str += fmt.Sprintf("\n\t%d | %s", p.line, "at end")
	} else {
		str += fmt.Sprintf("\n\t%d | %s", p.line, p.where)
	}

	return str
}

// PrintError prints the error to the console
func PrintError(err error) {
	_, _ = fmt.Fprint(os.Stderr, err.Error())
}
