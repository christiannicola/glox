package lexer

import "fmt"

// SyntaxError contains the relevant information to inform the user of a
// syntax error in the source code.
type SyntaxError struct {
	line    int64
	where   string
	message string
}

// Error formats the syntax error info into an easy-to-read format.
func (s SyntaxError) Error() string {
	str := fmt.Sprintf("Error: %s\n", s.message)
	str += fmt.Sprintf("\n\t%d | %s", s.line, s.where)

	return str
}
