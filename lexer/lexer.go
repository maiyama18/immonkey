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
	return token.Token{}
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
