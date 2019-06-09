package parser

import (
	"fmt"

	"github.com/maiyama18/immonkey/ast"
	"github.com/maiyama18/immonkey/lexer"
	"github.com/maiyama18/immonkey/token"
)

type Parser struct {
	lxr *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []error
}

func New(lxr *lexer.Lexer) *Parser {
	p := &Parser{lxr: lxr}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []error {
	return p.errors
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lxr.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	if !p.expectPeekTokenType(token.IDENTIFIER) {
		return nil
	}

	ident := &ast.Identifier{Token: p.currentToken, Name: p.currentToken.Literal}

	if !p.expectPeekTokenType(token.ASSIGN) {
		return nil
	}

	// TODO: parse expression value
	for !p.isCurrentTokenType(token.SEMICOLON) {
		p.nextToken()
	}

	return &ast.LetStatement{Token: token.New(token.LET, "let"), Identifier: ident}
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	// TODO: parse expression value
	for !p.isCurrentTokenType(token.SEMICOLON) {
		p.nextToken()
	}

	return &ast.ReturnStatement{Token: token.New(token.RETURN, "return")}
}

func (p *Parser) expectPeekTokenType(tokenType token.Type) bool {
	if !p.isPeekTokenType(tokenType) {
		p.addPeekError(tokenType)
		return false
	}
	p.nextToken()
	return true
}

func (p *Parser) isPeekTokenType(tokenType token.Type) bool {
	return p.peekToken.Type == tokenType
}

func (p *Parser) isCurrentTokenType(tokenType token.Type) bool {
	return p.currentToken.Type == tokenType
}

func (p *Parser) addPeekError(tokenType token.Type) {
	p.errors = append(p.errors, fmt.Errorf("expect next token to be %s, but got %s", tokenType, p.peekToken.Type))
}
