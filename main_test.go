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
		// Functions
		{input: "i n c 1", expected: "2"},
		{input: "2 - inc(23)", expected: "-22"},
		{input: "inc - 2", expected: "-1"},
		{input: "dec-1", expected: "-2"},
		{input: "dec(5)", expected: "4"},
		{input: "abs-10", expected: "10"},
		{input: "abs(-10)", expected: "10"},
		{input: "abs(2 - 1)", expected: "1"},
		{input: "4 - abs-1 + 3", expected: "6"},
		{input: "4 - abs(-1 + 3)", expected: "2"},
		{input: "4 - abs(-1) + 3", expected: "6"},
	}
	for _, tc := range testCases {
		got := calculate(tc.input, false)
		if got != tc.expected {
			t.Errorf("Testing calculator input %s, expected %s but got %s", tc.input, tc.expected, got)
		}
	}
}
