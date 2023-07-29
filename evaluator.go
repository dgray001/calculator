package main

import (
	"errors"
)

func (node *AstNode) evaluate() (Value, error) {
	if len(node.values) == 0 {
		return Value{}, errors.New("Can't evaluate node without values")
	}
	// pass for recursive evaluation
	for i, value := range node.values {
		var evaluated_value, evaluate_error = value.evaluate()
		if evaluate_error != nil {
			return Value{}, evaluate_error
		}
		node.values[i] = evaluated_value
	}
	var pass_error error = nil
	var values_deleted int = 0
	// pass for exponentiation
	// pass for multiplication
	values_deleted, pass_error = node.evaluatePass(MULTIPLY, DIVIDE, values_deleted)
	if pass_error != nil {
		return Value{}, pass_error
	}
	// pass for addition
	values_deleted, pass_error = node.evaluatePass(PLUS, MINUS, values_deleted)
	if pass_error != nil {
		return Value{}, pass_error
	}
	if len(node.values) != 1 {
		return Value{}, errors.New("Evaluation didn't result in a value")
	}
	if node.function != nil {
		return evaluateFunction(*node.function, node.values[0])
	}
	return node.values[0], nil
}

func (node *AstNode) evaluatePass(start Token, end Token, values_deleted int) (int, error) {
	var has_unary_operator = len(node.operators) == (len(node.values) + values_deleted)
	for i, operator := range node.operators {
		if operator >= start && operator <= end {
			// check for unary operator
			if i == 0 && has_unary_operator {
				if !operator.isUnaryOperator() {
					return values_deleted, errors.New("Non-unary operator in a unary operator position")
				}
				var result, unary_error = unaryOperation(operator, node.values[i])
				if unary_error != nil {
					return values_deleted, unary_error
				}
				node.values[i] = result
			} else {
				if !operator.isBinaryOperator() {
					return values_deleted, errors.New("Non-binary operator in a binary operator position")
				}
				var j = i - values_deleted
				if has_unary_operator {
					j--
				}
				var result, binary_error = binaryOperation(operator, node.values[j], node.values[j+1])
				if binary_error != nil {
					return values_deleted, binary_error
				}
				node.values = append(node.values[0:j], node.values[j+1:]...)
				values_deleted++
				node.values[j] = result
			}
		}
	}
	return values_deleted, nil
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
