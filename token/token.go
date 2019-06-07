package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOR     = "EOF"

	// identifier & literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
