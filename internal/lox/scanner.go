package lox

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	lineFeed       rune = '\x0A'
	carriageReturn rune = '\x0D'
	whitespace     rune = ' '
	horizontalTab  rune = '\x09'
	formFeed       rune = '\x0C'
	nullTerminator rune = '\x00'
)

var keywords map[string]TokenType

func init() {
	keywords = make(map[string]TokenType)

	keywords["and"] = And
	keywords["class"] = Class
	keywords["else"] = Else
	keywords["false"] = False
	keywords["for"] = For
	keywords["fun"] = Fun
	keywords["if"] = If
	keywords["nil"] = Nil
	keywords["or"] = Or
	keywords["print"] = Print
	keywords["return"] = Return
	keywords["super"] = Super
	keywords["this"] = This
	keywords["true"] = True
	keywords["var"] = Var
	keywords["while"] = While
}

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

	s.tokens = append(s.tokens, *NewToken(EOF, "", nil, s.line))

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
					if _, err = s.advance(); err != nil {
						return err
					}
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
		if s.isDigit(r) {
			return s.number()
		}

		if s.isAlpha(r) {
			return s.identifier()
		}

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
	r = nullTerminator

	if s.isAtEnd() {
		return
	}

	r, err = s.advance()

	s.current--

	return
}

func (s *Scanner) peekNext() (r rune, err error) {
	if r, err = s.advance(); err != nil {
		return
	}

	r, err = s.peek()

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

	if s.isAtEnd() {
		return SyntaxError{
			line:    s.line,
			where:   s.source[s.start:s.current],
			message: "Unterminated string",
		}
	}

	// NOTE (c.nicola): Advances to the closing "
	if _, err := s.advance(); err != nil {
		return err
	}

	// NOTE (c.nicola): Trim the surrounding " from the string and add the token
	s.addToken(String, s.source[s.start+1:s.current-1])

	return nil
}

func (s *Scanner) isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func (s *Scanner) isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func (s *Scanner) isAlphaNumeric(r rune) bool {
	return s.isAlpha(r) || s.isDigit(r)
}

func (s *Scanner) number() error {
	advanceUntilNonDigit := func() error {
		for {
			r, err := s.peek()
			if err != nil {
				return err
			}

			if !s.isDigit(r) {
				return nil
			}

			if _, err = s.advance(); err != nil {
				return err
			}
		}
	}

	if err := advanceUntilNonDigit(); err != nil {
		return err
	}

	// NOTE (c.nicola): Check if this is a floating point number
	r, err := s.peek()
	if err != nil {
		return err
	}

	if r == '.' {
		r, err = s.peekNext()
		if err != nil {
			return err
		}

		if s.isDigit(r) {
			if _, err = s.advance(); err != nil {
				return err
			}

			if err = advanceUntilNonDigit(); err != nil {
				return err
			}
		}
	}

	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return err
	}

	s.addToken(Number, f)

	return nil
}

func (s *Scanner) identifier() error {
	for {
		r, err := s.peek()
		if err != nil {
			return err
		}

		if !s.isAlpha(r) {
			break
		}

		if _, err = s.advance(); err != nil {
			return err
		}
	}

	text := s.source[s.start:s.current]

	if k, ok := keywords[text]; ok {
		s.addEmptyToken(k)
	} else {
		s.addToken(Identifier, text)
	}

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

	s.tokens = append(s.tokens, *NewToken(tokenType, text, literal, s.line))
}
