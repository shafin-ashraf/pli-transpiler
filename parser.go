package main

import "strings"

type Program struct {
	Procedures []Procedure
}

type Procedure struct {
	Name       string
	Statements []Node
}

type CallStatement struct {
	ProcedureName string
}

type Declaration struct {
	Name    string
	Type    string
	Initial string
}

type Assignment struct {
	Left  string
	Right string
}

type DoLoop struct {
	Variable string
	Start    string
	End      string
	Step     string
	Body     []Node
}

type IfStatement struct {
	Condition     string
	ThenBody      []Node
	ElseBody      []Node
	HasElseClause bool
}

type Parser struct {
	tokens []Token
	pos    int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens: tokens,
		pos:    0,
	}
}

func (p *Parser) current() Token {
	if p.pos >= len(p.tokens) {
		return Token{Type: "EOF", Value: ""}
	}
	return p.tokens[p.pos]
}

func (p *Parser) peek() Token {
	if p.pos+1 >= len(p.tokens) {
		return Token{Type: "EOF", Value: ""}
	}
	return p.tokens[p.pos+1]
}

func (p *Parser) advance() {
	if p.pos < len(p.tokens) {
		p.pos++
	}
}

func (p *Parser) Parse() Program {
	program := Program{
		Procedures: make([]Procedure, 0),
	}

	mainProc := Procedure{
		Name:       "MAIN",
		Statements: make([]Node, 0),
	}
	hasMainProc := false

	insideProcedure := false
	var currentProc Procedure

	for p.current().Type != "EOF" {
		if p.current().Type == "IDENTIFIER" && p.peek().Type == "COLON" {
			if insideProcedure {
				program.Procedures = append(program.Procedures, currentProc)
			}
			currentProc = p.parseProcedure()
			if currentProc.Name == "MAIN" {
				hasMainProc = true
			}
			insideProcedure = true
		} else if p.current().Type == "PROCEDURE" {
			if insideProcedure {
				program.Procedures = append(program.Procedures, currentProc)
			}
			currentProc = p.parseProcedure()
			if currentProc.Name == "MAIN" {
				hasMainProc = true
			}
			insideProcedure = true
		} else {
			stmt := p.parseStatement()
			if stmt != nil {
				if !insideProcedure {
					mainProc.Statements = append(mainProc.Statements, stmt)
				} else {
					currentProc.Statements = append(currentProc.Statements, stmt)
				}
			}
		}
	}

	if insideProcedure {
		program.Procedures = append(program.Procedures, currentProc)
	}

	if len(mainProc.Statements) > 0 && !hasMainProc {
		program.Procedures = append([]Procedure{mainProc}, program.Procedures...)
	}

	return program
}

func (p *Parser) parseProcedure() Procedure {
	proc := Procedure{
		Statements: make([]Node, 0),
	}

	// optional label
	if p.current().Type == "IDENTIFIER" && p.peek().Type == "COLON" {
		proc.Name = p.current().Value
		// skipping indetifier
		p.advance()
		// skipping colon
		p.advance()
		if p.current().Type == "PROCEDURE" {
			// skipping PROCEDURE
			p.advance()
		}
	} else if p.current().Type == "PROCEDURE" {
		// skipping PROCEDURE
		p.advance()
		if p.current().Type == "IDENTIFIER" {
			proc.Name = p.current().Value
			p.advance()
		}
	}

	// skipping semicolon after procedure declaration
	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	// parsing procedure body
	for p.current().Type != "EOF" && p.current().Type != "END" {
		stmt := p.parseStatement()
		if stmt != nil {
			proc.Statements = append(proc.Statements, stmt)
		}
	}

	// skipping END and  semicolon
	if p.current().Type == "END" {
		p.advance()
		if p.current().Type == "SEMICOLON" {
			p.advance()
		}
	}

	return proc
}

func (p *Parser) parseStatement() Node {
	switch p.current().Type {
	case "DECLARE":
		return p.parseDeclaration()
	case "IDENTIFIER":
		return p.parseAssignment()
	case "DO":
		return p.parseDoLoop()
	case "IF":
		return p.parseIfStatement()
	case "CALL":
		return p.parseCallStatement()
	default:
		p.advance()
		return nil
	}
}

func (p *Parser) parseCallStatement() Node {
	// skipping CALL keyword
	p.advance()

	call := CallStatement{
		ProcedureName: p.current().Value,
	}

	// skipping procedure name
	p.advance()

	// skipping semicolon if present
	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	return call
}

func (p *Parser) parseDeclaration() Node {
	// skipping DECLARE
	p.advance()

	decl := Declaration{
		Type: "AUTO",
	}

	if p.current().Type == "IDENTIFIER" {
		decl.Name = p.current().Value
		p.advance()

		if p.current().Type == "FIXED" || p.current().Type == "FLOAT" || p.current().Type == "CHARACTER" {
			decl.Type = p.current().Value
			p.advance()
		}

		if p.current().Type == "EQUALS" {
			p.advance()
			if p.current().Type == "NUMBER" || p.current().Type == "IDENTIFIER" {
				decl.Initial = p.current().Value
				p.advance()
			}
		}
	}

	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	return decl
}

func (p *Parser) parseIfStatement() Node {
	ifStmt := IfStatement{
		ThenBody: make([]Node, 0),
		ElseBody: make([]Node, 0),
	}

	// skipping IF
	p.advance()

	// parsing condition insifde if  until THEN
	var condition strings.Builder
	for p.current().Type != "THEN" && p.current().Type != "EOF" {
		if p.current().Type == "OPERATOR" {
			if condition.Len() > 0 && !strings.HasSuffix(condition.String(), " ") {
				condition.WriteString(" ")
			}
			condition.WriteString(p.current().Value)
			if p.peek().Type != "THEN" {
				condition.WriteString(" ")
			}
		} else {
			condition.WriteString(p.current().Value)
			if p.peek().Type != "OPERATOR" && p.peek().Type != "THEN" {
				condition.WriteString(" ")
			}
		}
		p.advance()
	}
	ifStmt.Condition = strings.TrimSpace(condition.String())

	// skipping THEN
	if p.current().Type == "THEN" {
		p.advance()
	}

	// parsing THEN body until ELSE or END
	for p.current().Type != "EOF" && p.current().Type != "END" && p.current().Type != "ELSE" {
		stmt := p.parseStatement()
		if stmt != nil {
			ifStmt.ThenBody = append(ifStmt.ThenBody, stmt)
		}
	}

	//checking for ELSE
	if p.current().Type == "ELSE" {
		ifStmt.HasElseClause = true
		p.advance()

		// parsing ELSE body until END
		for p.current().Type != "EOF" && p.current().Type != "END" {
			stmt := p.parseStatement()
			if stmt != nil {
				ifStmt.ElseBody = append(ifStmt.ElseBody, stmt)
			}
		}
	}

	// skipping END
	if p.current().Type == "END" {
		p.advance()
	}

	// skipping semicolon
	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	return ifStmt
}

func (p *Parser) parseDoLoop() Node {
	doLoop := DoLoop{
		Body: make([]Node, 0),
	}

	// skipping DO
	p.advance()

	if p.current().Type == "IDENTIFIER" {
		doLoop.Variable = p.current().Value
		p.advance()
	}

	// parsing start, end, and step
	if p.current().Type == "EQUALS" {
		p.advance()
		doLoop.Start = p.parseExpression()

		if p.current().Type == "TO" {
			p.advance()
			doLoop.End = p.parseExpression()
		}
		if p.current().Type == "BY" {
			p.advance()
			doLoop.Step = p.parseExpression()
		}
	}

	for p.current().Type != "END" && p.current().Type != "EOF" {
		doLoop.Body = append(doLoop.Body, p.parseStatement())
	}

	// skipping "END" and semicolon
	if p.current().Type == "END" {
		p.advance()
	}
	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	return doLoop
}

func (p *Parser) parseAssignment() Node {
	assign := Assignment{
		Left: p.current().Value,
	}
	p.advance()

	if p.current().Type == "EQUALS" {
		p.advance()
		assign.Right = p.parseExpression()
	}

	// skipping semicolon
	if p.current().Type == "SEMICOLON" {
		p.advance()
	}

	return assign
}

func (p *Parser) parseExpression() string {
	var expr strings.Builder

	// captuing everything until semicolon
	for p.current().Type != "SEMICOLON" && p.current().Type != "EOF" && p.current().Type != "TO" && p.current().Type != "BY" {
		expr.WriteString(p.current().Value)
		if p.peek().Type != "SEMICOLON" {
			expr.WriteString(" ")
		}
		p.advance()
	}

	return strings.TrimSpace(expr.String())
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}

func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
