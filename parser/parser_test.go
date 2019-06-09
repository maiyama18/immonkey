package parser

import (
	"testing"

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
	require.NotNil(t, program)
	require.Equal(t, 2, len(program.Statements))

	testLetStatement(t, program.Statements[0], "x")
	testLetStatement(t, program.Statements[1], "foo")
}

func testLetStatement(t *testing.T, stmt ast.Statement, name string) {
	t.Helper()

	require.Equal(t, "let", stmt.TokenLiteral())

	letStmt, ok := stmt.(*ast.LetStatement)
	require.True(t, ok)

	require.Equal(t, name, letStmt.Identifier.Name)
}
