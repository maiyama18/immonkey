package lexer

import (
	"testing"

	"github.com/maiyama18/immonkey/token"
	"github.com/stretchr/testify/require"
)

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedTokens []token.Token
	}{
		{
			name:  "single-char tokens",
			input: `=!+-*/><(){},;`,
			expectedTokens: []token.Token{
				{token.ASSIGN, "="},
				{token.BANG, "!"},
				{token.PLUS, "+"},
				{token.MINUS, "-"},
				{token.ASTERISK, "*"},
				{token.SLASH, "/"},
				{token.GT, ">"},
				{token.LT, "<"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
			},
		},
		{
			name: "multi-char tokens",
			input: `
10 == 10;
5 != 10;
`,
			expectedTokens: []token.Token{
				{token.INT, "10"},
				{token.EQ, "=="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},

				{token.INT, "5"},
				{token.NOTEQ, "!="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},
			},
		},
		{
			name: "identifiers/keywords",
			input: `
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(5, ten);

if (5 < 10) {
	return true;
} else {
	return false;
}
`,
			expectedTokens: []token.Token{
				{token.LET, "let"},
				{token.IDENTIFIER, "ten"},
				{token.ASSIGN, "="},
				{token.INT, "10"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENTIFIER, "add"},
				{token.ASSIGN, "="},
				{token.FUNCTION, "fn"},
				{token.LPAREN, "("},
				{token.IDENTIFIER, "x"},
				{token.COMMA, ","},
				{token.IDENTIFIER, "y"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.IDENTIFIER, "x"},
				{token.PLUS, "+"},
				{token.IDENTIFIER, "y"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.SEMICOLON, ";"},

				{token.LET, "let"},
				{token.IDENTIFIER, "result"},
				{token.ASSIGN, "="},
				{token.IDENTIFIER, "add"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.COMMA, ","},
				{token.IDENTIFIER, "ten"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},

				{token.IF, "if"},
				{token.LPAREN, "("},
				{token.INT, "5"},
				{token.LT, "<"},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.TRUE, "true"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
				{token.ELSE, "else"},
				{token.LBRACE, "{"},
				{token.RETURN, "return"},
				{token.FALSE, "false"},
				{token.SEMICOLON, ";"},
				{token.RBRACE, "}"},
			},
		},
	}

	for _, test := range tests {
		l := New(test.input)
		for i, expected := range test.expectedTokens {
			actual := l.NextToken()
			require.Equal(t, expected, actual, "%d-th token wrong", i)
		}
	}
}
