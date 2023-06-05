package main

func unaryOperation(operator Token, value Value) (Value, error) {
	switch operator {
	case PLUS:
		switch value.value_type {
		case INTEGER:
			return value, nil
		default:
			panic("Invalid value type for unary plus operation: " + value.value_type.toString())
		}
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

func binaryOperation(operator Token, value1 Value, value2 Value) (Value, error) {
	switch operator {
	case PLUS:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			panic("Cannot add node value types")
		}
		return intValue(value1.integer.add(*value2.integer)), nil
	case MINUS:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			panic("Cannot add node value types")
		}
		return intValue(value1.integer.subtract(*value2.integer)), nil
	case MULTIPLY:
		panic("Not implemented")
	default:
		panic("Invalid binary operator: " + operator.toString())
	}
}
