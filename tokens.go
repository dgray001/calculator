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
	tokenLimit // to loop over possible tokens
)

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
	default:
		return ""
	}
}
