package main

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		input         string
		expected      Integer
		error_message string
	}{
		{"unknown", newInteger(), "Unrecogized character: u"},
		{"1 ! 3", newInteger().addDigit(1), "Unrecogized character: !"},
		{"0", newInteger().addDigit(0), ""},
		{"0 \t 5 \n 9", newInteger().addDigit(0).addDigit(5).addDigit(9), ""},
	}
	for _, tc := range testCases {
		got, error := tokenize(tc.input)
		if !got.equals(tc.expected) {
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
