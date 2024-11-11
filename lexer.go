package main

import "strings"

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	input  string
	pos    int
	tokens []Token
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:  input,
		pos:    0,
		tokens: make([]Token, 0),
	}
}

func (l *Lexer) current() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	return l.input[l.pos]
}

func (l *Lexer) advance() {
	l.pos++
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && isWhitespace(l.current()) {
		l.advance()
	}
}

func (l *Lexer) Tokenize() []Token {
	for l.pos < len(l.input) {
		l.skipWhitespace()

		if l.pos >= len(l.input) {
			break
		}

		switch {
		case isLetter(l.current()):
			l.tokenizeIdentifier()
		case isDigit(l.current()):
			l.tokenizeNumber()
		case l.current() == ':':
			l.tokens = append(l.tokens, Token{Type: "COLON", Value: ":"})
			l.advance()
		case l.current() == ';':
			l.tokens = append(l.tokens, Token{Type: "SEMICOLON", Value: ";"})
			l.advance()
		case l.current() == '=':
			l.tokens = append(l.tokens, Token{Type: "EQUALS", Value: "="})
			l.advance()
		case l.current() == '+':
			l.tokens = append(l.tokens, Token{Type: "PLUS", Value: "+"})
			l.advance()
		case l.current() == '-':
			l.tokens = append(l.tokens, Token{Type: "MINUS", Value: "-"})
			l.advance()
		case l.current() == '*':
			l.tokens = append(l.tokens, Token{Type: "MULTIPLY", Value: "*"})
			l.advance()
		case l.current() == '/':
			l.tokens = append(l.tokens, Token{Type: "DIVIDE", Value: "/"})
			l.advance()
		case l.current() == '(':
			l.tokens = append(l.tokens, Token{Type: "LPAREN", Value: "("})
			l.advance()
		case l.current() == ')':
			l.tokens = append(l.tokens, Token{Type: "RPAREN", Value: ")"})
			l.advance()
		default:
		case l.current() == '>':
			if l.pos+1 < len(l.input) && l.input[l.pos+1] == '=' {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: ">="})
				l.advance()
				l.advance()
			} else {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: ">"})
				l.advance()
			}
		case l.current() == '<':
			if l.pos+1 < len(l.input) && l.input[l.pos+1] == '=' {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: "<="})
				l.advance()
				l.advance()
			} else {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: "<"})
				l.advance()
			}
		case l.current() == '¬':
			if l.pos+1 < len(l.input) && l.input[l.pos+1] == '=' {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: "!="})
				l.advance()
				l.advance()
			} else {
				l.tokens = append(l.tokens, Token{Type: "OPERATOR", Value: "!"})
				l.advance()
			}
			// skipping unknown
			l.advance()
		}
	}

	l.tokens = append(l.tokens, Token{Type: "EOF", Value: ""})
	return l.tokens
}

func (l *Lexer) tokenizeIdentifier() {
	var value strings.Builder

	for l.pos < len(l.input) && (isLetter(l.current()) || isDigit(l.current()) || l.current() == '_') {
		value.WriteByte(l.current())
		l.advance()
	}

	identifier := value.String()
	tokenType := "IDENTIFIER"

	// supported keywords
	switch strings.ToUpper(identifier) {
	case "PROCEDURE":
		tokenType = "PROCEDURE"
	case "DECLARE", "DCL":
		tokenType = "DECLARE"
	case "END":
		tokenType = "END"
	case "FIXED":
		tokenType = "FIXED"
	case "FLOAT":
		tokenType = "FLOAT"
	case "CHARACTER", "CHAR":
		tokenType = "CHARACTER"
	case "DO":
		tokenType = "DO"
	case "WHILE":
		tokenType = "WHILE"
	case "TO":
		tokenType = "TO"
	case "BY":
		tokenType = "BY"
	case "IF":
		tokenType = "IF"
	case "THEN":
		tokenType = "THEN"
	case "ELSE":
		tokenType = "ELSE"
	}

	l.tokens = append(l.tokens, Token{Type: tokenType, Value: identifier})
}

func (l *Lexer) tokenizeNumber() {
	var value strings.Builder

	for l.pos < len(l.input) && (isDigit(l.current()) || l.current() == '.') {
		value.WriteByte(l.current())
		l.advance()
	}

	l.tokens = append(l.tokens, Token{Type: "NUMBER", Value: value.String()})
}
