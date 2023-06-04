package main

import (
	"testing"
)

func TestAstNodeEquals(t *testing.T) {
	testCases := []struct {
		left     AstNode
		right    AstNode
		expected bool
	}{
		{AstNode{}, AstNode{}, true},
		{
			constructAstNode(INCREMENT, []Value{}, []Token{}, newInteger(), AstNode{}, true, true),
			constructAstNode(INCREMENT, []Value{}, []Token{}, newInteger(), AstNode{}, true, true),
			true,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			true,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(INCREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(DECREMENT, []Value{intValue(newInteger())}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE, THREE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{lastAddedValue: true}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false).construct(), AstNode{}, true, true),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false), AstNode{}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false), AstNode{}, false, true),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false), AstNode{}, true, true),
			false,
		},
		{
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false), AstNode{}, true, false),
			constructAstNode(DECREMENT, []Value{}, []Token{FIVE},
				newInteger().addDigit(THREE.toInt(), false), AstNode{}, true, true),
			false,
		},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(true), tc.right.toString(true), got)
		}
	}
}

func TestAddToken(t *testing.T) {
	var new_integer Integer
	var added_integer Integer
	testCases := []struct {
		start       AstNode
		token       Token
		expectError bool
		expected    AstNode
	}{
		// int tokens
		{
			AstNode{lastAddedValue: true},
			ONE,
			true,
			AstNode{lastAddedValue: true},
		},
		{
			AstNode{constructingNode: &AstNode{}},
			MINUS,
			false,
			AstNode{constructingNode: &AstNode{lastAddedOperator: true, operators: []Token{MINUS}}},
		},
		{
			AstNode{constructingInt: &new_integer},
			TWO,
			false,
			AstNode{constructingInt: &added_integer},
		},
		{
			AstNode{},
			THREE,
			false,
			AstNode{constructingInt: &added_integer},
		},
		{
			AstNode{lastAddedOperator: true},
			FOUR,
			false,
			AstNode{lastAddedOperator: true, constructingInt: &added_integer},
		},
		{
			AstNode{lastAddedOperator: true},
			MINUS,
			true,
			AstNode{lastAddedOperator: true},
		},
		{
			AstNode{},
			MINUS,
			false,
			AstNode{lastAddedOperator: true, operators: []Token{MINUS}},
		},
		{
			AstNode{lastAddedValue: true},
			MINUS,
			false,
			AstNode{lastAddedOperator: true, operators: []Token{MINUS}},
		},
		{
			AstNode{constructingInt: &added_integer},
			MINUS,
			false,
			AstNode{lastAddedOperator: true, values: []Value{intValue(newInteger().addDigit(FIVE.toInt(), false).construct())}, operators: []Token{MINUS}},
		},
		{
			AstNode{},
			CLOSE_PAREN,
			true,
			AstNode{},
		},
		{
			AstNode{constructingNode: &AstNode{}},
			PLUS,
			false,
			AstNode{constructingNode: &AstNode{operators: []Token{PLUS}, lastAddedOperator: true}},
		},
		{
			AstNode{constructingNode: &AstNode{}},
			CLOSE_PAREN,
			false,
			AstNode{values: []Value{{value_type: AST_NODE, ast_node: &AstNode{}}}, lastAddedValue: true},
		},
		{
			AstNode{},
			INCREMENT,
			true,
			AstNode{},
		},
	}
	for _, tc := range testCases {
		new_integer = newInteger()
		if tc.token.isInt() {
			added_integer = newInteger().addDigit(tc.token.toInt(), false)
		} else {
			added_integer = newInteger().addDigit(FIVE.toInt(), false)
		}
		got := tc.start.addToken(tc.token)
		if (got != nil) != tc.expectError {
			t.Errorf("For test case adding token %s to %s, got error %s but expected %t",
				tc.token.toString(), tc.start.toString(true), got.Error(), tc.expectError)
		} else if !tc.start.equals(tc.expected) {
			t.Errorf("For test case adding token %s, got %s but expected %s",
				tc.token.toString(), tc.start.toString(false), tc.expected.toString(false))
		}
	}
}
