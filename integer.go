package main

import (
	"strconv"
)

type Integer struct {
	int_sign    bool
	digits      []uint8
	constructed bool
}

func newInteger() Integer {
	return Integer{
		int_sign:    true,
		digits:      make([]uint8, 0),
		constructed: false,
	}
}

func copyInt(i Integer) Integer {
	digits := make([]uint8, len(i.digits))
	copy(digits, i.digits)
	return Integer{
		int_sign:    i.int_sign,
		digits:      digits,
		constructed: i.constructed,
	}
}

func constructInt(i int) Integer {
	s := strconv.Itoa(i)
	j := newInteger()
	for _, sd := range s {
		if sd == '-' {
			j.int_sign = false
			continue
		}
		s, e := strconv.Atoi(string(sd))
		if e != nil {
			panic("Non integer rune being used to construct int")
		}
		j = j.addDigit(uint8(s), false)
	}
	return j.construct()
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
	ic := copyInt(i)
	ic.int_sign = !ic.int_sign
	return ic
}

func (i Integer) abs() Integer {
	ic := copyInt(i)
	ic.int_sign = true
	return ic
}

func (i Integer) increment() Integer {
	var one = newInteger().addDigit(ONE.toInt(), false).construct()
	return i.add(one)
}

func (i Integer) decrement() Integer {
	var one = newInteger().addDigit(ONE.toInt(), false).construct()
	return i.subtract(one)
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

/*
For reference:
https://math.stackexchange.com/questions/623449/negative-number-divided-by-positive-number-what-would-be-remainder and
https://ericlippert.com/2013/10/31/math-from-scratch-part-eleven-integer-division/
So following the second reference, we have:
i / j = q + r / j so that i = j * q + r
(-i) / j == i / (-j) == -(i / j)
so once we have quotient we can calculate remainder by:
r = i - j * q
*/
func (i Integer) longDivision(j Integer) (Integer, Integer) {
	if j.isZero() {
		panic("Divide by zero error")
	}
	ic := copyInt(i)
	jc := copyInt(j)
	end_sign := ic.int_sign == jc.int_sign
	ic.int_sign = true
	jc.int_sign = true
	quotient := constructInt(0)
	for ic.compare(jc) != LESSER_THAN {
		quotient = quotient.add(constructInt(1))
		ic = ic.subtract(jc)
	}
	if !quotient.isZero() {
		quotient.int_sign = end_sign
	}
	if !ic.isZero() {
		ic.int_sign = i.int_sign
	}
	return quotient, ic
}

func (i Integer) add(j Integer) Integer {
	if !i.constructed || !j.constructed {
		panic("Cannot add unconstructed integers")
	}
	return_integer := newInteger()
	subtraction := false
	if !i.int_sign && !j.int_sign {
		return_integer.int_sign = false
	} else if !i.int_sign || !j.int_sign {
		subtraction = true
	}

	if subtraction {
		abs_compare := i.abs().compare(j.abs())
		switch_operands := false
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
			minuend = copyInt(j)
			subtrahend = copyInt(i)
		} else {
			minuend = copyInt(i)
			subtrahend = copyInt(j)
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
	return i.add(j.invert())
}

func (i Integer) multiply(j Integer) Integer {
	if !i.constructed || !j.constructed {
		panic("Cannot multiply unconstructed integers")
	}
	var return_integer = newInteger().construct()
	// the "long multiplication" algorithm
	for multiplicand := 0; multiplicand < len(i.digits); multiplicand++ {
		var i_v = i.digits[len(i.digits)-1-multiplicand]
		var addend = newInteger()
		var remainder uint8 = 0
		for multiplier := 0; multiplier < len(j.digits); multiplier++ {
			var j_v = j.digits[len(j.digits)-1-multiplier]
			var product = i_v*j_v + remainder
			addend = addend.addDigit(product%10, true)
			remainder = product / 10
		}
		addend = addend.addDigit(remainder, true)
		for extra_zero := 0; extra_zero < multiplicand; extra_zero++ {
			addend = addend.addDigit(0, false)
		}
		return_integer = return_integer.add(addend.construct())
	}
	if i.int_sign != j.int_sign && !return_integer.isZero() {
		return_integer.int_sign = false
	} else {
		return_integer.int_sign = true
	}
	return return_integer
}

func (i Integer) toRational() RationalNumber {
	return newRationalNumber(i, constructInt(1))
}
