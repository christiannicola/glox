package lexer

import (
	"fmt"
	"io"
	"strings"
)

// Scanner is responsible for scanning source code line by line and generating the tokens from
// said source code.
type Scanner struct {
	source  string
	reader  *strings.Reader
	tokens  []Token
	start   int64
	current int64
	line    int64
}

// NewScanner returns a new Scanner that can scan an extract tokens from source
func NewScanner(source string) Scanner {
	return Scanner{source, strings.NewReader(source), make([]Token, 0), 0, 0, 1}
}

// ScanTokens scans the source code and returns a slice of Token from it. Returns an error if
// the source code contains syntax errors.
func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		if err := s.scanToken(); err != nil {
			return nil, err
		}
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))

	return s.tokens, nil
}

func (s Scanner) isAtEnd() bool {
	return s.current >= int64(len(s.source))
}

func (s *Scanner) scanToken() error {
	r, err := s.advance()
	if err != nil {
		return err
	}

	switch r {
	case '(':
		s.addEmptyToken(LeftParen)
	case ')':
		s.addEmptyToken(RightParen)
	case '{':
		s.addEmptyToken(LeftBrace)
	case '}':
		s.addEmptyToken(RightBrace)
	case ',':
		s.addEmptyToken(Comma)
	case '.':
		s.addEmptyToken(Dot)
	case '-':
		s.addEmptyToken(Minus)
	case '+':
		s.addEmptyToken(Plus)
	case ';':
		s.addEmptyToken(Semicolon)
	case '*':
		s.addEmptyToken(Star)
	default:
		return SyntaxError{
			line:    s.line,
			where:   "somewhere in the code",
			message: fmt.Sprintf("Unexpected character %s", string(r)),
		}
	}

	return nil
}

func (s *Scanner) advance() (r rune, err error) {
	_, err = s.reader.Seek(s.current, io.SeekStart)
	if err != nil {
		return
	}

	s.current++

	r, _, err = s.reader.ReadRune()

	return
}

func (s *Scanner) addEmptyToken(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}
