package main

import (
	"fmt"
	"strings"
)

type Node interface {
	ToJS() string
}

func (p Program) ToJS() string {
	var js strings.Builder
	js.WriteString("(function() {\n")
	for _, proc := range p.Procedures {
		js.WriteString(proc.ToJS())
		js.WriteString("\n")
	}
	js.WriteString("})();")
	return js.String()
}

func (p Procedure) ToJS() string {
	var js strings.Builder
	js.WriteString(fmt.Sprintf("function %s() {\n", p.Name))
	for _, stmt := range p.Statements {
		js.WriteString("  " + stmt.ToJS() + "\n")
	}
	js.WriteString("}")
	return js.String()
}

func (d Declaration) ToJS() string {
	if d.Initial != "" {
		return fmt.Sprintf("let %s = %s;", d.Name, d.Initial)
	}
	return fmt.Sprintf("let %s;", d.Name)
}

func (a Assignment) ToJS() string {
	return fmt.Sprintf("%s = %s;", a.Left, a.Right)
}
