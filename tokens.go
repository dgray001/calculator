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
	DIVIDE

	// parentheses
	OPEN_PAREN
	CLOSE_PAREN
	OPEN_BRACKET
	CLOSE_BRACKET

	// functions
	INCREMENT
	DECREMENT
	ABSOLUTE
	INVERT

	// to loop over possible tokens
	tokenLimit
)

func (i Token) equals(untyped interface{}) bool {
	var j = untyped.(Token)
	return i == j
}

func (token Token) isInt() bool {
	if token >= ZERO && token <= NINE {
		return true
	}
	return false
}

func (token Token) isOperator() bool {
	if token >= PLUS && token <= DIVIDE {
		return true
	}
	return false
}

func (token Token) isUnaryOperator() bool {
	switch token {
	case PLUS:
		return true
	case MINUS:
		return true
	default:
		return false
	}
}

func (token Token) isBinaryOperator() bool {
	return token.isOperator()
}

func (token Token) isParentheses() bool {
	if token >= OPEN_PAREN && token <= CLOSE_BRACKET {
		return true
	}
	return false
}

func (token Token) isOpenParens() bool {
	if token == OPEN_PAREN || token == OPEN_BRACKET {
		return true
	}
	return false
}

func (token Token) isCloseParens() bool {
	if token == CLOSE_PAREN || token == CLOSE_BRACKET {
		return true
	}
	return false
}

func (token Token) isFunction() bool {
	if token >= INCREMENT && token <= INVERT {
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
	case DIVIDE:
		return "/"
	case OPEN_PAREN:
		return "("
	case CLOSE_PAREN:
		return ")"
	case OPEN_BRACKET:
		return "["
	case CLOSE_BRACKET:
		return "]"
	case INCREMENT:
		return "inc"
	case DECREMENT:
		return "dec"
	case ABSOLUTE:
		return "abs"
	case INVERT:
		return "inv"
	default:
		panic("Unrecognized token")
	}
}

func (token Token) toRunes() []rune {
	return []rune(token.toString())
}
