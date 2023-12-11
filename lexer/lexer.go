package lexer

import (
	token "DakInterpreter/token"
)

type Lexer struct {
	input        string
	position     int  // points to the position of 'ch'
	readPosition int  // comes after the curr position
	ch           byte // current char under examination
}

// New initializes a lexer object with an input at the first ch of that input.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// If we reached the end of our input.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (lex *Lexer) NextToken() token.Token {
	var currentToken token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case '=':
		// If also the next char is = this means we have a == case.
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			currentToken = token.Token{Type: token.EQ, Literal: string(ch) + string(lex.ch)}
		} else {
			currentToken = newToken(token.ASSIGN, lex.ch)
		}
	case '+':
		currentToken = newToken(token.PLUS, lex.ch)
	case '-':
		currentToken = newToken(token.MINUS, lex.ch)
	case '!':
		if lex.peekChar() == '=' {
			ch := lex.ch
			lex.readChar()
			currentToken = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(lex.ch)}
		} else {
			currentToken = newToken(token.BANG, lex.ch)
		}
	case '/':
		currentToken = newToken(token.SLASH, lex.ch)
	case '*':
		currentToken = newToken(token.ASTERISK, lex.ch)
	case '<':
		currentToken = newToken(token.LT, lex.ch)
	case '>':
		currentToken = newToken(token.GT, lex.ch)
	case ';':
		currentToken = newToken(token.SEMICOLON, lex.ch)
	case ',':
		currentToken = newToken(token.COMMA, lex.ch)
	case '{':
		currentToken = newToken(token.LBRACE, lex.ch)
	case '}':
		currentToken = newToken(token.RBRACE, lex.ch)
	case '(':
		currentToken = newToken(token.LPAREN, lex.ch)
	case ')':
		currentToken = newToken(token.RPAREN, lex.ch)
	case 0:
		currentToken.Literal = ""
		currentToken.Type = token.EOF
	default:
		if isLetter(lex.ch) {
			currentToken.Literal = lex.readIdentifier()
			currentToken.Type = token.LookupIdent(currentToken.Literal)
			return currentToken
		} else if isDigit(lex.ch) {
			currentToken.Literal = lex.readNumber()
			currentToken.Type = token.INT
			return currentToken
		} else {
			currentToken = newToken(token.ILLEGAL, lex.ch)
		}
	}
	lex.readChar()
	return currentToken

}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch < +'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
