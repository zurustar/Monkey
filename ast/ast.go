//
// ASTとはAbstruct Syntax Treeの略で、日本語では抽象構文木という。
// プログラム全体を木構造で持つもの。
//
//

package ast

import (
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
}

// 式
// 式は値を返す
type Statement interface {
	Node
	statementNode()
}

// 文
// 分は値を返さない
type Expression interface {
	Node
	expressionNode()
}

// プログラム全体のASTのルート
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


type LetStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal}