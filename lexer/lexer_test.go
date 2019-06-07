package lexer

import (
	"testing"

	"github.com/maiyama18/immonkey/token"
	"github.com/stretchr/testify/require"
)

func TestLexer_NextToken(t *testing.T) {
	tests := []struct {
		input          string
		expectedTokens []token.Token
	}{
		{
			input: `=+(){},;`,
			expectedTokens: []token.Token{
				{token.ASSIGN, "="},
				{token.PLUS, "+"},
				{token.LPAREN, "("},
				{token.RPAREN, ")"},
				{token.LBRACE, "{"},
				{token.RBRACE, "}"},
				{token.COMMA, ","},
				{token.SEMICOLON, ";"},
			},
		},
		{
			input: `
let five = 5;

let add = fn(x, y) {
	x + y;
};

let result = add(five, 10);
`,
			expectedTokens: []token.Token{
				{token.LET, "let"},
				{token.IDENTIFIER, "five"},
				{token.ASSIGN, "="},
				{token.INT, "5"},
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
				{token.IDENTIFIER, "five"},
				{token.COMMA, ","},
				{token.INT, "10"},
				{token.RPAREN, ")"},
				{token.SEMICOLON, ";"},
			},
		},
	}

	for _, test := range tests {
		l := New(test.input)
		for _, expected := range test.expectedTokens {
			actual := l.NextToken()
			require.Equal(t, expected, actual)
		}
	}
}
