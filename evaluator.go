package main

import "errors"

func (node *AstNode) evaluate() (Value, error) {
	if len(node.values) == 0 {
		return Value{}, errors.New("Can't evaluate node without values")
	}
	// pass for recursive evaluation
	for _, value := range node.values {
		value.evaluate()
	}
	// pass for exponentiation
	// pass for multiplication
	node.evaluatePass(MULTIPLY, MULTIPLY)
	// pass for addition
	node.evaluatePass(PLUS, MINUS)
	return node.values[0], nil
}

func (node *AstNode) evaluatePass(start Token, end Token) {
	for i, operator := range node.operators {
		if operator >= start && operator <= end {
			// check for unary operator
			if i == 0 && len(node.operators) == len(node.values) {
			}
		}
	}
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
