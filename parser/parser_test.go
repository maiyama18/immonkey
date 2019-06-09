package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/maiyama18/immonkey/token"

	"github.com/maiyama18/immonkey/ast"
	"github.com/maiyama18/immonkey/lexer"
	"github.com/stretchr/testify/require"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 42;
let foo = 1;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	require.NotNil(t, program)
	require.Equal(t, 2, len(program.Statements))

	testLetStatement(t, program.Statements[0], "x")
	testLetStatement(t, program.Statements[1], "foo")
}

func TestParser_Errors(t *testing.T) {
	input := `
let x 42;
let = 1;
let 99;
`

	l := lexer.New(input)
	p := New(l)

	_ = p.ParseProgram()
	errs := p.Errors()

	require.NotNil(t, errs)
	require.Equal(t, 3, len(errs))

	tokenTypes := []token.Type{token.ASSIGN, token.IDENTIFIER, token.IDENTIFIER}
	for i, tokenType := range tokenTypes {
		err := errs[i]
		expectMsg := fmt.Sprintf("expect next token to be %s", tokenType)
		require.True(t, strings.Contains(err.Error(), expectMsg),
			"expect '%s' to contain '%s'", err.Error(), expectMsg)
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) {
	t.Helper()

	require.Equal(t, "let", stmt.TokenLiteral())

	letStmt, ok := stmt.(*ast.LetStatement)
	require.True(t, ok)

	require.Equal(t, name, letStmt.Identifier.Name)
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()
	if len(errs) == 0 {
		return
	}

	t.Errorf("%d parser errors", len(errs))
	for _, err := range errs {
		t.Error(err)
	}
	t.FailNow()
}
