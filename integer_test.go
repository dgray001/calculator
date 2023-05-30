package main

import (
	"testing"
)

func testEquals(t *testing.T) {
	testCases := []struct {
		left     Integer
		right    Integer
		expected bool
	}{
		{Integer{digits: []uint8{}}, Integer{digits: []uint8{}}, true},
		{Integer{digits: []uint8{1, 2}}, Integer{digits: []uint8{1, 2}}, true},
		{Integer{digits: []uint8{4, 6, 9}}, Integer{digits: []uint8{5, 6, 9}}, true},
		{Integer{digits: []uint8{}}, Integer{digits: []uint8{0}}, false},
		{Integer{digits: []uint8{4, 9}}, Integer{digits: []uint8{7}}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{0, 7}}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{2}}, false},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(), tc.right.toString(), got)
		}
	}
}

func TestAddDigit(t *testing.T) {
	testCases := []struct {
		starting Integer
		input    uint8
		expected Integer
	}{
		{Integer{digits: []uint8{}}, 0, Integer{digits: []uint8{0}}},
		{Integer{digits: []uint8{}}, 4, Integer{digits: []uint8{4}}},
		{Integer{digits: []uint8{4}}, 2, Integer{digits: []uint8{4, 2}}},
		{Integer{digits: []uint8{8, 0}}, 0, Integer{digits: []uint8{8, 0, 0}}},
	}
	for _, tc := range testCases {
		got := tc.starting.addDigit(tc.input)
		if !got.equals(tc.expected) {
			t.Errorf("For test case expecting %s after adding %d, got %s", tc.expected.toString(), tc.input, got.toString())
		}
	}
}

func TestAddDigitPanic(t *testing.T) {
	testCases := []struct {
		starting Integer
		input    uint8
	}{
		{Integer{digits: []uint8{}}, 11},
	}
	for _, tc := range testCases {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Test case adding %d to %s expecting panic did not panic", tc.input, tc.starting.toString())
			}
		}()
		tc.starting.addDigit(tc.input)
	}
}
