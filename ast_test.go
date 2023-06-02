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
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			true,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger().addDigit(ONE.toInt(), false), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: PLUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger().addDigit(ONE.toInt(), false), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: false,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: false, hasLeft: true, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: false, hasOperator: true, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: false, hasRight: true},
			false,
		},
		{
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: true},
			AstNode{left: newInteger(), operator: MINUS, right: newInteger(), constructingLeftInt: true,
				constructingRightInt: true, hasLeft: true, hasOperator: true, hasRight: false},
			false,
		},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(), tc.right.toString(), got)
		}
	}
}

func TestAddToken(t *testing.T) {
	testCases := []struct {
		start       AstNode
		token       Token
		expectError bool
		expected    AstNode
	}{
		// non-error int cases
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
		// error int cases
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
			AstNode{hasRight: true},
			ONE,
			true,
			AstNode{hasRight: true},
		},
		{
			AstNode{constructingRightInt: true},
			ONE,
			true,
			AstNode{constructingRightInt: true},
		},
		// non error add operator cases
		{
			AstNode{},
			PLUS,
			false,
			AstNode{hasOperator: true, operator: PLUS},
		},
		{
			AstNode{hasLeft: true},
			PLUS,
			false,
			AstNode{hasLeft: true, hasOperator: true, operator: PLUS},
		},
		{
			AstNode{constructingLeftInt: true, left: newInteger().addDigit(TWO.toInt(), false)},
			PLUS,
			false,
			AstNode{hasLeft: true, left: newInteger().addDigit(TWO.toInt(), false).construct(), hasOperator: true, operator: PLUS},
		},
		// error add operator cases
		{
			AstNode{hasOperator: true},
			PLUS,
			true,
			AstNode{hasOperator: true},
		},
		{
			AstNode{hasRight: true},
			PLUS,
			true,
			AstNode{hasRight: true},
		},
		{
			AstNode{constructingRightInt: true},
			PLUS,
			true,
			AstNode{constructingRightInt: true},
		},
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
