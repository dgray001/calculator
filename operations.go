package main

import "errors"

func unaryOperation(operator Token, value Value) (Value, error) {
	switch operator {
	case PLUS:
		switch value.value_type {
		case INTEGER:
			return value, nil
		case RATIONAL_NUMBER:
			return value, nil
		default:
			return Value{}, errors.New("invalid value type for unary plus operation: " + value.value_type.toString())
		}
	case MINUS:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.invert()), nil
		case RATIONAL_NUMBER:
			var v = value.rational.invert()
			value.rational = &v
			return value, nil
		default:
			return Value{}, errors.New("invalid value type for unary minus operation: " + value.value_type.toString())
		}
	default:
		return Value{}, errors.New("invalid unary operator: " + operator.toString())
	}
}

func binaryOperation(operator Token, value1 Value, value2 Value) (Value, error) {
	switch operator {

	case PLUS:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			return Value{}, errors.New("cannot add node value types")
		}
		if value1.value_type == RATIONAL_NUMBER || value2.value_type == RATIONAL_NUMBER {
			return rationalValue(value1.asRational().add(value2.asRational())), nil
		}
		return intValue(value1.integer.add(*value2.integer)), nil

	case MINUS:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			return Value{}, errors.New("cannot subtract node value types")
		}
		if value1.value_type == RATIONAL_NUMBER || value2.value_type == RATIONAL_NUMBER {
			return rationalValue(value1.asRational().subtract(value2.asRational())), nil
		}
		return intValue(value1.integer.subtract(*value2.integer)), nil

	case MULTIPLY:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			return Value{}, errors.New("cannot multiply node value types")
		}
		if value1.value_type == RATIONAL_NUMBER || value2.value_type == RATIONAL_NUMBER {
			return rationalValue(value1.asRational().multiply(value2.asRational())), nil
		}
		return intValue(value1.integer.multiply(*value2.integer)), nil

	case INT_DIVIDE:
		if value1.value_type != INTEGER || value2.value_type != INTEGER {
			return Value{}, errors.New("can only int divide int value types")
		}
		if value2.integer.isZero() {
			return Value{}, errors.New("cannot divide by zero")
		}
		q, _ := value1.integer.longDivision(*value2.integer)
		return intValue(q), nil

	case DIVIDE:
		if value1.value_type == AST_NODE || value2.value_type == AST_NODE {
			return Value{}, errors.New("cannot divide node value types")
		}
		if value1.value_type == RATIONAL_NUMBER || value2.value_type == RATIONAL_NUMBER {
			return rationalValue(value1.asRational().divide(value2.asRational())), nil
		}
		if value2.integer.isZero() {
			return Value{}, errors.New("cannot divide by zero")
		}
		return rationalValue(newRationalNumber(*value1.integer, *value2.integer)), nil

	default:
		return Value{}, errors.New("invalid binary operator: " + operator.toString())
	}
}

func evaluateFunction(function Token, value Value) (Value, error) {
	switch function {

	case INCREMENT:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.increment()), nil
		case RATIONAL_NUMBER:
			return rationalValue(value.rational.increment()), nil
		default:
			return Value{}, errors.New("invalid value type for increment function: " + value.value_type.toString())
		}

	case DECREMENT:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.decrement()), nil
		case RATIONAL_NUMBER:
			return rationalValue(value.rational.decrement()), nil
		default:
			return Value{}, errors.New("invalid value type for decrement function: " + value.value_type.toString())
		}

	case ABSOLUTE:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.abs()), nil
		case RATIONAL_NUMBER:
			return rationalValue(value.rational.abs()), nil
		default:
			return Value{}, errors.New("invalid value type for absolute value function: " + value.value_type.toString())
		}

	case INVERT:
		switch value.value_type {
		case INTEGER:
			return intValue(value.integer.invert()), nil
		case RATIONAL_NUMBER:
			return rationalValue(value.rational.invert()), nil
		default:
			return Value{}, errors.New("invalid value type for absolute value function: " + value.value_type.toString())
		}

	default:
		return Value{}, errors.New("invalid function: " + function.toString())
	}
}
