package main

type Token int8

const (
	// ints
	ZERO Token = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE

	// operators
	PLUS
	MINUS
	MULTIPLY

	// other
	OPEN_PAREN
	CLOSE_PAREN

	// functions
	INCREMENT
	DECREMENT

	// to loop over possible tokens
	tokenLimit
)

func (token Token) isInt() bool {
	if token >= ZERO && token <= NINE {
		return true
	}
	return false
}

func (token Token) isOperator() bool {
	if token >= PLUS && token <= MULTIPLY {
		return true
	}
	return false
}

func (token Token) isFunction() bool {
	if token >= INCREMENT && token <= DECREMENT {
		return true
	}
	return false
}

func (token Token) toInt() uint8 {
	switch token {
	case ZERO:
		return 0
	case ONE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	default:
		panic("Token is not an integer: " + token.toString())
	}
}

func (token Token) toString() string {
	switch token {
	case ZERO:
		return "0"
	case ONE:
		return "1"
	case TWO:
		return "2"
	case THREE:
		return "3"
	case FOUR:
		return "4"
	case FIVE:
		return "5"
	case SIX:
		return "6"
	case SEVEN:
		return "7"
	case EIGHT:
		return "8"
	case NINE:
		return "9"
	case PLUS:
		return "+"
	case MINUS:
		return "-"
	case MULTIPLY:
		return "*"
	case OPEN_PAREN:
		return "("
	case CLOSE_PAREN:
		return ")"
	case INCREMENT:
		return "inc"
	case DECREMENT:
		return "dec"
	default:
		return ""
	}
}

func (token Token) toRunes() []rune {
	return []rune(token.toString())
}
