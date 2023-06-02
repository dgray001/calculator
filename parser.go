package main

// Returns the parsed integer
func parse(tokens []Token) (AstNode, error) {
	var current_node = newAstNode()
	for _, token := range tokens {
		var e = current_node.addToken(token)
		if e != nil {
			return current_node, e
		}
	}
	var e = current_node.endTokens()
	if e != nil {
		return current_node, e
	}
	return current_node, nil
}
