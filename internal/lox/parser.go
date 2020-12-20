package lox

import (
	"fmt"
)

// TODO: add a helper for calling fmt.Errorf that appends current line number

// Parser parser
type Parser struct {
	tokens []Token
	current int
}

// NewParser create a parser from tokens
func NewParser(tokens []Token) *Parser {
	parser := Parser{}
	parser.tokens = tokens
	parser.current = 0
	return &parser
}

func (p *Parser) expression() (Expr, error) {
	return p.equality()
}

func (p *Parser) equality() (Expr, error) {
	expr, err := p.comparison()
	if err != nil {
		return nil, fmt.Errorf("Parse error getting equality expression: %w", err)
	}

	for p.match(BangEqual, EqualEqual) {
		operator := p.previous()
		right, err := p.comparison()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting equality expression: %w", err)
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) match(types... TokenType) bool {
	for _, tokenType := range types {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false;
}

func (p *Parser) check(typeOf TokenType) bool {
	if p.isEnd() {
		return false
	}

	return p.peek().typeOf == typeOf
}

func (p *Parser) advance() Token {
	if !p.isEnd() {
		p.current = p.current + 1
	}
	return p.previous()
}

func (p *Parser) isEnd() bool {
	return p.peek().typeOf == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() Token {
	return p.tokens[p.current - 1]
}

func (p *Parser) comparison() (Expr, error) {
	expr, err := p.term()
	if err != nil {
		return nil, fmt.Errorf("Parse error getting comparison expression: %w", err)
	}

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.previous()
		right, err := p.term()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting comparison expression: %w", err)
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) term() (Expr, error) {
	expr, err := p.factor()
	if err != nil {
		return nil, fmt.Errorf("Parse error getting term expression: %w", err)
	}

	for p.match(Minus, Plus) {
		operator := p.previous()
		right, err := p.factor()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting term expression: %w", err)
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) factor() (Expr, error) {
	expr, err := p.unary()
	if err != nil {
		return nil, fmt.Errorf("Parse error getting factor expression: %w", err)
	}

	for p.match(Slash, Star) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting factor expression: %w", err)
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) unary() (Expr, error) {
	if p.match(Bang, Minus) {
		operator := p.previous()
		right, err := p.unary()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting unary expression: %w", err)
		}

		return NewUnaryExpr(operator, right), nil
	}

	expr, err := p.primary()
	if err != nil {
		return nil, fmt.Errorf("Parse error getting primary expression: %w", err)
	}

	return expr, nil
}

func (p *Parser) primary() (Expr, error) {
	if p.match(False) {
		return NewLiteralExpr(false), nil
	}

	if p.match(True) {
		return NewLiteralExpr(true), nil
	}

	if p.match(Nil) {
		return NewLiteralExpr(nil), nil
	}

	if p.match(Number, String) {
		return NewLiteralExpr(p.previous().literal), nil
	}

	if p.match(LeftParen) {
		// TODO finish this and write Parser.consume method
		expr, err := p.expression()
		if err != nil {
			return nil, fmt.Errorf("Parse error getting primary expression: %w", err)
		}

		_, err = p.consume(RightParen, "Expect ')' after expression.")
		if err != nil {
			return nil, fmt.Errorf("Parse error getting primary expression: %w", err)
		}
		return NewGroupingExpr(expr), nil
	}

	return nil, nil
}

func (p *Parser) consume(typeOf TokenType, msg string) (Token, error) {
	if p.check(typeOf) {
		return p.advance(), nil
	}

	var tok Token
	return tok, fmt.Errorf("Parse error - peek: %s message: %s", p.peek(), msg)
}