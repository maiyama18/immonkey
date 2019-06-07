package lexer

import "github.com/maiyama18/immonkey/token"

type Lexer struct {
	input        string
	position     int
	peekPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	switch l.char {
	case '=':
		tk = token.New(token.ASSIGN, "=")
	case '+':
		tk = token.New(token.PLUS, "+")
	case '(':
		tk = token.New(token.LPAREN, "(")
	case ')':
		tk = token.New(token.RPAREN, ")")
	case '{':
		tk = token.New(token.LBRACE, "{")
	case '}':
		tk = token.New(token.RBRACE, "}")
	case ',':
		tk = token.New(token.COMMA, ",")
	case ';':
		tk = token.New(token.SEMICOLON, ";")
	case 0:
		tk = token.New(token.EOF, "")
	}

	l.readChar()
	return tk
}

func (l *Lexer) readChar() {
	if l.peekPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.peekPosition]
	}
	l.position = l.peekPosition
	l.peekPosition++
}
