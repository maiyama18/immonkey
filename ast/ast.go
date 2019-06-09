package ast

import (
	"strings"

	"github.com/maiyama18/immonkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) == 0 {
		return ""
	}
	return p.Statements[0].TokenLiteral()
}
func (p *Program) String() string {
	var out strings.Builder

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token      token.Token
	Identifier *Identifier
	Expression Expression
}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) String() string {
	var out strings.Builder

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Identifier.Name + " = ")
	if ls.Expression != nil {
		out.WriteString(ls.Expression.String())
	}
	out.WriteString(";")

	return out.String()
}

type ReturnStatement struct {
	Token      token.Token
	Expression Expression
}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.Expression != nil {
		out.WriteString(rs.Expression.String())
	}
	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) String() string {
	if es.Expression == nil {
		return ""
	}
	return es.Expression.String()
}

type Identifier struct {
	Token token.Token
	Name  string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string       { return i.Name }

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) TokenLiteral() string { return i.Token.Literal }
func (i *IntegerLiteral) expressionNode()      {}
func (i *IntegerLiteral) String() string       { return i.Token.Literal }
