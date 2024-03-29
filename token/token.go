package token

const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	// identifier & literals
	IDENTIFIER Type = "IDENTIFIER"
	INT        Type = "INT"

	// operators
	ASSIGN   Type = "="
	PLUS     Type = "+"
	MINUS    Type = "-"
	ASTERISK Type = "*"
	SLASH    Type = "/"

	BANG  Type = "!"
	LT    Type = "<"
	GT    Type = ">"
	EQ    Type = "=="
	NOTEQ Type = "!="

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
	IF       Type = "IF"
	ELSE     Type = "ELSE"
	RETURN   Type = "RETURN"
	TRUE     Type = "true"
	FALSE    Type = "false"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}

func New(tokenType Type, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

func TypeOf(literal string) Type {
	switch literal {
	case "let":
		return LET
	case "fn":
		return FUNCTION
	case "if":
		return IF
	case "else":
		return ELSE
	case "return":
		return RETURN
	case "true":
		return TRUE
	case "false":
		return FALSE
	default:
		return IDENTIFIER
	}
}
