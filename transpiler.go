package main

import (
	"fmt"
	"strings"
)

// The Node interface and various types implementing ToCS() are defined here.

type Node interface {
	ToCS(padding int) string
}

func (p Program) ToCS(padding int) string {
	var cs strings.Builder
	cs.WriteString("using System;\n\n")
	cs.WriteString("namespace PLIProgram\n{\n")
	cs.WriteString("    public class Program\n    {\n")

	hasMain := false
	for _, proc := range p.Procedures {
		if proc.Name == "MAIN" {
			hasMain = true
			break
		}
	}

	if !hasMain {
		cs.WriteString("        public static void Main(string[] args)\n")
		cs.WriteString("        {\n")
		cs.WriteString("            // Entry point\n")
		if len(p.Procedures) > 0 {
			cs.WriteString("            var program = new Program();\n")
			cs.WriteString(fmt.Sprintf("            program.%s();\n", p.Procedures[0].Name))
		}
		cs.WriteString("        }\n\n")
	}

	for _, proc := range p.Procedures {
		cs.WriteString(proc.ToCS(2))
		cs.WriteString("\n")
	}

	cs.WriteString("    }\n")
	cs.WriteString("}\n")
	return cs.String()
}

func (p Procedure) ToCS(padding int) string {
	var cs strings.Builder
	indent := strings.Repeat("    ", padding)

	if p.Name == "MAIN" {
		cs.WriteString(indent + "public static void Main(string[] args)\n")
	} else {
		cs.WriteString(indent + fmt.Sprintf("public void %s()\n", p.Name))
	}

	cs.WriteString(indent + "{\n")
	for _, stmt := range p.Statements {
		cs.WriteString(stmt.ToCS(padding+1) + "\n")
	}
	cs.WriteString(indent + "}")
	return cs.String()
}

func (d Declaration) ToCS(padding int) string {
	typeMap := map[string]string{
		"FIXED":     "int",
		"FLOAT":     "double",
		"CHARACTER": "string",
		"AUTO":      "var",
	}

	csType := typeMap[d.Type]
	if csType == "" {
		csType = "var"
	}

	indent := strings.Repeat("    ", padding)
	if d.Initial != "" {
		if d.Type == "CHARACTER" && !strings.HasPrefix(d.Initial, "\"") {
			d.Initial = fmt.Sprintf("\"%s\"", d.Initial)
		}
		return fmt.Sprintf("%s%s %s = %s;", indent, csType, d.Name, d.Initial)
	}
	return fmt.Sprintf("%s%s %s;", indent, csType, d.Name)
}

func (a Assignment) ToCS(padding int) string {
	indent := strings.Repeat("    ", padding)
	return fmt.Sprintf("%s%s = %s;", indent, a.Left, a.Right)
}

func (d DoLoop) ToCS(padding int) string {
	var cs strings.Builder
	indent := strings.Repeat("    ", padding)

	// default step is 1 if not defined
	step := "1"
	if d.Step != "" {
		step = d.Step
	}

	cs.WriteString(fmt.Sprintf("%sfor (int %s = %s; %s <= %s; %s += %s) {\n",
		indent, d.Variable, d.Start, d.Variable, d.End, d.Variable, step))

	for _, stmt := range d.Body {
		if stmt != nil {
			cs.WriteString(stmt.ToCS(padding+1) + "\n")
		}
	}

	cs.WriteString(indent + "}")
	return cs.String()
}
