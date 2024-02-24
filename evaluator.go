package main

import (
	"errors"
)

func (node *AstNode) evaluate() (Value, error) {
	if len(node.values) == 0 {
		return Value{}, errors.New("can't evaluate node without values")
	}
	// pass for recursive evaluation
	for i, value := range node.values {
		var evaluated_value, evaluate_error = value.evaluate()
		if evaluate_error != nil {
			return Value{}, evaluate_error
		}
		node.values[i] = evaluated_value
	}
	// pass for exponentiation
	// pass for multiplication
	pass_error := node.evaluatePass(MULTIPLY, DIVIDE)
	if pass_error != nil {
		return Value{}, pass_error
	}
	// pass for addition
	pass_error = node.evaluatePass(PLUS, MINUS)
	if pass_error != nil {
		return Value{}, pass_error
	}
	if len(node.values) != 1 || len(node.operators) > 0 {
		return Value{}, errors.New("evaluation didn't result in a value")
	}
	// pass for function
	if node.function != nil {
		return evaluateFunction(*node.function, node.values[0])
	}
	return node.values[0], nil
}

func (node *AstNode) evaluatePass(start Token, end Token) error {
	var has_unary_operator = len(node.operators) == len(node.values)
	node_operators := make([]Token, len(node.operators))
	copy(node_operators, node.operators)
	operators_removed := 0
	values_removed := 0
	for i, operator := range node_operators {
		if operator < start || operator > end {
			continue
		}
		// check for unary operator
		if i == 0 && has_unary_operator {
			if !operator.isUnaryOperator() {
				return errors.New("non-unary operator in a unary operator position")
			}
			result, unary_error := unaryOperation(operator, node.values[i])
			if unary_error != nil {
				return unary_error
			}
			node.values[i] = result
		} else {
			if !operator.isBinaryOperator() {
				return errors.New("non-binary operator in a binary operator position")
			}
			j := i - values_removed
			if has_unary_operator {
				j--
			}
			var result, binary_error = binaryOperation(operator, node.values[j], node.values[j+1])
			if binary_error != nil {
				return binary_error
			}
			node.values = append(node.values[0:j], node.values[j+1:]...)
			values_removed++
			node.values[j] = result
		}
		node.operators = append(node.operators[0:i-operators_removed], node.operators[i+1-operators_removed:]...)
		operators_removed++
	}
	return nil
}

func (value *Value) evaluate() (Value, error) {
	switch value.value_type {
	case INTEGER:
		return *value, nil
	case AST_NODE:
		return value.ast_node.evaluate()
	default:
		return Value{}, nil
	}
}
