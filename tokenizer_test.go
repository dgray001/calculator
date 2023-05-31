package main

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input         string
		expected      []Token
		error_message string
	}{
		{"unknown", []Token{}, "Unrecogized character: u"},
		{"1 ! 3", []Token{ONE}, "Unrecogized character: !"},
		{"in+", []Token{}, "Unrecogized character: +"},
		{"decin", []Token{DECREMENT}, "Unfinished tokens at end of input"},
		{"0", []Token{ZERO}, ""},
		{"0 \t 5 \n 9", []Token{ZERO, FIVE, NINE}, ""},
		{"+2", []Token{PLUS, TWO}, ""},
		{"-4", []Token{MINUS, FOUR}, ""},
		{"6*", []Token{SIX, MULTIPLY}, ""},
		{") (", []Token{CLOSE_PAREN, OPEN_PAREN}, ""},
		{"inc5-1", []Token{INCREMENT, FIVE, MINUS, ONE}, ""},
		{"( dec( ) )", []Token{OPEN_PAREN, DECREMENT, OPEN_PAREN, CLOSE_PAREN, CLOSE_PAREN}, ""},
	}
	for _, tc := range testCases {
		got, error := tokenize(tc.input)
		if !arrayEquals(got, tc.expected) {
			t.Errorf("For test case (%s, %d, %s), tokenize returned %d", tc.input, tc.expected, tc.error_message, got)
		}
		var error_message = ""
		if error != nil {
			error_message = error.Error()
		}
		if error_message != tc.error_message {
			t.Errorf("For test case (%s, %d, %s), tokenize returned error %s", tc.input, tc.expected, tc.error_message, error_message)
		}
	}
}
