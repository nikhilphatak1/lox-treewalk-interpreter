package lox

import (
	"fmt"
	"strconv"
	"unicode"
)

// Scanner scan tokens
type Scanner struct {
	source string
	tokens []Token
	start int
	current int
	line int
}

// NewScanner Scanner constructor
func NewScanner(source string) *Scanner {
	scanner := Scanner{}
	scanner.source = source
	scanner.tokens = make([]Token, 0)
	return &scanner
}

// ScanTokens scan tokens
func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return nil, err
		}
	}

	s.tokens = append(s.tokens, *NewToken(EOF, "", nil, s.line))
	return s.tokens, nil
}

func (s *Scanner) isEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() error {
	var err error = nil
	switch c := s.advance(); c {
	case '(': s.addToken(LeftParen, nil)
	case ')': s.addToken(RightParen, nil)
	case '{': s.addToken(LeftBrace, nil)
	case '}': s.addToken(RightBrace, nil)
	case ',': s.addToken(Comma, nil)
	case '.': s.addToken(Dot, nil)
	case '-': s.addToken(Minus, nil)
	case '+': s.addToken(Plus, nil)
	case ';': s.addToken(Semicolon, nil)
	case '*': s.addToken(Star, nil)
	case '!':
		if s.match("=") {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}
	case '=':
		if s.match("=") {
			s.addToken(EqualEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}
	case '<':
		if s.match("=") {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Less, nil)
		}
	case '>':
		if s.match("=") {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		}
	case '/':
		if s.match("/") {
			for s.peek() != '\n' && !s.isEnd() {
				s.advance()
			}
		} else {
			s.addToken(Slash, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n': s.line = s.line + 1
	case '"':
		err = s.string()
	default:
		if unicode.IsDigit(c) {
			err = s.number()
		} else if unicode.IsLetter(c) {
			s.identifier()
		} else {
			return fmt.Errorf("Unexpected character at line %d", s.line)
		}
	}

	return err
}

func (s *Scanner) match(expected string) bool {
	if s.isEnd() {
		return false
	}

	if string(s.source[s.current]) != expected {
		return false
	}

	s.current = s.current + 1
	return true
}

func (s *Scanner) peek() rune {
	if s.isEnd() {
		// golang needs 3 octal digits for a null character
		return '\x00'
	}
	return rune(s.source[s.current])
}

func (s *Scanner) peekAnother() rune {
	if s.current + 1 >= len(s.source) {
		// null character
		return '\x00'
	}
	return rune(s.source[s.current + 1])
}

func (s *Scanner) identifier() {
	c := s.peek();
	for unicode.IsDigit(c) || unicode.IsLetter(c) {
		s.advance()
		c = s.peek()
	}
	s.addToken(Identifier, nil)
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isEnd() {
		if s.peek() == '\n' {
			s.line  = s.line + 1
		}
		s.advance()
	}

	if s.isEnd() {
		return fmt.Errorf("String not terminated at line %d", s.line)
	}

	s.advance()

	// get contents of string
	s.addToken(String, string(s.source[s.start + 1:s.current - 1]))

	return nil
}

func (s *Scanner) number() error {
	for (unicode.IsDigit(s.peek())) {
		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(s.peekAnother()) {
		s.advance()
		for unicode.IsDigit(s.peek()) {
			s.advance()
		}
	}

	num, err := strconv.ParseFloat(string(s.source[s.start:s.current]), 64)
	if err != nil {
		return err
	}
	s.addToken(Number, num)
	return nil
}

func (s *Scanner) advance() rune {
	s.current = s.current + 1
	return rune(s.source[s.current - 1])
}

func (s *Scanner) addToken(typeOf TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, *NewToken(typeOf, text, literal, s.line))
}
