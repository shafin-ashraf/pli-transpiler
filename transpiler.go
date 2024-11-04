package main

import (
	"fmt"
	"strings"
)

//type Node interface {
//	ToJS() string
//}

type Node interface {
	ToCS() string
}

//func (p Program) ToJS() string {
//	var js strings.Builder
//	js.WriteString("(function() {\n")
//	for _, proc := range p.Procedures {
//		js.WriteString(proc.ToJS())
//		js.WriteString("\n")
//	}
//	js.WriteString("})();")
//	return js.String()
//}
//
//func (p Procedure) ToJS() string {
//	var js strings.Builder
//	js.WriteString(fmt.Sprintf("function %s() {\n", p.Name))
//	for _, stmt := range p.Statements {
//		js.WriteString("  " + stmt.ToJS() + "\n")
//	}
//	js.WriteString("}")
//	return js.String()
//}
//
//func (d Declaration) ToJS() string {
//	if d.Initial != "" {
//		return fmt.Sprintf("let %s = %s;", d.Name, d.Initial)
//	}
//	return fmt.Sprintf("let %s;", d.Name)
//}
//
//func (a Assignment) ToJS() string {
//	return fmt.Sprintf("%s = %s;", a.Left, a.Right)
//}

func (p Program) ToCS() string {
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
			cs.WriteString(fmt.Sprintf("            var program = new Program();\n"))
			cs.WriteString(fmt.Sprintf("            program.%s();\n", p.Procedures[0].Name))
		}
		cs.WriteString("        }\n\n")
	}

	for _, proc := range p.Procedures {
		cs.WriteString(proc.ToCS())
		cs.WriteString("\n")
	}

	cs.WriteString("    }\n")
	cs.WriteString("}\n")
	return cs.String()
}

func (p Procedure) ToCS() string {
	var cs strings.Builder

	if p.Name == "MAIN" {
		cs.WriteString(fmt.Sprintf("        public static void Main(string[] args)\n"))
	} else {
		cs.WriteString(fmt.Sprintf("        public void %s()\n", p.Name))
	}

	cs.WriteString("        {\n")
	for _, stmt := range p.Statements {
		cs.WriteString("            " + stmt.ToCS() + "\n")
	}
	cs.WriteString("        }")
	return cs.String()
}

func (d Declaration) ToCS() string {
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

	if d.Initial != "" {
		if d.Type == "CHARACTER" && !strings.HasPrefix(d.Initial, "\"") {
			d.Initial = fmt.Sprintf("\"%s\"", d.Initial)
		}
		return fmt.Sprintf("%s %s = %s;", csType, d.Name, d.Initial)
	}
	return fmt.Sprintf("%s %s;", csType, d.Name)
}

func (a Assignment) ToCS() string {
	return fmt.Sprintf("%s = %s;", a.Left, a.Right)
}
