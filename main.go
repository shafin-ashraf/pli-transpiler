package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: transpiler <input-file>")
		os.Exit(1)
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
	}(inputFile)

	var input strings.Builder
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		input.WriteString(scanner.Text() + "\n")
	}

	lexer := NewLexer(input.String())
	tokens := lexer.Tokenize()

	parser := NewParser(tokens)
	program := parser.Parse()

	csharp := program.ToCS(0)

	fmt.Println(csharp)
}
