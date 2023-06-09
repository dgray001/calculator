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
		return_string += strconv.FormatInt(int64(digit), 10)
	}
	if !i.int_sign && return_string != "0" {
		return "-" + return_string
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

func (i Integer) addDigit(digit uint8, add_to_front bool) Integer {
	if i.constructed {
		panic("Cannot add digit when integer is constructed")
	}
	if digit > 9 {
		panic("Cannot add " + strconv.FormatInt(int64(digit), 10) + " since it is > 9")
	}
	if add_to_front {
		i.digits = append([]uint8{digit}, i.digits...)
	} else {
		i.digits = append(i.digits, digit)
	}
	return i
}

func (i Integer) construct() Integer {
	if i.constructed {
		panic("Cannot construct an integer that is already constructed")
	}
	i.constructed = true
	var slice_index = 0
	for j := 0; j < len(i.digits)-1; j++ {
		if i.digits[j] == 0 {
			slice_index++
		} else {
			break
		}
	}
	i.digits = i.digits[slice_index:]
	if len(i.digits) == 0 {
		i.digits = append(i.digits, 0)
		i.int_sign = true
	}
	return i
}

func (i Integer) invert() Integer {
	i.int_sign = !i.int_sign
	return i
}

func (i Integer) abs() Integer {
	i.int_sign = true
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
		for k := 0; k < len(i.digits); k++ {
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

	if subtraction {
		var abs_compare = i.abs().compare(j.abs())
		var switch_operands = false
		switch abs_compare {
		case EQUAL_TO:
			return return_integer.construct()
		case GREATER_THAN:
			switch_operands = false
			return_integer.int_sign = i.int_sign
		case LESSER_THAN:
			switch_operands = true
			return_integer.int_sign = j.int_sign
		default:
			panic("Invalid compare result when subtracting")
		}
		var minuend Integer
		var subtrahend Integer
		if switch_operands {
			minuend = j
			subtrahend = i
		} else {
			minuend = i
			subtrahend = j
		}
		for place := 0; place < len(minuend.digits); place++ {
			var minuend_v = minuend.digits[len(minuend.digits)-1-place]
			var subtrahend_v uint8 = 0
			if place < len(subtrahend.digits) {
				subtrahend_v = subtrahend.digits[len(subtrahend.digits)-1-place]
			}
			if subtrahend_v > minuend_v {
				minuend_v += 10
				minuend.digits[len(minuend.digits)-2-place]--
			}
			return_integer = return_integer.addDigit(minuend_v-subtrahend_v, true)
		}
	} else {
		var remainder uint8 = 0
		for place := 0; place < len(i.digits) || place < len(j.digits); place++ {
			var i_v uint8 = 0
			var j_v uint8 = 0
			if place < len(i.digits) {
				i_v = i.digits[len(i.digits)-1-place]
			}
			if place < len(j.digits) {
				j_v = j.digits[len(j.digits)-1-place]
			}
			var sum = i_v + j_v + remainder
			return_integer = return_integer.addDigit(sum%10, true)
			remainder = sum / 10
		}
		return_integer = return_integer.addDigit(remainder, true)
	}
	return return_integer.construct()
}

func (i Integer) subtract(j Integer) Integer {
	j.int_sign = !j.int_sign
	return i.add(j)
}
