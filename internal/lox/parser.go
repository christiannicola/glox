package lox

import (
	"fmt"
	"github.com/christiannicola/glox/internal/lox/ast"
)

var errNoPrevious = fmt.Errorf("no previous token available")

// Parser is responsible for constructing an AST out of Tokens.
type Parser struct {
	tokens  []ast.Token
	current int64
}

// NewParser returns a pointer to a parser.
func NewParser(tokens []ast.Token) *Parser {
	return &Parser{tokens, 0}
}

// Parse parses a single expression and returns it. Can also return a ParseError if the tokens
// do not adhere to the language grammar.
func (p *Parser) Parse() (ast.Expression, error) {
	return p.expression()
}

func (p *Parser) expression() (ast.Expression, error) {
	return p.equality()
}

func (p *Parser) equality() (ast.Expression, error) {
	var (
		matchesType bool
		expr        ast.Expression
		operator    *ast.Token
		right       ast.Expression
		err         error
	)

	if expr, err = p.comparison(); err != nil {
		return nil, err
	}

	for {
		if matchesType, err = p.match(ast.BangEqual, ast.EqualEqual); err != nil {
			return nil, err
		}

		if !matchesType {
			break
		}

		if operator, err = p.previous(); err != nil {
			return nil, err
		}

		if right, err = p.comparison(); err != nil {
			return nil, err
		}

		expr = ast.NewBinary(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) comparison() (ast.Expression, error) {
	var (
		matchesType bool
		expr        ast.Expression
		operator    *ast.Token
		right       ast.Expression
		err         error
	)

	if expr, err = p.term(); err != nil {
		return nil, err
	}

	for {
		if matchesType, err = p.match(ast.Greater, ast.GreaterEqual, ast.Less, ast.LessEqual); err != nil {
			return nil, err
		}

		if !matchesType {
			break
		}

		if operator, err = p.previous(); err != nil {
			return nil, err
		}

		if right, err = p.term(); err != nil {
			return nil, err
		}

		expr = ast.NewBinary(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) term() (ast.Expression, error) {
	var (
		matchesType bool
		expr        ast.Expression
		operator    *ast.Token
		right       ast.Expression
		err         error
	)

	if expr, err = p.factor(); err != nil {
		return nil, err
	}

	for {
		if matchesType, err = p.match(ast.Minus, ast.Plus); err != nil {
			return nil, err
		}

		if !matchesType {
			break
		}

		if operator, err = p.previous(); err != nil {
			return nil, err
		}

		if right, err = p.factor(); err != nil {
			return nil, err
		}

		expr = ast.NewBinary(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) factor() (ast.Expression, error) {
	var (
		matchesType bool
		expr        ast.Expression
		operator    *ast.Token
		right       ast.Expression
		err         error
	)

	if expr, err = p.unary(); err != nil {
		return nil, err
	}

	for {
		if matchesType, err = p.match(ast.Slash, ast.Star); err != nil {
			return nil, err
		}

		if !matchesType {
			break
		}

		if operator, err = p.previous(); err != nil {
			return nil, err
		}

		if right, err = p.unary(); err != nil {
			return nil, err
		}

		expr = ast.NewBinary(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) unary() (ast.Expression, error) {
	var (
		matchesType bool
		err         error
		operator    *ast.Token
		right       ast.Expression
	)

	if matchesType, err = p.match(ast.Bang, ast.Minus); err != nil {
		return nil, err
	}

	if matchesType {
		if operator, err = p.previous(); err != nil {
			return nil, err
		}

		if right, err = p.unary(); err != nil {
			return nil, err
		}

		return ast.NewUnary(operator, right), nil
	}

	return p.primary()
}

func (p *Parser) primary() (ast.Expression, error) {
	switch p.peek().Type() {
	case ast.False:
		p.current++
		return ast.NewLiteral(false), nil
	case ast.True:
		p.current++
		return ast.NewLiteral(true), nil
	case ast.Nil:
		p.current++
		return ast.NewLiteral(nil), nil
	case ast.Number:
		fallthrough
	case ast.String:
		p.current++
		prev, err := p.previous()

		if err != nil {
			return nil, err
		}

		return ast.NewLiteral(prev.Literal()), nil
	case ast.LeftParen:
		p.current++

		expr, err := p.expression()

		if err != nil {
			return nil, err
		}

		if _, err := p.consume(ast.RightParen, "Expect ')' after expression."); err != nil {
			return nil, err
		}

		return ast.NewGrouping(expr), nil
	default:
		return nil, p.error(p.peek(), "Expect expression.")
	}
}

func (p *Parser) consume(t ast.TokenType, msg string) (*ast.Token, error) {
	if p.check(t) {
		return p.advance()
	}

	token := p.peek()

	return token, p.error(token, msg)
}

func (p Parser) error(t *ast.Token, msg string) error {
	err := ParseError{
		token:   t,
		line:    t.Line(),
		where:   t.Lexeme(),
		message: msg,
	}

	return err
}

func (p *Parser) synchronize() error {
	var (
		prev *ast.Token
		err  error
	)
	if _, err = p.advance(); err != nil {
		return err
	}

	for !p.isAtEnd() {
		if prev, err = p.previous(); err != nil {
			return err
		}

		if prev.Type() == ast.Semicolon {
			return nil
		}

		switch p.peek().Type() {
		case ast.Class:
		case ast.For:
		case ast.Fun:
		case ast.If:
		case ast.Print:
		case ast.Return:
		case ast.Var:
		case ast.While:
		default:
			if _, err = p.advance(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *Parser) match(types ...ast.TokenType) (bool, error) {
	for i := range types {
		if p.check(types[i]) {
			_, err := p.advance()

			return true, err
		}
	}

	return false, nil
}

func (p Parser) check(t ast.TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().Type() == t
}

func (p *Parser) advance() (*ast.Token, error) {
	if !p.isAtEnd() {
		p.current++
	}

	return p.previous()
}

func (p Parser) isAtEnd() bool {
	return p.peek().Type() == ast.EOF
}

func (p Parser) peek() *ast.Token {
	return &p.tokens[p.current]
}

func (p Parser) previous() (*ast.Token, error) {
	if p.current-1 < 0 {
		return nil, errNoPrevious
	}

	return &p.tokens[p.current-1], nil
}
