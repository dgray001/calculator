package main

import (
	"errors"
	"unicode"
)

// Returns the parsed integer
func tokenize(input string) ([]Token, error) {
	// initialize state
	var tokens = []Token{}
	var partial_tokens = make(map[Token]int)

	for _, char := range input {
		if unicode.IsSpace(char) {
			continue
		}
		var found_token = false
		if len(partial_tokens) > 0 {
			for token := range partial_tokens {
				var token_found, finished_token = checkToken(char, token, partial_tokens, &tokens)
				if token_found {
					found_token = true
				}
				if finished_token {
					break
				}
			}
		} else {
			for token := Token(0); token < tokenLimit; token++ {
				var token_found, finished_token = checkToken(char, token, partial_tokens, &tokens)
				if token_found {
					found_token = true
				}
				if finished_token {
					break
				}
			}
		}
		if !found_token {
			return tokens, errors.New("Unrecogized character: " + string(char))
		}
	}
	if len(partial_tokens) > 0 {
		return tokens, errors.New("Unfinished tokens at end of input")
	}

	return tokens, nil
}

// checks input token and returns whether it was added to the token list
func checkToken(char rune, token Token, partial_tokens map[Token]int, tokens *[]Token) (bool, bool) {
	var runes = token.toRunes()
	var current_rune = runes[partial_tokens[token]]
	if current_rune == char {
		partial_tokens[token]++
		if partial_tokens[token] >= len(runes) {
			for k := range partial_tokens {
				delete(partial_tokens, k)
			}
			*tokens = append(*tokens, token)
			return true, true
		}
		return true, false
	}
	return false, false
}
