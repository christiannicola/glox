package lexer

import (
	"io"
	"strings"
)

type Scanner struct {
	source string
	reader *strings.Reader
	tokens []Token
	start int64
	current int64
	line int64
}

func NewScanner(source string) Scanner {
	return Scanner{source, strings.NewReader(source), make([]Token, 0), 0, 0, 1}
}


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
	case '(': s.addEmptyToken(LEFT_PAREN)
	case ')': s.addEmptyToken(RIGHT_PAREN)
	case '{': s.addEmptyToken(LEFT_BRACE)
	case '}': s.addEmptyToken(RIGHT_BRACE)
	case ',': s.addEmptyToken(COMMA)
	case '.': s.addEmptyToken(DOT)
	case '-': s.addEmptyToken(MINUS)
	case '+': s.addEmptyToken(PLUS)
	case ';': s.addEmptyToken(SEMICOLON)
	case '*': s.addEmptyToken(STAR)
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


