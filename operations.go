package main

func unaryOperation(operator Token, value Value) (Value, error) {
	switch operator {
	case PLUS: // doesn't actually do anything
		return value, nil
	case MINUS:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.invert()), nil
		default:
			panic("Invalid value type for unary minus operation: " + value.value_type.toString())
		}
	default:
		panic("Invalid unary operator: " + operator.toString())
	}
}
