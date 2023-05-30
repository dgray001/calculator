package main

import (
	"errors"
	"unicode"
)

// Returns the parsed integer
func tokenize(input string) (Integer, error) {
	// initialize state
	var return_int = newInteger()

	for _, rune := range input {
		if unicode.IsSpace(rune) {
			continue
		}
		var found_token *Token
		for token := Token(0); token < tokenLimit; token++ {
			if token.toString() == string(rune) {
				found_token = &token
				return_int = return_int.addDigit(token.toInt())
				break
			}
		}
		if found_token == nil {
			return return_int, errors.New("Unrecogized character: " + string(rune))
		}
	}

	return return_int, nil
}
