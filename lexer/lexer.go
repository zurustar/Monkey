package lexer

import (
	"log"
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	log.Println("New(", input, ")")
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	log.Println("readChar()")
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	log.Println("(l *Lexer)NextToken()")
	var tkn token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		tkn = newToken(token.ASSIGN, l.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '{':
		tkn = newToken(token.LBRACE, l.ch)
	case '}':
		tkn = newToken(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			log.Println("***", tkn)
			return tkn
		} else {
			tkn = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	log.Println("xxx", tkn)
	return tkn
}

func isLetter(ch byte) bool {
	log.Println("isLetter(", ch, ")")
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	log.Println("(l *Lexer) readIdentifier()")
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	log.Println(position, l.position, l.input[position:l.position])
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	log.Println(tokenType, "¥t", ch)
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

