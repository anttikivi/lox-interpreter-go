package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
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

	// TODO: Actually add the tokens somewhere.
	for _, c := range string(fileContents) {
		switch c {
		case LEFT_PAREN:
			fmt.Printf("LEFT_PAREN %c null\n", LEFT_PAREN)
		case RIGHT_PAREN:
			fmt.Printf("RIGHT_PAREN %c null\n", RIGHT_PAREN)
		}
	}
	fmt.Println("EOF  null")
}
