package lexer

import (
	"bytes"

	"github.com/maiyama18/immonkey/token"
)

type Lexer struct {
	input        string
	position     int
	peekPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.consumeChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipSpaces()

	var tk token.Token
	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			tk = token.New(token.EQ, "==")
			l.consumeChar()
		} else {
			tk = token.New(token.ASSIGN, "=")
		}
	case '!':
		if l.peekChar() == '=' {
			tk = token.New(token.NOTEQ, "!=")
			l.consumeChar()
		} else {
			tk = token.New(token.BANG, "!")
		}
	case '<':
		tk = token.New(token.LT, "<")
	case '>':
		tk = token.New(token.GT, ">")
	case '+':
		tk = token.New(token.PLUS, "+")
	case '-':
		tk = token.New(token.MINUS, "-")
	case '*':
		tk = token.New(token.ASTERISK, "*")
	case '/':
		tk = token.New(token.SLASH, "/")
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
	default:
		if isLetter(l.char) {
			literal := l.readName()
			return token.New(token.TypeOf(literal), literal)
		} else if isDigit(l.char) {
			literal := l.readNumber()
			return token.New(token.INT, literal)
		} else {
			return token.New(token.ILLEGAL, string(l.char))
		}
	}

	l.consumeChar()
	return tk
}

func (l *Lexer) consumeChar() {
	if l.peekPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.peekPosition]
	}
	l.position = l.peekPosition
	l.peekPosition++
}

func (l *Lexer) peekChar() byte {
	if l.peekPosition >= len(l.input) {
		return 0
	}
	return l.input[l.peekPosition]
}

func (l *Lexer) readName() string {
	var b bytes.Buffer
	for isLetter(l.char) {
		b.WriteByte(l.char)
		l.consumeChar()
	}
	return b.String()
}

func (l *Lexer) readNumber() string {
	var b bytes.Buffer
	for isDigit(l.char) {
		b.WriteByte(l.char)
		l.consumeChar()
	}
	return b.String()
}

func (l *Lexer) skipSpaces() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.consumeChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && 'z' <= char || 'A' <= char && 'Z' <= char || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
