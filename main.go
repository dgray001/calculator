package main

import "fmt"

func main() {
	var tokens, tokenize_error = tokenize("1	5- 3")
	if tokenize_error != nil {
		fmt.Println(tokenize_error)
		return
	}
	fmt.Println("Tokens: ", tokens)
	var ast, parse_error = parse(tokens)
	if parse_error != nil {
		fmt.Println(parse_error)
		return
	}
	fmt.Println("Node: ", ast.toString())
}
