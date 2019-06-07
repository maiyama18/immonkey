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
	}

	for _, test := range tests {
		l := New(test.input)
		for _, expected := range test.expectedTokens {
			actual := l.NextToken()
			require.Equal(t, expected, actual)
		}
	}
}
