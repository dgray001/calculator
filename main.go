package main

import "fmt"

func main() {
	for {
		var input string
		fmt.Print(" > ")
		fmt.Scanln(&input)
		var tokens, tokenize_error = tokenize(input)
		if tokenize_error != nil {
			fmt.Println(tokenize_error)
			continue
		}
		//fmt.Println("Tokens: ", tokens)
		var ast, parse_error = parse(tokens)
		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}
		//fmt.Println("\nNode: ", ast.toString(false))
		var result, evaluate_error = ast.evaluate()
		if evaluate_error != nil {
			fmt.Println(evaluate_error)
			return
		}
		//fmt.Println("\nResult: ", result.toString(false))
		fmt.Println(result.toResultString())
	}
}
