package main

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	testCases := []struct {
		root_node       AstNode
		expected_error  string
		expected_result Value
	}{
		{
			AstNode{},
			"Can't evaluate node without values",
			Value{},
		},
		{
			AstNode{values: []Value{intValue(newInteger().addDigit(TWO.toInt(), false).construct())}, operators: []Token{PLUS}},
			"",
			Value{value_type: INTEGER, integer: &Integer{digits: []uint8{2}, constructed: true, int_sign: true}},
		},
		{
			AstNode{values: []Value{intValue(newInteger().addDigit(TWO.toInt(), false).construct())}, operators: []Token{MINUS}},
			"",
			Value{value_type: INTEGER, integer: &Integer{digits: []uint8{2}, constructed: true, int_sign: false}},
		},
	}
	for _, tc := range testCases {
		var value, err = tc.root_node.evaluate()
		var got_err = ""
		if err != nil {
			got_err = err.Error()
		}
		if got_err != tc.expected_error {
			t.Errorf("For case evaluating %s, got error %s but expected %s",
				tc.root_node.toDebugString("  "), got_err, tc.expected_error)
		}
		if !value.equals(tc.expected_result) {
			t.Errorf("For case evaluating %s, got %s but expected %s",
				tc.root_node.toDebugString("  "), value.toString(false), tc.expected_result.toString(false))
		}
	}
}
