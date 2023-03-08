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

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// 次のトークンを取得する
func (l *Lexer) NextToken() token.Token {
	var tkn token.Token

	l.skipWhitespace() // 空白を飛ばす

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tkn = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tkn = newToken(token.PLUS, l.ch)
	case '-':
		tkn = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tkn = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tkn = newToken(token.BANG, l.ch)
		}
	case '/':
		tkn = newToken(token.SLASH, l.ch)
	case '*':
		tkn = newToken(token.ASTERISK, l.ch)
	case '<':
		tkn = newToken(token.LT, l.ch)
	case '>':
		tkn = newToken(token.GT, l.ch)
	case ';':
		tkn = newToken(token.SEMICOLON, l.ch)
	case '(':
		tkn = newToken(token.LPAREN, l.ch)
	case ')':
		tkn = newToken(token.RPAREN, l.ch)
	case ',':
		tkn = newToken(token.COMMA, l.ch)
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
			return tkn
		} else if isDigit(l.ch) {
			tkn.Type = token.INT
			tkn.Literal = l.readNumber()
			return tkn
		} else {
			tkn = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tkn
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}



func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	log.Println(position, l.position, l.input[position:l.position])
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	log.Println(position, l.position, l.input[position:l.position])
	return l.input[position:l.position]
}

// 1文字先読みする
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

