package main

import (
	"testing"
)

func TestRationalNumberEquals(t *testing.T) {
	var ints = []Integer{
		{digits: []uint8{5}},
		{digits: []uint8{1, 2}},
		{digits: []uint8{1, 2}, constructed: true},
	}
	var rationals = []RationalNumber{
		{},
		{rational_sign: true, numerator: ints[0], denominator: ints[1]},
		{rational_sign: false, numerator: ints[0], denominator: ints[1]},
		{rational_sign: true, numerator: ints[1], denominator: ints[2]},
		{rational_sign: true, numerator: ints[2], denominator: ints[0]},
	}
	type TestCase struct {
		left     RationalNumber
		right    RationalNumber
		expected bool
	}
	var testCases = []TestCase{}
	for i := range rationals {
		for j := range rationals {
			if i == j {
				testCases = append(testCases, TestCase{rationals[i], rationals[j], true})
			} else {
				testCases = append(testCases, TestCase{rationals[i], rationals[j], false})
			}
		}
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(), tc.right.toString(), got)
		}
	}
}
