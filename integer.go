package main

import (
	"strconv"
)

type Integer struct {
	digits      []uint8
	constructed bool
}

func newInteger() Integer {
	return Integer{
		digits:      []uint8{},
		constructed: false,
	}
}

func (i Integer) toString() string {
	var return_string = ""
	for _, digit := range i.digits {
		return_string = strconv.FormatInt(int64(digit), 10) + return_string
	}
	return return_string
}

func (left Integer) equals(untyped interface{}) bool {
	var right = untyped.(Integer)
	if left.constructed != right.constructed {
		return false
	}
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

func (i Integer) addDigit(digit uint8, add_to_end bool) Integer {
	if i.constructed {
		panic("Cannot add digit when integer is constructed")
	}
	if digit > 9 {
		panic("Cannot add " + strconv.FormatInt(int64(digit), 10) + " since it is > 9")
	}
	if add_to_end {
		i.digits = append(i.digits, digit)
	} else {
		i.digits = append([]uint8{digit}, i.digits...)
	}
	return i
}

func (i Integer) construct() Integer {
	if i.constructed {
		panic("Cannot construct an integer that is already constructed")
	}
	i.constructed = true
	var slice_index = len(i.digits)
	for j := len(i.digits) - 1; j >= 0; j-- {
		if i.digits[j] == 0 {
			slice_index--
		} else {
			break
		}
	}
	i.digits = i.digits[0:slice_index]
	if len(i.digits) == 0 {
		i.digits = append(i.digits, 0)
	}
	return i
}

func (i Integer) add(j Integer) Integer {
	if !i.constructed || !j.constructed {
		panic("Cannot add unconstructed integers")
	}
	var return_integer = Integer{}
	var remainder uint8 = 0
	for place := 0; place < len(i.digits) || place < len(j.digits); place++ {
		var i_v uint8 = 0
		var j_v uint8 = 0
		if place < len(i.digits) {
			i_v = i.digits[place]
		}
		if place < len(j.digits) {
			j_v = j.digits[place]
		}
		var sum = i_v + j_v + remainder
		return_integer = return_integer.addDigit(sum%10, true)
		remainder = sum / 10
	}
	return_integer = return_integer.addDigit(remainder, true)
	return return_integer.construct()
}
