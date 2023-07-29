package main

type RationalNumber struct {
	numerator     Integer
	denominator   Integer
	rational_sign bool
	simplified    bool
}

func newRationalNumber(i Integer, j Integer) RationalNumber {
	var rational_sign = i.int_sign == j.int_sign
	i.int_sign = true
	j.int_sign = true
	return RationalNumber{
		numerator:     i,
		denominator:   j,
		rational_sign: rational_sign,
		simplified:    false,
	}
}

func (i RationalNumber) toString() string {
	var return_string = ""
	if !i.rational_sign {
		return_string += "-"
	}
	return_string += i.numerator.toString()
	return_string += "/"
	return_string += i.denominator.toString()
	return return_string
}

// Note this function is for object equality so doesn't account for simplification
// Use Value::equals for mathematical equality
func (left RationalNumber) equals(untyped interface{}) bool {
	var right = untyped.(RationalNumber)
	if left.rational_sign != right.rational_sign {
		return false
	}
	if !left.numerator.equals(right.numerator) {
		return false
	}
	if !right.denominator.equals(right.denominator) {
		return false
	}
	return true
}

func (i RationalNumber) invert() RationalNumber {
	i.rational_sign = !i.rational_sign
	return i
}
