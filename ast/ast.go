package ast

import "monkey/token"
// Node should return the literal value of the token
// TokenLiteral() will be used only for debugging and testing
type Node interface {
	TokenLiteral() string
}

// Statement is not always necessary
type Statement interface {
	Node
	statementNode()
}

// Expression is not always necessary
type Expression interface {
	Node
	expressionNode()
}

// Program  is the root node of every AST
type Program struct {
	Statements []Statement
}

// TokenLiteral return
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls* LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}