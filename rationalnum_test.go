package main

import (
	"testing"
)

func TestRationalNumberEquals(t *testing.T) {
	type TestCase struct {
		left     RationalNumber
		right    RationalNumber
		expected bool
	}
	var testCases = []TestCase{
		{left: constructRationalNumber(12, 5), right: constructRationalNumber(24, 10), expected: false},
		{left: constructRationalNumber(12, 5), right: constructRationalNumber(12, 5), expected: true},
		{left: constructRationalNumber(12, 5), right: constructRationalNumber(-12, 5), expected: false},
		{left: constructRationalNumber(12, 5), right: constructRationalNumber(12, -5), expected: false},
		{left: constructRationalNumber(3, 1), right: constructRationalNumber(3, 2), expected: false},
		{left: constructRationalNumber(3, 1), right: constructRationalNumber(2, 1), expected: false},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(), tc.right.toString(), got)
		}
	}
}

func TestInvert(t *testing.T) {
	type TestCase struct {
		initial  RationalNumber
		expected RationalNumber
	}
	var testCases = []TestCase{
		{initial: constructRationalNumber(12, 5), expected: constructRationalNumber(-12, 5)},
		{initial: constructRationalNumber(-12, 5), expected: constructRationalNumber(12, 5)},
		{initial: constructRationalNumber(12, -5), expected: constructRationalNumber(12, 5)},
		{initial: constructRationalNumber(-12, -5), expected: constructRationalNumber(-12, 5)},
	}
	for _, tc := range testCases {
		got := tc.initial.invert()
		if !got.equals(tc.expected) {
			t.Errorf("For test case inverting %s, got %s", tc.initial.toString(), got.toString())
		}
	}
}

func TestAbs(t *testing.T) {
	type TestCase struct {
		initial  RationalNumber
		expected RationalNumber
	}
	var testCases = []TestCase{
		{initial: constructRationalNumber(12, 5), expected: constructRationalNumber(12, 5)},
		{initial: constructRationalNumber(-12, 5), expected: constructRationalNumber(12, 5)},
		{initial: constructRationalNumber(12, -5), expected: constructRationalNumber(12, 5)},
		{initial: constructRationalNumber(-12, -5), expected: constructRationalNumber(12, 5)},
	}
	for _, tc := range testCases {
		got := tc.initial.abs()
		if !got.equals(tc.expected) {
			t.Errorf("For test case taking abs of %s, got %s", tc.initial.toString(), got.toString())
		}
	}
}
