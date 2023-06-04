package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		tokens         []Token
		expected_error string
		expected_ast   AstNode
	}{
		{
			[]Token{FIVE, TWO, PLUS, OPEN_BRACKET, NINE, CLOSE_BRACKET, MULTIPLY, ONE},
			"",
			AstNode{
				values: []Value{
					{value_type: INTEGER, integer: &Integer{digits: []uint8{2, 5}, constructed: true, int_sign: true}},
					{value_type: AST_NODE, ast_node: &AstNode{
						values:         []Value{{value_type: INTEGER, integer: &Integer{digits: []uint8{9}, constructed: true, int_sign: true}}},
						lastAddedValue: true,
					}},
					{value_type: INTEGER, integer: &Integer{digits: []uint8{1}, constructed: true, int_sign: true}},
				},
				operators:      []Token{PLUS, MULTIPLY},
				lastAddedValue: true,
			},
		},
	}
	for _, tc := range testCases {
		var tokens_strings = []string{}
		for _, token := range tc.tokens {
			tokens_strings = append(tokens_strings, token.toString())
		}
		var tokens_string = strings.Join(tokens_strings, " ,")
		var node, e = parse(tc.tokens)
		var error = ""
		if e != nil {
			error = e.Error()
		}
		if tc.expected_error != error {
			t.Error("For test case parsing &s, expected error &s but got &s", tokens_string, tc.expected_error, error)
		}
		if !node.equals(tc.expected_ast) {
			t.Errorf("For test case parsing %s, expected %s but got %s", tokens_string, tc.expected_ast.toDebugString("  "), node.toDebugString("  "))
		}
	}
}
