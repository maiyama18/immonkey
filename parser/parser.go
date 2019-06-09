package parser

import (
	"fmt"
	"strconv"

	"github.com/maiyama18/immonkey/ast"
	"github.com/maiyama18/immonkey/lexer"
	"github.com/maiyama18/immonkey/token"
)

const (
	LOWEST      = iota
	EQUALS      // ==, !=
	LESSGREATER // >, <
	SUM         // +, -
	PRODUCT     // *, /
	PREFIX      // -X, !X
	CALL        // add(1, 2)
)

type (
	parsePrefixFn func() ast.Expression
	parseInfixFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lxr *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	errors []error

	parsePrefixFns map[token.Type]parsePrefixFn
	parseInfixFns  map[token.Type]parseInfixFn
}

func New(lxr *lexer.Lexer) *Parser {
	p := &Parser{lxr: lxr}

	p.nextToken()
	p.nextToken()

	p.parsePrefixFns = map[token.Type]parsePrefixFn{
		token.IDENTIFIER: p.parseIdentifier,
		token.INT:        p.parseIntegerLiteral,
	}
	p.parseInfixFns = map[token.Type]parseInfixFn{}

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
		return p.parseExpressionStatement()
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
	p.nextToken()

	// TODO: parse expression value
	for !p.isCurrentTokenType(token.SEMICOLON) {
		p.nextToken()
	}

	return &ast.ReturnStatement{Token: token.New(token.RETURN, "return")}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	tk := p.currentToken

	exp := p.parseExpression(LOWEST)

	if p.isPeekTokenType(token.SEMICOLON) {
		p.nextToken()
	}

	return &ast.ExpressionStatement{Token: tk, Expression: exp}
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefixFn, ok := p.parsePrefixFns[p.currentToken.Type]
	if !ok {
		p.addError("no function found to parse %s", p.currentToken.Type)
		return nil
	}
	left := prefixFn()

	return left
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Name: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	v, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		p.addError("could not parse '%s' as integer", p.currentToken.Literal)
		return nil
	}
	return &ast.IntegerLiteral{Token: p.currentToken, Value: v}
}

func (p *Parser) expectPeekTokenType(tokenType token.Type) bool {
	if !p.isPeekTokenType(tokenType) {
		p.addError("expect next token to be %s, but got %s", tokenType, p.peekToken.Type)
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

func (p *Parser) addError(format string, a ...interface{}) {
	p.errors = append(p.errors, fmt.Errorf(format, a...))
}
