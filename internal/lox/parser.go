package lox

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

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.previous()
		right := p.comparison()
		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr
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

func (p *Parser) comparison() Expr {
	expr := p.term()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.previous()
		right := p.term()
		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for p.match(Minus, Plus) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.factor()

	for p.match(Slash, Star) {
		operator := p.previous()
		right := p.factor()
		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(Bang, Minus) {
		operator := p.previous()
		right := p.unary()
		return NewUnaryExpr(operator, right)
	}
	
	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match(False) {
		return NewLiteralExpr(false)
	}

	if p.match(True) {
		return NewLiteralExpr(true)
	}

	if p.match(Nil) {
		return NewLiteralExpr(nil)
	}

	if p.match(Number, String) {
		return NewLiteralExpr(p.previous().literal)
	}

	if p.match(LeftParen) {
		// TODO finish this and write Parser.consume method
		expr := p.expression()
		p.consume(RightParen, "Expect ')' after expression.")
		return NewGroupingExpr()
	}
}