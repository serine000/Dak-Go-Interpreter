package lexer

import (
	token "DakInterpreter/token"
	"testing"
)

func TestNxtToken(t *testing.T) {
	input := `=+-(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, testToken := range tests {
		currentToken := lexer.NextToken()

		if currentToken.Type != testToken.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, testToken.expectedType, currentToken.Type)
		}

		if currentToken.Literal != testToken.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, testToken.expectedLiteral, currentToken.Literal)
		}
	}
}
