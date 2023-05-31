package main

type Token int8

const (
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
	PLUS
	MINUS
	MULTIPLY
	OPEN_PAREN
	CLOSE_PAREN
	tokenLimit // to loop over possible tokens
)

func (token Token) isInt() bool {
	switch token {
	case ZERO:
	case ONE:
	case TWO:
	case THREE:
	case FOUR:
	case FIVE:
	case SIX:
	case SEVEN:
	case EIGHT:
	case NINE:
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
	default:
		return ""
	}
}

func (token Token) toRune() rune {
	return []rune(token.toString())[0]
}
