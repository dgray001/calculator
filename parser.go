package main

// Constructs and returns the AST from the array of tokens
func parse(tokens []Token) (AstNode, error) {
	var root_node = newAstNode()
	for _, token := range tokens {
		var error = root_node.addToken(token)
		if error != nil {
			return root_node, error
		}
	}
	var end_error = root_node.endTokens()
	return root_node, end_error
}
