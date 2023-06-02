package main

import (
	"testing"
)

func TestAddToken(t *testing.T) {
	testCases := []struct {
		start       AstNode
		token       Token
		expectError bool
		expected    AstNode
	}{
		// non-error cases
		{
			AstNode{},
			TWO,
			false,
			AstNode{constructingLeftInt: true, left: newInteger().addDigit(TWO.toInt(), false)},
		},
		{
			AstNode{constructingLeftInt: true, left: newInteger().addDigit(ONE.toInt(), false)},
			THREE,
			false,
			AstNode{constructingLeftInt: true,
				left: newInteger().addDigit(ONE.toInt(), false).addDigit(THREE.toInt(), false)},
		},
		{
			AstNode{hasOperator: true},
			FOUR,
			false,
			AstNode{hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(FOUR.toInt(), false)},
		},
		{
			AstNode{hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(ONE.toInt(), false)},
			FIVE,
			false,
			AstNode{hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(ONE.toInt(), false).addDigit(FIVE.toInt(), false)},
		},
		{
			AstNode{hasLeft: true, hasOperator: true},
			SIX,
			false,
			AstNode{hasLeft: true, hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(SIX.toInt(), false)},
		},
		{
			AstNode{hasLeft: true, hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(ZERO.toInt(), false)},
			SEVEN,
			false,
			AstNode{hasLeft: true, hasOperator: true, constructingRightInt: true,
				right: newInteger().addDigit(ZERO.toInt(), false).addDigit(SEVEN.toInt(), false)},
		},
		// error cases
		{
			AstNode{hasLeft: true, constructingLeftInt: true},
			ONE,
			true,
			AstNode{hasLeft: true, constructingLeftInt: true},
		},
		{
			AstNode{hasLeft: true, hasRight: true},
			ONE,
			true,
			AstNode{hasLeft: true, hasRight: true},
		},
		{
			AstNode{hasLeft: true, hasOperator: false},
			ONE,
			true,
			AstNode{hasLeft: true, hasOperator: false},
		},
		{
			AstNode{hasOperator: true, constructingLeftInt: true},
			ONE,
			true,
			AstNode{hasOperator: true, constructingLeftInt: true},
		},
		{
			AstNode{hasOperator: true, hasRight: true},
			ONE,
			true,
			AstNode{hasOperator: true, hasRight: true},
		},
		{
			AstNode{constructingRightInt: true},
			ONE,
			true,
			AstNode{constructingRightInt: true},
		},
		// TODO: Change when rest is implemented
	}
	for _, tc := range testCases {
		got := tc.start.addToken(tc.token)
		if (got != nil) != tc.expectError {
			t.Errorf("For test case adding token %s to %s, got error %s but expected %t",
				tc.token.toString(), tc.start.toString(), got.Error(), tc.expectError)
		} else if !tc.start.equals(tc.expected) {
			t.Errorf("For test case adding token %s, got %s but expected %s",
				tc.token.toString(), tc.start.toString(), tc.expected.toString())
		}
	}
}
