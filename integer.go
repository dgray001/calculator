package main

import (
	"strconv"
)

type CompareResult int8

const (
	ERROR_CompareResult CompareResult = iota
	LESSER_THAN
	EQUAL_TO
	GREATER_THAN
)

func (compare_result CompareResult) invert() CompareResult {
	switch compare_result {
	case LESSER_THAN:
		return GREATER_THAN
	case GREATER_THAN:
		return LESSER_THAN
	default:
		return compare_result
	}
}

type Integer struct {
	int_sign    bool
	digits      []uint8
	constructed bool
}

func newInteger() Integer {
	return Integer{
		int_sign:    true,
		digits:      []uint8{},
		constructed: false,
	}
}

func (i Integer) toString() string {
	var return_string = ""
	for _, digit := range i.digits {
		return_string = strconv.FormatInt(int64(digit), 10) + return_string
	}
	if !i.int_sign && return_string != "0" {
		return_string = "-" + return_string
	}
	return return_string
}

func (left Integer) equals(untyped interface{}) bool {
	var right = untyped.(Integer)
	if left.int_sign != right.int_sign {
		return false
	}
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

func (i Integer) invert() Integer {
	i.int_sign = !i.int_sign
	return i
}

func (i Integer) isZero() bool {
	return len(i.digits) == 1 && i.digits[0] == 0
}

func (i Integer) compare(j Integer) CompareResult {
	if !i.constructed || !j.constructed {
		panic("Cannot compare unconstructed ints")
	}
	if i.isZero() && j.isZero() {
		return EQUAL_TO
	}
	if i.int_sign && !j.int_sign {
		return GREATER_THAN
	} else if !i.int_sign && j.int_sign {
		return LESSER_THAN
	}
	// have to actually calculate which has a greater absolute value
	var compare_result = ERROR_CompareResult
	if len(i.digits) > len(j.digits) {
		compare_result = GREATER_THAN
	} else if len(i.digits) < len(j.digits) {
		compare_result = LESSER_THAN
	} else {
		compare_result = EQUAL_TO
		for k := len(i.digits) - 1; k >= 0; k-- {
			if i.digits[k] > j.digits[k] {
				compare_result = GREATER_THAN
				break
			} else if i.digits[k] < j.digits[k] {
				compare_result = LESSER_THAN
				break
			}
		}
	}
	if !i.int_sign && !j.int_sign {
		compare_result = compare_result.invert()
	}
	return compare_result
}

func (i Integer) add(j Integer) Integer {
	if !i.constructed || !j.constructed {
		panic("Cannot add unconstructed integers")
	}
	var return_integer = newInteger()
	var subtraction = false
	if !i.int_sign && !j.int_sign {
		return_integer.int_sign = false
	} else if !i.int_sign {
		subtraction = true
	} else if !j.int_sign {
		subtraction = true
	}

	var remainder int8 = 0
	for place := 0; place < len(i.digits) || place < len(j.digits); place++ {
		var last_digit = place+1 >= len(i.digits) && place+1 >= len(j.digits)
		var i_v int8 = 0
		var j_v int8 = 0
		if place < len(i.digits) {
			if i.int_sign {
				i_v = int8(i.digits[place])
			} else {
				i_v = -int8(i.digits[place])
			}
		}
		if place < len(j.digits) {
			if j.int_sign {
				j_v = int8(j.digits[place])
			} else {
				j_v = -int8(j.digits[place])
			}
		}
		var sum = i_v + j_v + remainder
		remainder = sum / 10
		var sum_digit = sum % 10
		if subtraction {
			//
		} else if sum_digit < 0 {
			sum_digit = -sum_digit
		}
		if sum_digit < 0 {
			if i.int_sign != j.int_sign {
				if last_digit {
					sum_digit = -sum_digit
					return_integer.int_sign = false
				} else {
					sum_digit += 10
					remainder = -1
				}
			} else {
				sum_digit = -sum_digit
			}
		}
		return_integer = return_integer.addDigit(uint8(sum_digit), true)
	}
	if remainder < 0 {
		remainder = -remainder
		return_integer.int_sign = false
	}
	return_integer = return_integer.addDigit(uint8(remainder), true)
	return return_integer.construct()
}

func (i Integer) subtract(j Integer) Integer {
	j.int_sign = !j.int_sign
	return i.add(j)
}
