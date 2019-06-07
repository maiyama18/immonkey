package lexer

import (
	"testing"

	"github.com/maiyama18/immonkey/token"
	"github.com/stretchr/testify/require"
)

func TestLexer_NextToken(t *testing.T) {
	input := `=+(){},;`
	l := New(input)

	expectedTokens := []token.Token{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	for _, expected := range expectedTokens {
		actual := l.NextToken()
		require.Equal(t, expected, actual)
	}
}
