package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

func New(tokenType Type, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	// identifier & literals
	IDENTIFIER Type = "IDENTIFIER"
	INT        Type = "INT"

	// operators
	ASSIGN Type = "="
	PLUS   Type = "+"

	// delimiters
	COMMA     Type = ","
	SEMICOLON Type = ";"

	LPAREN Type = "("
	RPAREN Type = ")"
	LBRACE Type = "{"
	RBRACE Type = "}"

	// keywords
	FUNCTION Type = "FUNCTION"
	LET      Type = "LET"
)
