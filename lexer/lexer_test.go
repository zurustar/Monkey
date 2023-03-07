package lexer

import (
	"log"
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tkn := l.NextToken()
		if tkn.Type != tt.expectedType {
			t.Fatalf("tests[%d] expected=%q, got=%q", i, tt.expectedType, tkn.Type)
		}
		if tkn.Literal != tt.expectedLiteral {
			t.Fatalf("test #%d, expected=%q, got=%q", i, tt.expectedLiteral, tkn.Literal)
		}
	}
}

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};
		let result = add(five, ten);
	`

	log.Println("tokenize:", input)
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tkn := l.NextToken()
		if tkn.Type != tt.expectedType {
			t.Fatalf("tests[%d] expected=%q, got=%q", i, tt.expectedType, tkn.Type)
		}
		if tkn.Literal != tt.expectedLiteral {
			t.Fatalf("test #%d, expected=%q, got=%q", i, tt.expectedLiteral, tkn.Literal)
		}
	}
}