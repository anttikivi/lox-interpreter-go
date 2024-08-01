package main

import "fmt"

type TokenType uint

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	STRING
	NUMBER

	EOF
)

var tokenNames = map[TokenType]string{
	LEFT_PAREN:  "LEFT_PAREN",
	RIGHT_PAREN: "RIGHT_PAREN",
	LEFT_BRACE:  "LEFT_BRACE",
	RIGHT_BRACE: "RIGHT_BRACE",
	COMMA:       "COMMA",
	DOT:         "DOT",
	MINUS:       "MINUS",
	PLUS:        "PLUS",
	SEMICOLON:   "SEMICOLON",
	SLASH:       "SLASH",
	STAR:        "STAR",

	BANG:          "BANG",
	BANG_EQUAL:    "BANG_EQUAL",
	EQUAL:         "EQUAL",
	EQUAL_EQUAL:   "EQUAL_EQUAL",
	GREATER:       "GREATER",
	GREATER_EQUAL: "GREATER_EQUAL",
	LESS:          "LESS",
	LESS_EQUAL:    "LESS_EQUAL",

	STRING: "STRING",
	NUMBER: "NUMBER",

	EOF: "EOF",
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %v", tokenNames[t.Type], t.Lexeme, t.Literal)
}
