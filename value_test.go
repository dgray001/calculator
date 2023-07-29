package main

import (
	"testing"
)

func TestValueEquals(t *testing.T) {
	testCases := []struct {
		left     Value
		right    Value
		expected bool
	}{
		{Value{value_type: INTEGER, integer: &Integer{}}, Value{value_type: INTEGER, integer: &Integer{}}, true},
		{
			Value{value_type: INTEGER, integer: &Integer{digits: []uint8{1, 2}, constructed: true}},
			Value{value_type: INTEGER, integer: &Integer{digits: []uint8{1, 2}, constructed: true}},
			true,
		},
		{Value{value_type: AST_NODE, ast_node: &AstNode{}}, Value{value_type: AST_NODE, ast_node: &AstNode{}}, true},
		{
			Value{value_type: AST_NODE, ast_node: &AstNode{operators: []Token{NINE}, lastAddedValue: true}},
			Value{value_type: AST_NODE, ast_node: &AstNode{operators: []Token{NINE}, lastAddedValue: true}},
			true,
		},
		{Value{value_type: INTEGER, integer: &Integer{}}, Value{value_type: INTEGER, integer: &Integer{constructed: true}}, false},
		{Value{value_type: AST_NODE, ast_node: &AstNode{}}, Value{value_type: AST_NODE, ast_node: &AstNode{lastAddedValue: true}}, false},
		{Value{value_type: RATIONAL_NUMBER, rational: &RationalNumber{}}, Value{value_type: INTEGER, integer: &Integer{constructed: true}}, false},
		{Value{value_type: RATIONAL_NUMBER, rational: &RationalNumber{}}, Value{value_type: AST_NODE, ast_node: &AstNode{lastAddedValue: true}}, false},
		{Value{value_type: RATIONAL_NUMBER, rational: &RationalNumber{}}, Value{value_type: RATIONAL_NUMBER, rational: &RationalNumber{}}, true},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(true), tc.right.toString(true), got)
		}
	}
}
