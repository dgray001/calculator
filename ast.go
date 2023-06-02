package main

import (
	"errors"
	"strings"
)

type AstNode struct {
	function *Token

	values    []Value
	operators []Token

	constructingInt   *Integer
	lastAddedValue    bool
	lastAddedOperator bool
}

func newAstNode() AstNode {
	return AstNode{
		function:          nil,
		values:            make([]Value, 0),
		operators:         make([]Token, 0),
		constructingInt:   nil,
		lastAddedValue:    false,
		lastAddedOperator: false,
	}
}

func constructAstNode(fn Token, vs []Value, os []Token, in Integer, lastV bool, lastO bool) AstNode {
	return AstNode{
		function:          &fn,
		values:            vs,
		operators:         os,
		constructingInt:   &in,
		lastAddedValue:    lastV,
		lastAddedOperator: lastO,
	}
}

func (node AstNode) toString(shallow bool) string {
	var return_string strings.Builder
	return_string.WriteString("AstNode {")
	if node.function != nil {
		return_string.WriteString("\n  function: " + node.function.toString())
	}
	return_string.WriteString("\n  values: [")
	var first = true
	for _, value := range node.values {
		if !first {
			return_string.WriteString(", ")
		}
		first = false
		return_string.WriteString(value.toString(shallow))
	}
	return_string.WriteString("]")
	return_string.WriteString("\n  operators: [")
	first = true
	for _, operator := range node.operators {
		if !first {
			return_string.WriteString(", ")
		}
		first = false
		return_string.WriteString(operator.toString())
	}
	return_string.WriteString("]")
	return_string.WriteString("\n}")
	return return_string.String()
}

func (i AstNode) equals(untyped interface{}) bool {
	var j = untyped.(AstNode)
	if (i.function != nil || j.function != nil) && *i.function != *j.function {
		return false
	}
	if !arrayEquals(i.values, j.values) {
		return false
	}
	if !arrayEquals(i.operators, j.operators) {
		return false
	}
	if i.constructingInt != nil && j.constructingInt != nil {
		if !i.constructingInt.equals(*j.constructingInt) {
			return false
		}
	} else if i.constructingInt != nil || j.constructingInt != nil {
		return false
	}
	if i.lastAddedValue != j.lastAddedValue {
		return false
	}
	if i.lastAddedOperator != j.lastAddedOperator {
		return false
	}
	return true
}

// Used by the parser
func (node *AstNode) addToken(token Token) error {
	if token.isInt() {
		if node.lastAddedValue {
			return errors.New("Can't add two values in a row")
		} else if node.constructingInt != nil {
			var added_int = node.constructingInt.addDigit(token.toInt(), false)
			node.constructingInt = &added_int
		} else {
			var new_integer = newInteger().addDigit(token.toInt(), false)
			node.constructingInt = &new_integer
		}
	} else if token.isOperator() {
		if node.lastAddedOperator {
			return errors.New("Can't add two operators in a row")
		} else {
			node.operators = append(node.operators, token)
			node.lastAddedOperator = true
			node.lastAddedValue = false
			if node.constructingInt != nil {
				var new_integer = node.constructingInt.construct()
				node.values = append(node.values, intValue(new_integer))
				node.constructingInt = nil
			}
		}
	} else {
		return errors.New("Only ints and operators implemented so far")
	}
	return nil
}

func (node *AstNode) endTokens() error {
	if node.constructingInt != nil {
		var new_integer = node.constructingInt.construct()
		node.values = append(node.values, intValue(new_integer))
		node.lastAddedValue = true
		node.lastAddedOperator = false
	}
	if node.lastAddedOperator {
		return errors.New("Can't end node on an operator")
	}
	var length_difference = len(node.values) - len(node.operators)
	if length_difference < 0 || length_difference > 1 {
		return errors.New("Incorrect number of values and operators")
	}
	return nil
}
