package main

import (
	"testing"
)

func TestIntegerEquals(t *testing.T) {
	testCases := []struct {
		left     Integer
		right    Integer
		expected bool
	}{
		{Integer{digits: []uint8{}}, Integer{digits: []uint8{}}, true},
		{Integer{digits: []uint8{2, 1}}, Integer{digits: []uint8{2, 1}}, true},
		{Integer{digits: []uint8{4, 6, 9}}, Integer{digits: []uint8{4, 6, 9}}, true},
		{Integer{digits: []uint8{}}, Integer{digits: []uint8{0}}, false},
		{Integer{digits: []uint8{4, 9}}, Integer{digits: []uint8{7}}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{7, 0}}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{0, 7}}, false},
		{Integer{digits: []uint8{0, 7}}, Integer{digits: []uint8{0, 7}}, true},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{2}}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{7}, constructed: true}, false},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{7}, int_sign: true}, false},
	}
	for _, tc := range testCases {
		got := tc.left.equals(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s to %s, got %t", tc.left.toString(), tc.right.toString(), got)
		}
	}
}

func TestIntegerToString(t *testing.T) {
	testCases := []struct {
		i        Integer
		expected string
	}{
		{constructInt(0), "0"},
		{constructInt(11), "11"},
		{constructInt(-11), "-11"},
	}
	for _, tc := range testCases {
		got := tc.i.toString()
		if got != tc.expected {
			t.Errorf("When converting int to string got %s, expected %s", got, tc.expected)
		}
	}
}

func TestConstructInt(t *testing.T) {
	testCases := []struct {
		i        int
		expected string
	}{
		{0, "0"},
		{-0, "0"},
		{4, "4"},
		{-4, "-4"},
		{7019, "7019"},
		{-7019, "-7019"},
	}
	for _, tc := range testCases {
		var got = constructInt(tc.i)
		if got.toString() != tc.expected {
			t.Errorf("Test case constructing %d expected %s but received %s", tc.i, tc.expected, got.toString())
		}
	}
}

func TestConstruct(t *testing.T) {
	testCases := []struct {
		i        Integer
		expected Integer
	}{
		{Integer{}, Integer{digits: []uint8{0}, constructed: true, int_sign: true}},
		{Integer{digits: []uint8{0, 2}}, Integer{digits: []uint8{2}, constructed: true}},
		{Integer{digits: []uint8{4, 0, 2}}, Integer{digits: []uint8{4, 0, 2}, constructed: true}},
		{Integer{digits: []uint8{4, 0}}, Integer{digits: []uint8{4, 0}, constructed: true}},
		{Integer{digits: []uint8{4, 2, 0, 0, 0}}, Integer{digits: []uint8{4, 2, 0, 0, 0}, constructed: true}},
		{Integer{digits: []uint8{0, 0, 0}}, Integer{digits: []uint8{0}, constructed: true}},
		{Integer{digits: []uint8{0, 0, 0, 1, 0}}, Integer{digits: []uint8{1, 0}, constructed: true}},
	}
	for _, tc := range testCases {
		var got = tc.i.construct()
		if !got.equals(tc.expected) {
			t.Errorf("Test case constructing %s expected %s but received %s", tc.i.toString(), tc.expected.toString(), got.toString())
		}
	}
}

func TestConstructPanic(t *testing.T) {
	testCases := []struct {
		i Integer
	}{
		{Integer{constructed: true}},
		{Integer{digits: []uint8{4}, constructed: true}},
	}
	for _, tc := range testCases {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Test case constructing %s expecting panic did not panic", tc.i.toString())
			}
		}()
		tc.i.construct()
	}
}

func TestAddDigit(t *testing.T) {
	testCases := []struct {
		starting Integer
		input    uint8
		expected Integer
	}{
		{Integer{}, 0, Integer{digits: []uint8{0}}},
		{Integer{}, 4, Integer{digits: []uint8{4}}},
		{Integer{digits: []uint8{4}}, 2, Integer{digits: []uint8{4, 2}}},
		{Integer{digits: []uint8{0, 8}}, 0, Integer{digits: []uint8{0, 8, 0}}},
	}
	for _, tc := range testCases {
		got := tc.starting.addDigit(tc.input, false)
		if !got.equals(tc.expected) {
			t.Errorf("For test case expecting %s after adding %d, got %s", tc.expected.toString(), tc.input, got.toString())
		}
	}
}

func TestCompareIntegers(t *testing.T) {
	testCases := []struct {
		left     Integer
		right    Integer
		expected CompareResult
	}{
		{Integer{}, Integer{digits: []uint8{0}}, EQUAL_TO},
		{Integer{digits: []uint8{2, 1}, int_sign: true}, Integer{digits: []uint8{2, 1}, int_sign: true}, EQUAL_TO},
		{Integer{digits: []uint8{4, 6, 9}, int_sign: true}, Integer{digits: []uint8{4, 6, 9}, int_sign: true}, EQUAL_TO},
		{Integer{digits: []uint8{4, 9}, int_sign: true}, Integer{digits: []uint8{7}, int_sign: true}, GREATER_THAN},
		{Integer{digits: []uint8{7}, int_sign: true}, Integer{digits: []uint8{7, 0}, int_sign: true}, LESSER_THAN},
		{Integer{digits: []uint8{7}, int_sign: true}, Integer{digits: []uint8{0, 7}, int_sign: true}, EQUAL_TO},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{2}}, LESSER_THAN},
		{Integer{digits: []uint8{7}}, Integer{digits: []uint8{7, 0}}, GREATER_THAN},
		{Integer{digits: []uint8{2, 3}, int_sign: true}, Integer{digits: []uint8{3, 1}, int_sign: true}, LESSER_THAN},
	}
	for _, tc := range testCases {
		tc.left = tc.left.construct()
		tc.right = tc.right.construct()
		got := tc.left.compare(tc.right)
		if got != tc.expected {
			t.Errorf("For test case comparing %s with %s, got %d but expected %d", tc.left.toString(), tc.right.toString(), got, tc.expected)
		}
	}
}

func TestAddDigitPanic(t *testing.T) {
	testCases := []struct {
		starting Integer
		input    uint8
	}{
		{Integer{}, 11},
		{Integer{constructed: true}, 2},
	}
	for _, tc := range testCases {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Test case adding %d to %s expecting panic did not panic", tc.input, tc.starting.toString())
			}
		}()
		tc.starting.addDigit(tc.input, false)
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		augend   Integer
		addend   Integer
		expected Integer
	}{
		// two positive nums
		{
			Integer{},
			Integer{},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1}, int_sign: true},
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{3}, int_sign: true},
		},
		{
			Integer{digits: []uint8{8, 9}, int_sign: true},
			Integer{digits: []uint8{2, 4}, int_sign: true},
			Integer{digits: []uint8{1, 1, 3}, int_sign: true},
		},
		// two negative nums
		{
			Integer{digits: []uint8{6}, int_sign: false},
			Integer{digits: []uint8{7}, int_sign: false},
			Integer{digits: []uint8{1, 3}, int_sign: false},
		},
		{
			Integer{digits: []uint8{8, 9}, int_sign: false},
			Integer{digits: []uint8{2, 4}, int_sign: false},
			Integer{digits: []uint8{1, 1, 3}, int_sign: false},
		},
		// positive + negative
		{
			Integer{digits: []uint8{5}, int_sign: true},
			Integer{digits: []uint8{2}, int_sign: false},
			Integer{digits: []uint8{3}, int_sign: true},
		},
		{
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{5}, int_sign: false},
			Integer{digits: []uint8{3}, int_sign: false},
		},
		{
			Integer{digits: []uint8{4, 2, 3}, int_sign: true},
			Integer{digits: []uint8{6, 2}, int_sign: false},
			Integer{digits: []uint8{3, 6, 1}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1, 2, 3}, int_sign: true},
			Integer{digits: []uint8{1, 1, 9}, int_sign: false},
			Integer{digits: []uint8{4}, int_sign: true},
		},
		{
			Integer{digits: []uint8{6, 2}, int_sign: true},
			Integer{digits: []uint8{4, 2, 3}, int_sign: false},
			Integer{digits: []uint8{3, 6, 1}, int_sign: false},
		},
		{
			Integer{digits: []uint8{1, 1, 9}, int_sign: true},
			Integer{digits: []uint8{1, 2, 3}, int_sign: false},
			Integer{digits: []uint8{4}, int_sign: false},
		},
		// negative + positive
		{
			Integer{digits: []uint8{5}, int_sign: false},
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{3}, int_sign: false},
		},
		{
			Integer{digits: []uint8{2}, int_sign: false},
			Integer{digits: []uint8{5}, int_sign: true},
			Integer{digits: []uint8{3}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1, 1, 9}, int_sign: false},
			Integer{digits: []uint8{1, 2, 3}, int_sign: true},
			Integer{digits: []uint8{4}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1, 2, 3}, int_sign: false},
			Integer{digits: []uint8{1, 1, 9}, int_sign: true},
			Integer{digits: []uint8{4}, int_sign: false},
		},
	}
	for _, tc := range testCases {
		tc.augend = tc.augend.construct()
		tc.addend = tc.addend.construct()
		tc.expected = tc.expected.construct()
		var got = tc.augend.add(tc.addend)
		if !got.equals(tc.expected) {
			t.Errorf("Test case adding %s and %s expected %s but got %s", tc.augend.toString(), tc.addend.toString(), tc.expected.toString(), got.toString())
		}
	}
}

func TestAddPanic(t *testing.T) {
	testCases := []struct {
		augend Integer
		addend Integer
	}{
		{Integer{}, Integer{}},
		{Integer{digits: []uint8{4}, constructed: true}, Integer{digits: []uint8{4}}},
		{Integer{}, Integer{constructed: true}},
	}
	for _, tc := range testCases {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Test case adding %s and %s expecting panic did not panic", tc.augend.toString(), tc.addend.toString())
			}
		}()
		tc.augend.add(tc.addend)
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		multiplicand Integer
		multiplier   Integer
		product      Integer
	}{
		// two positive nums
		{
			Integer{},
			Integer{},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{0}, int_sign: true},
			Integer{digits: []uint8{2, 1}, int_sign: true},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{2, 2}, int_sign: true},
			Integer{digits: []uint8{0}, int_sign: true},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1}, int_sign: true},
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{2}, int_sign: true},
		},
		{
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{3}, int_sign: true},
			Integer{digits: []uint8{6}, int_sign: true},
		},
		{
			Integer{digits: []uint8{8, 9}, int_sign: true},
			Integer{digits: []uint8{2, 4}, int_sign: true},
			Integer{digits: []uint8{2, 1, 3, 6}, int_sign: true},
		},
		// two negative nums
		{
			Integer{digits: []uint8{6}, int_sign: false},
			Integer{digits: []uint8{7}, int_sign: false},
			Integer{digits: []uint8{4, 2}, int_sign: true},
		},
		{
			Integer{digits: []uint8{1, 2}, int_sign: false},
			Integer{digits: []uint8{2, 4}, int_sign: false},
			Integer{digits: []uint8{2, 8, 8}, int_sign: true},
		},
		// positive * negative
		{
			Integer{digits: []uint8{5}, int_sign: true},
			Integer{digits: []uint8{0}, int_sign: false},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{5}, int_sign: true},
			Integer{digits: []uint8{2}, int_sign: false},
			Integer{digits: []uint8{1, 0}, int_sign: false},
		},
		// negative * positive
		{
			Integer{digits: []uint8{0}, int_sign: false},
			Integer{digits: []uint8{2}, int_sign: true},
			Integer{digits: []uint8{0}, int_sign: true},
		},
		{
			Integer{digits: []uint8{5}, int_sign: false},
			Integer{digits: []uint8{3}, int_sign: true},
			Integer{digits: []uint8{1, 5}, int_sign: false},
		},
	}
	for _, tc := range testCases {
		tc.multiplicand = tc.multiplicand.construct()
		tc.multiplier = tc.multiplier.construct()
		tc.product = tc.product.construct()
		var got = tc.multiplicand.multiply(tc.multiplier)
		if !got.equals(tc.product) {
			t.Errorf("Test case multiplying %s and %s expected %s but got %s", tc.multiplicand.toString(), tc.multiplier.toString(), tc.product.toString(), got.toString())
		}
	}
}

func TestMultiplyPanic(t *testing.T) {
	testCases := []struct {
		left  Integer
		right Integer
	}{
		{Integer{}, Integer{}},
		{Integer{digits: []uint8{4}, constructed: true}, Integer{digits: []uint8{4}}},
		{Integer{}, Integer{constructed: true}},
	}
	for _, tc := range testCases {
		defer func() {
			r := recover()
			if r == nil {
				t.Errorf("Test case adding %s and %s expecting panic did not panic", tc.left.toString(), tc.right.toString())
			}
		}()
		tc.left.multiply(tc.right)
	}
}

func TestLongDivision(t *testing.T) {
	testCases := []struct {
		x int
		y int
		q int
		r int
	}{
		{x: 3, y: 1, q: 3, r: 0},
		{x: 11, y: 7, q: 1, r: 4},
		{x: -11, y: 7, q: -1, r: -4},
		{x: 11, y: -7, q: -1, r: 4},
		{x: -11, y: -7, q: 1, r: -4},
		{x: 27, y: 5, q: 5, r: 2},
		{x: -27, y: 5, q: -5, r: -2},
		{x: 27, y: -5, q: -5, r: 2},
		{x: -27, y: -5, q: 5, r: -2},
		{x: 3, y: 8, q: 0, r: 3},
		{x: -3, y: 8, q: 0, r: -3},
		{x: 3, y: -8, q: 0, r: 3},
		{x: -3, y: -8, q: 0, r: -3},
	}
	for _, tc := range testCases {
		dividend := constructInt(tc.x)
		divisor := constructInt(tc.y)
		quotient := constructInt(tc.q)
		remainder := constructInt(tc.r)
		var got1, got2 = dividend.longDivision(divisor)
		if !got1.equals(quotient) {
			t.Errorf("Test case long dividing %s and %s got quotient %s", dividend.toString(), divisor.toString(), got1.toString())
		}
		if !got2.equals(remainder) {
			t.Errorf("Test case long dividing %s and %s got remainder %s", dividend.toString(), divisor.toString(), got2.toString())
		}
	}
}
