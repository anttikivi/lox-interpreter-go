package main

import (
	"fmt"
	"math"
	"os"
)

var (
	hasError        = false
	hasRuntimeError = false
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	source := string(fileContents)
	scanner := Scanner{source: source, tokens: []Token{}, start: 0, current: 0, line: 1}
	tokens := scanner.scanTokens()

	print(tokens)
	if hasError {
		os.Exit(65)
	}
}

func print(tokens []Token) {
	for _, token := range tokens {
		if token.Literal == nil {
			fmt.Printf("%s %s %s\n", tokenNames[token.Type], token.Lexeme, "null")
		} else {
			if token.Type == NUMBER {
				value := token.Literal.(float64)
				result := math.Floor(value)
				if result == value {
					fmt.Printf("%s %s %v.0\n", tokenNames[token.Type], token.Lexeme, token.Literal)
				} else {
					fmt.Printf("%s %s %v\n", tokenNames[token.Type], token.Lexeme, token.Literal)
				}
			} else {
				fmt.Printf("%s %s %v\n", tokenNames[token.Type], token.Lexeme, token.Literal)
			}
		}
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
	hasError = true
}
