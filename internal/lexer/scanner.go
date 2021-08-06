package lexer

import (
	"fmt"
	"io"
	"strings"
)

const (
	lineFeed       rune = '\x0A'
	carriageReturn rune = '\x0D'
	whitespace     rune = ' '
	horizontalTab  rune = '\x09'
	formFeed       rune = '\x0C'
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
	case '"':
		return s.string()
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
	case '!':
		isBangEqual, err := s.matchNext('=')
		if err != nil {
			return err
		}

		if isBangEqual {
			s.addEmptyToken(BangEqual)
		} else {
			s.addEmptyToken(Bang)
		}
	case '=':
		isEqualEqual, err := s.matchNext('=')
		if err != nil {
			return err
		}

		if isEqualEqual {
			s.addEmptyToken(EqualEqual)
		} else {
			s.addEmptyToken(Equal)
		}
	case '<':
		isLessEqual, err := s.matchNext('=')
		if err != nil {
			return err
		}

		if isLessEqual {
			s.addEmptyToken(LessEqual)
		} else {
			s.addEmptyToken(Less)
		}
	case '>':
		isGreaterEqual, err := s.matchNext('=')
		if err != nil {
			return err
		}

		if isGreaterEqual {
			s.addEmptyToken(GreaterEqual)
		} else {
			s.addEmptyToken(Greater)
		}
	case '/':
		isComment, err := s.matchNext('/')

		if err != nil {
			return err
		}

		if !isComment {
			s.addEmptyToken(Slash)
		} else {
			for {
				if s.isAtEnd() {
					break
				}

				newLine, err := s.peek()
				if err != nil {
					return err
				}

				if newLine != lineFeed {
					s.advance()
				}
			}
		}
	case lineFeed:
		s.line++
	case carriageReturn:
	case whitespace:
	case horizontalTab:
	case formFeed:
	default:
		return SyntaxError{
			line:    s.line,
			where:   "somewhere in the code",
			message: fmt.Sprintf("Unexpected character %s", string(r)),
		}
	}

	return nil
}

func (s *Scanner) seek(pos int64) error {
	_, err := s.reader.Seek(pos, io.SeekStart)

	return err
}

func (s *Scanner) peek() (r rune, err error) {
	r = '\x00'

	if s.isAtEnd() {
		return
	}

	r, err = s.advance()

	s.current--

	return
}

func (s *Scanner) advance() (r rune, err error) {
	if err = s.seek(s.current); err != nil {
		return
	}

	s.current++

	r, _, err = s.reader.ReadRune()

	return
}

func (s *Scanner) matchNext(expected rune) (bool, error) {
	if s.isAtEnd() {
		return false, nil
	}

	actual, err := s.advance()

	if actual != expected {
		s.current--

		return false, err
	}

	s.current++

	return true, err
}

func (s *Scanner) string() error {
	for !s.isAtEnd() {
		nextRune, err := s.peek()
		if err != nil {
			return err
		}

		if nextRune == '"' {
			break
		}

		if nextRune == lineFeed {
			s.line++
		}

		if _, err = s.advance(); err != nil {
			return err
		}
	}

	// NOTE (c.nicola): Advances to the closing "
	if _, err := s.advance(); err != nil {
		return err
	}

	// NOTE (c.nicola): Trim the surrounding " from the string and add the token
	s.addToken(String, s.source[s.start + 1:s.current - 1])

	return nil
}

func (s *Scanner) addEmptyToken(tokenType TokenType) {
	s.addToken(tokenType, nil)
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	var text string

	if s.current > int64(len(s.source)) {
		text = s.source[s.start : len(s.source)-1]
	} else {
		text = s.source[s.start:s.current]
	}

	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}
