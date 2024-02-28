package main

import (
	"errors"
	"strings"
	"unicode"
)

// Returns an array of tokens fron the input string
func tokenize(input string) ([]Token, error) {
	whitespace_filter := func(r rune) bool {
		return !unicode.IsSpace(r)
	}
	input_runes := arrayFilter([]rune(input), whitespace_filter)
	tokens := []Token{}

	i := 0
	for i < len(input_runes) {
		partial := string(input_runes[i:])
		found_token := false
		for token := Token(0); token < tokenLimit; token++ {
			if strings.HasPrefix(partial, token.toString()) {
				tokens = append(tokens, token)
				i += len(token.toRunes())
				found_token = true
				break
			}
		}
		if found_token {
			continue
		}
		return tokens, errors.New("unrecognized character sequence starting at: " + string(input_runes[i:]))
	}

	return tokens, nil
}
