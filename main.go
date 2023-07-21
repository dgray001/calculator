package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// config
	var debug = false
	// main loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" >  ")
		scanner.Scan()
		input := scanner.Text()
		fmt.Println(calculate(input, debug))
	}
}

func calculate(input string, debug bool) string {
	var tokens, tokenize_error = tokenize(input)
	if tokenize_error != nil {
		return tokenize_error.Error()
	}
	if debug {
		fmt.Println("Tokens: ", tokens)
	}
	var ast, parse_error = parse(tokens)
	if parse_error != nil {
		return parse_error.Error()
	}
	if debug {
		fmt.Println("\nNode: ", ast.toDebugString("  "))
	}
	var result, evaluate_error = ast.evaluate()
	if evaluate_error != nil {
		return evaluate_error.Error()
	}
	if debug {
		fmt.Println("\nResult: ", result.toString(false))
	}
	return result.toResultString()
}
