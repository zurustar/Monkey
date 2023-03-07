package lexer

import (
	"log"
	"monkey/token"
)

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer)NextToken() token.Token {
	var tkn token.Token
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
			return tkn
		}
	}
	l.readChar()
	return tkn
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter (l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}


func newToken(tokenType token.TokenType, ch byte) token.Token {
	log.Println(tokenType, "¥t", ch)
	return token.Token{Type: tokenType, Literal: string(ch)}
}

