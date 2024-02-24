package main

type RationalNumber struct {
	numerator     Integer
	denominator   Integer
	rational_sign bool
	simplified    bool
}

func newRationalNumber(i Integer, j Integer) RationalNumber {
	if j.isZero() {
		panic("Creating rational number with 0 as denominator")
	}
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

func constructRationalNumber(i int, j int) RationalNumber {
	return newRationalNumber(constructInt(i), constructInt(j))
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
	if !left.denominator.equals(right.denominator) {
		return false
	}
	return true
}

func (i RationalNumber) invert() RationalNumber {
	i.rational_sign = !i.rational_sign
	return i
}

func (i RationalNumber) abs() RationalNumber {
	i.rational_sign = true
	return i
}

func (i RationalNumber) increment() RationalNumber {
	if i.rational_sign {
		i.numerator = i.numerator.add(i.denominator)
	} else {
		i.numerator = i.numerator.subtract(i.denominator)
	}
	return i
}

func (i RationalNumber) decrement() RationalNumber {
	if i.rational_sign {
		i.numerator = i.numerator.subtract(i.denominator)
	} else {
		i.numerator = i.numerator.add(i.denominator)
	}
	return i
}

func (i RationalNumber) isZero() bool {
	return i.numerator.isZero()
}

func (i RationalNumber) compare(j RationalNumber) CompareResult {
	if i.isZero() && j.isZero() {
		return EQUAL_TO
	}
	if i.rational_sign && !j.rational_sign {
		return GREATER_THAN
	} else if !i.rational_sign && j.rational_sign {
		return LESSER_THAN
	}
	return i.numerator.multiply(j.denominator).compare(i.denominator.multiply(j.numerator))
}

func (i RationalNumber) add(j RationalNumber) RationalNumber {
	// a/b + c/d = (ad + bc) / bd
	i.numerator.int_sign = i.rational_sign
	j.numerator.int_sign = j.rational_sign
	ad := i.numerator.multiply(j.denominator)
	bc := i.denominator.multiply(j.numerator)
	bd := i.denominator.multiply(j.denominator)
	return newRationalNumber(ad.add(bc), bd)
}

func (i RationalNumber) subtract(j RationalNumber) RationalNumber {
	// a/b - c/d = (ad - bc) / bd
	i.numerator.int_sign = i.rational_sign
	j.numerator.int_sign = j.rational_sign
	ad := i.numerator.multiply(j.denominator)
	bc := i.denominator.multiply(j.numerator)
	bd := i.denominator.multiply(j.denominator)
	return newRationalNumber(ad.subtract(bc), bd)
}
