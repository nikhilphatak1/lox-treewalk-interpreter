package lox

import (
	"fmt"
)

// TokenType types of tokens
type TokenType int

// Token types
const (
	// Single-character tokens.
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
  
	// One or two character tokens.
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
  
	// Literals.
	Identifier
	String
	Number
  
	// Keywords.
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	EOF
)

// Token token
type Token struct {
	typeOf TokenType
	lexeme string
	literal interface{}
	line int
}

// NewToken Token constructor
func NewToken(typeOf TokenType, lexeme string, literal interface{}, line int) *Token {
	token := Token{}
	token.typeOf = typeOf
	token.lexeme = lexeme
	token.literal = literal
	token.line = line
	return &token
}

func (t *Token) String() string {
	typeString := [...]string{
		"LeftParen", "RightParen", "LeftBrace", "RightBrace", "Comma", "Dot",
		"Minus", "Plus", "Semicolon", "Slash", "Star", "Bang", "BangEqual",
		"Equal", "EqualEqual", "Greater", "GreaterEqual", "Less", "LessEqual",
		"Identifier", "String", "Number", "And", "Class", "Else", "False",
		"Fun", "For", "If", "Nil", "Or", "Print", "Return", "Super", "This",
		"True", "Var", "While", "EOF"}[t.typeOf]
	return fmt.Sprintf("%s %s %v", typeString, t.lexeme, t.literal)
}