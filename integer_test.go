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

func TestConstruct(t *testing.T) {
	testCases := []struct {
		i        Integer
		expected Integer
	}{
		{Integer{}, Integer{digits: []uint8{0}, constructed: true}},
		{Integer{digits: []uint8{0, 2}}, Integer{digits: []uint8{0, 2}, constructed: true}},
		{Integer{digits: []uint8{4, 0, 2}}, Integer{digits: []uint8{4, 0, 2}, constructed: true}},
		{Integer{digits: []uint8{4, 0}}, Integer{digits: []uint8{4}, constructed: true}},
		{Integer{digits: []uint8{4, 2, 0, 0, 0}}, Integer{digits: []uint8{4, 2}, constructed: true}},
		{Integer{digits: []uint8{0, 0, 0}}, Integer{digits: []uint8{0}, constructed: true}},
		{Integer{digits: []uint8{0, 0, 0, 1, 0}}, Integer{digits: []uint8{0, 0, 0, 1}, constructed: true}},
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
		tc.starting.addDigit(tc.input)
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		left     Integer
		right    Integer
		expected Integer
	}{
		{Integer{constructed: true}, Integer{constructed: true}, Integer{digits: []uint8{0}, constructed: true}},
		{
			Integer{digits: []uint8{1}, constructed: true},
			Integer{digits: []uint8{2}, constructed: true},
			Integer{digits: []uint8{3}, constructed: true},
		},
		{
			Integer{digits: []uint8{1, 0}, constructed: true},
			Integer{digits: []uint8{0}, constructed: true},
			Integer{digits: []uint8{1}, constructed: true},
		},
		{
			Integer{digits: []uint8{1, 1}, constructed: true},
			Integer{digits: []uint8{0, 1}, constructed: true},
			Integer{digits: []uint8{1, 2}, constructed: true},
		},
		{
			Integer{digits: []uint8{9, 0}, constructed: true},
			Integer{digits: []uint8{3}, constructed: true},
			Integer{digits: []uint8{2, 1}, constructed: true},
		},
		{
			Integer{digits: []uint8{9, 8}, constructed: true},
			Integer{digits: []uint8{4, 2}, constructed: true},
			Integer{digits: []uint8{3, 1, 1}, constructed: true},
		},
	}
	for _, tc := range testCases {
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
