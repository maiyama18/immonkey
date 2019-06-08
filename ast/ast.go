package ast

import "github.com/maiyama18/immonkey/token"

type Node interface {
	Token() token.Token
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

func (p *Program) Token() token.Token {
	if len(p.Statements) == 0 {
		return token.New(token.EOF, "")
	}
	return p.Statements[0].Token()
}

type LetStatement struct {
	tk         token.Token
	Identifier *Identifier
	Value      Expression
}

func (ls *LetStatement) Token() token.Token { return ls.tk }
func (ls *LetStatement) statementNode()     {}

type Identifier struct {
	tk   token.Token
	Name string
}

func (i *Identifier) Token() token.Token { return i.tk }
func (i *Identifier) expressionNode()    {}
