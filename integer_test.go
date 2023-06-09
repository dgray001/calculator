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
		left     Integer
		right    Integer
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
		tc.left = tc.left.construct()
		tc.right = tc.right.construct()
		tc.expected = tc.expected.construct()
		var got = tc.left.add(tc.right)
		if !got.equals(tc.expected) {
			t.Errorf("Test case adding %s and %s expected %s but got %s", tc.left.toString(), tc.right.toString(), tc.expected.toString(), got.toString())
		}
	}
}

func TestAddPanic(t *testing.T) {
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
		tc.left.add(tc.right)
	}
}
