package main

import (
	"strconv"
)

type Integer struct {
	digits []uint8
}

func newInteger() Integer {
	return Integer{
		digits: []uint8{},
	}
}

func (i Integer) toString() string {
	var return_string = ""
	for _, digit := range i.digits {
		return_string = strconv.FormatInt(int64(digit), 10) + return_string
	}
	return return_string
}

func (left Integer) equals(right Integer) bool {
	if len(left.digits) != len(right.digits) {
		return false
	}
	for i := 0; i < len(left.digits); i++ {
		if left.digits[i] != right.digits[i] {
			return false
		}
	}
	return true
}

func (i Integer) addDigit(digit uint8) Integer {
	if digit > 9 {
		panic("Cannot add " + strconv.FormatInt(int64(digit), 10) + " since it is > 9")
	}
	i.digits = append(i.digits, digit)
	return i
}
