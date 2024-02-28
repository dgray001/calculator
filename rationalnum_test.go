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

func TestInc(t *testing.T) {
	type TestCase struct {
		initial  RationalNumber
		expected RationalNumber
	}
	var testCases = []TestCase{
		{initial: constructRationalNumber(12, 5), expected: constructRationalNumber(17, 5)},
		{initial: constructRationalNumber(-12, 5), expected: constructRationalNumber(-7, 5)},
		{initial: constructRationalNumber(12, -5), expected: constructRationalNumber(-7, 5)},
		{initial: constructRationalNumber(-12, -5), expected: constructRationalNumber(17, 5)},
		{initial: constructRationalNumber(1, 3), expected: constructRationalNumber(4, 3)},
		{initial: constructRationalNumber(-1, 3), expected: constructRationalNumber(2, 3)},
	}
	for _, tc := range testCases {
		got := tc.initial.increment()
		if !got.equals(tc.expected) {
			t.Errorf("For test case taking inc of %s, got %s", tc.initial.toString(), got.toString())
		}
	}
}

func TestDec(t *testing.T) {
	type TestCase struct {
		initial  RationalNumber
		expected RationalNumber
	}
	var testCases = []TestCase{
		{initial: constructRationalNumber(12, 5), expected: constructRationalNumber(7, 5)},
		{initial: constructRationalNumber(-12, 5), expected: constructRationalNumber(-17, 5)},
		{initial: constructRationalNumber(12, -5), expected: constructRationalNumber(-17, 5)},
		{initial: constructRationalNumber(-12, -5), expected: constructRationalNumber(7, 5)},
		{initial: constructRationalNumber(1, 3), expected: constructRationalNumber(-2, 3)},
		{initial: constructRationalNumber(-1, 3), expected: constructRationalNumber(-4, 3)},
	}
	for _, tc := range testCases {
		got := tc.initial.decrement()
		if !got.equals(tc.expected) {
			t.Errorf("For test case taking dec of %s, got %s", tc.initial.toString(), got.toString())
		}
	}
}

func TestIsZero(t *testing.T) {
	type TestCase struct {
		initial  RationalNumber
		expected bool
	}
	var testCases = []TestCase{
		{initial: constructRationalNumber(0, 5), expected: true},
		{initial: constructRationalNumber(0, -1), expected: true},
		{initial: constructRationalNumber(1, 1), expected: false},
		{initial: constructRationalNumber(-1, 1), expected: false},
		{initial: constructRationalNumber(1, 9999), expected: false},
	}
	for _, tc := range testCases {
		got := tc.initial.isZero()
		if got != tc.expected {
			t.Errorf("For test case checking isZero of %s, got %t", tc.initial.toString(), got)
		}
	}
}

func TestCompare(t *testing.T) {
	type TestCase struct {
		i        RationalNumber
		j        RationalNumber
		expected CompareResult
	}
	var testCases = []TestCase{
		{i: constructRationalNumber(0, 5), j: constructRationalNumber(0, 5), expected: EQUAL_TO},
		{i: constructRationalNumber(12, 5), j: constructRationalNumber(-12, 5), expected: GREATER_THAN},
		{i: constructRationalNumber(12, -5), j: constructRationalNumber(12, 5), expected: LESSER_THAN},
		{i: constructRationalNumber(16, 4), j: constructRationalNumber(4, 1), expected: EQUAL_TO},
		{i: constructRationalNumber(99, 100), j: constructRationalNumber(98, 99), expected: GREATER_THAN},
		{i: constructRationalNumber(-0, 3), j: constructRationalNumber(-1, 6), expected: GREATER_THAN},
	}
	for _, tc := range testCases {
		got := tc.i.compare(tc.j)
		if got != tc.expected {
			t.Errorf("For test case comparing %s and %s, got %d", tc.i.toString(), tc.j.toString(), got)
		}
	}
}

func TestAddRational(t *testing.T) {
	type TestCase struct {
		i        RationalNumber
		j        RationalNumber
		expected string
	}
	var testCases = []TestCase{
		{i: constructRationalNumber(0, 5), j: constructRationalNumber(0, 5), expected: "0/25"},
		{i: constructRationalNumber(12, 5), j: constructRationalNumber(3, 5), expected: "75/25"},
	}
	for _, tc := range testCases {
		got := tc.i.add(tc.j)
		if got.toString() != tc.expected {
			t.Errorf("For test case adding %s and %s, got %s", tc.i.toString(), tc.j.toString(), got.toString())
		}
	}
}

func TestSubtractRational(t *testing.T) {
	type TestCase struct {
		i        RationalNumber
		j        RationalNumber
		expected string
	}
	var testCases = []TestCase{
		{i: constructRationalNumber(3, 5), j: constructRationalNumber(12, 5), expected: "-45/25"},
	}
	for _, tc := range testCases {
		got := tc.i.subtract(tc.j)
		if got.toString() != tc.expected {
			t.Errorf("For test case subtracting %s and %s, got %s", tc.i.toString(), tc.j.toString(), got.toString())
		}
	}
}

func TestMultiplyRational(t *testing.T) {
	type TestCase struct {
		i        RationalNumber
		j        RationalNumber
		expected string
	}
	var testCases = []TestCase{
		{i: constructRationalNumber(2, 1), j: constructRationalNumber(3, 1), expected: "6/1"},
		{i: constructRationalNumber(-2, 1), j: constructRationalNumber(3, 1), expected: "-6/1"},
		{i: constructRationalNumber(2, 1), j: constructRationalNumber(3, -1), expected: "-6/1"},
		{i: constructRationalNumber(2, -1), j: constructRationalNumber(-3, 1), expected: "6/1"},
		{i: constructRationalNumber(3, 5), j: constructRationalNumber(12, 5), expected: "36/25"},
		{i: constructRationalNumber(1, 3), j: constructRationalNumber(2, -5), expected: "-2/15"},
	}
	for _, tc := range testCases {
		got := tc.i.multiply(tc.j)
		if got.toString() != tc.expected {
			t.Errorf("For test case multiplying %s and %s, got %s", tc.i.toString(), tc.j.toString(), got.toString())
		}
	}
}

func TestDivideRational(t *testing.T) {
	type TestCase struct {
		i        RationalNumber
		j        RationalNumber
		expected string
	}
	var testCases = []TestCase{
		{i: constructRationalNumber(2, 1), j: constructRationalNumber(3, 1), expected: "2/3"},
		{i: constructRationalNumber(-2, 1), j: constructRationalNumber(3, 1), expected: "-2/3"},
		{i: constructRationalNumber(2, 1), j: constructRationalNumber(3, -1), expected: "-2/3"},
		{i: constructRationalNumber(2, -1), j: constructRationalNumber(-3, 1), expected: "2/3"},
		{i: constructRationalNumber(3, 5), j: constructRationalNumber(12, 5), expected: "15/60"},
		{i: constructRationalNumber(1, 3), j: constructRationalNumber(2, -5), expected: "-5/6"},
	}
	for _, tc := range testCases {
		got := tc.i.divide(tc.j)
		if got.toString() != tc.expected {
			t.Errorf("For test case dividing %s and %s, got %s", tc.i.toString(), tc.j.toString(), got.toString())
		}
	}
}
