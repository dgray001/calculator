package main

import (
	"errors"
	"unicode"
)

// Returns the parsed integer
func tokenize(input string) ([]Token, error) {
	// initialize state
	var tokens = []Token{}

	for _, rune := range input {
		if unicode.IsSpace(rune) {
			continue
		}
		var found_token *Token
		for token := Token(0); token < tokenLimit; token++ {
			if token.toRune() == rune {
				found_token = &token
				tokens = append(tokens, token)
				break
			}
		}
		if found_token == nil {
			return tokens, errors.New("Unrecogized character: " + string(rune))
		}
	}

	return tokens, nil
}
