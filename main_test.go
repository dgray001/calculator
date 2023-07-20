package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		// Errors
		// Addition
		{input: "2", expected: "2"},
		{input: "-0", expected: "0"},
		{input: "2+2", expected: "4"},
		{input: "  2 -3 ", expected: "-1"},
		{input: "1 + 2 + 3 + 4 - 5", expected: "5"},
		{input: "3 - (5+2)", expected: "-4"},
		{input: "3 - (5 + (2) - (65+9-71)) + 1", expected: "0"},
		// Multiplication
	}
	for _, tc := range testCases {
		got := calculate(tc.input, false)
		if got != tc.expected {
			t.Errorf("Testing calculator input %s, expected %s but got %s", tc.input, tc.expected, got)
		}
	}
}
