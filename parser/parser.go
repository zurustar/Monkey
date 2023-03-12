// プログラムのソースコードを解析してASTを構築する

package parser

import (
	"fmt"
	"log"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer
	curToken token.Token // 今のトークン
	peekToken token.Token // 次のトークンを覗き見る(peek)

	errors []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Errors() []string{
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("exprected next token to be %s, but %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// 次のトークンに移動
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// プログラムを解析する。
func (p *Parser) ParseProgram() *ast.Program {
	// 抽象構文木を準備
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// 次のトークンが最後のものになるまでループしてみていく
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// 現在見ているトークンのステートメントを解析する
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		// letだったらletとしての解析を実行
		log.Println("現在のTOKENはLETだ")
		return p.parseLetStatement()
	case token.RETURN:
		log.Println("現在のTOKENはRETURNだ")
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// letステートメントの解析用。
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 今の場所がletの場合にここに来るので、
	// おもむろに現在のトークンをTokenとしてもつLetStatementのデータを生成
	stmt := &ast.LetStatement{Token: p.curToken}
	// letの場合、次に来るのは識別子（変数名）のはずだ
	log.Println("次のトークンは識別子であるはずだ。なぜなら変数宣言のはずだからだ。")
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	// 識別子であることが確認できたので、Identifier型のデータを作る
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	log.Println(stmt.Name)
	// 初期値が入ることを期待しているので、次には代入が来ることを期待
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	log.Println("次は代入だった")
	log.Println("今は", p.curToken)
	for !p.curTokenIs(token.SEMICOLON) {
		log.Println("今セミコロンではないので次のトークンに進む")
		p.nextToken()
		log.Println("今は", p.curToken)
	}
	log.Println("セミコロンだったのでlet文の解析はおしまい")
	return stmt
}

// returnステートメントの解析用
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	log.Println("次のトークンは", p.curToken)
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
		log.Println("次のトークンは", p.curToken)
	}
	log.Println("セミコロンだったのでreturn文の解析はおしまい")
	return stmt
}

// 現在見ているトークンのタイプが与えられたものかどうかを調べる
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// 次のトークンのタイプが与えられたものかどうかを調べる
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// 期待するトークンかチェック。正しい場合は次のトークンに移動
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
