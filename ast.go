package main

import (
	"errors"
	"strconv"
	"strings"
)

type AstNode struct {
	function *Token

	values    []Value
	operators []Token

	constructingInt   *Integer
	constructingNode  *AstNode
	closedBracketType *Token // use nil for an implied bracket

	lastAddedValue    bool
	lastAddedOperator bool
}

func newAstNode() AstNode {
	return AstNode{
		function:          nil,
		values:            make([]Value, 0),
		operators:         make([]Token, 0),
		constructingInt:   nil,
		constructingNode:  nil,
		lastAddedValue:    false,
		lastAddedOperator: false,
	}
}

func constructAstNode(fn Token, vs []Value, os []Token, in Integer, an AstNode, lastV bool, lastO bool) AstNode {
	return AstNode{
		function:          &fn,
		values:            vs,
		operators:         os,
		constructingInt:   &in,
		constructingNode:  &an,
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

func (node AstNode) toDebugString(append_string string) string {
	var return_string strings.Builder
	return_string.WriteString("AstNode {")
	if node.function != nil {
		return_string.WriteString("\n" + append_string + "function: " + node.function.toString())
	}
	return_string.WriteString("\n" + append_string + "values: [")
	var first = true
	for _, value := range node.values {
		if !first {
			return_string.WriteString(", ")
		}
		first = false
		return_string.WriteString(value.toString(false))
	}
	return_string.WriteString("]")
	return_string.WriteString("\n" + append_string + "operators: [")
	first = true
	for _, operator := range node.operators {
		if !first {
			return_string.WriteString(", ")
		}
		first = false
		return_string.WriteString(operator.toString())
	}
	return_string.WriteString("]")
	if node.constructingInt != nil {
		return_string.WriteString("\n" + append_string + "constructingInt: " + node.constructingInt.toString())
	}
	if node.constructingNode != nil {
		return_string.WriteString("\n" + append_string + "constructingNode: " +
			node.constructingNode.toDebugString(append_string+"  "))
	}
	return_string.WriteString("\n" + append_string + "lastAddedValue: " + strconv.FormatBool(node.lastAddedValue))
	return_string.WriteString("\n" + append_string + "lastAddedOperator: " + strconv.FormatBool(node.lastAddedOperator))
	return_string.WriteString("\n" + append_string + "}")
	return return_string.String()
}

func (i AstNode) equals(untyped interface{}) bool {
	var j = untyped.(AstNode)
	if (i.function != nil) != (j.function != nil) {
		return false
	}
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
	if i.constructingNode != nil && j.constructingNode != nil {
		if !i.constructingNode.equals(*j.constructingNode) {
			return false
		}
	} else if i.constructingNode != nil || j.constructingNode != nil {
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

// Recursive token adder used by the parser
func (node *AstNode) addToken(token Token) error {
	if node.constructingNode != nil {
		if token.isCloseParens() && node.constructingNode.constructingNode == nil {
			if *node.closedBracketType != token {
				return errors.New("didn't close bracket on same type as opening it")
			}
			var error = node.constructingNode.endTokens()
			if error != nil {
				return error
			}
			node.finishConstructingNode()
			return nil
		} else if token.isOperator() && node.constructingNode.function != nil &&
			(len(node.constructingNode.values) > 0 || node.constructingNode.constructingInt != nil) {
			if node.closedBracketType != nil {
				return errors.New("implied closing of explicit bracket type")
			}
			var error = node.constructingNode.endTokens()
			if error != nil {
				return error
			}
			node.finishConstructingNode()
		} else {
			return node.constructingNode.addToken(token)
		}
	}
	if token.isInt() {
		if node.lastAddedValue {
			return errors.New("can't add two values in a row")
		} else if node.constructingInt != nil {
			var added_int = node.constructingInt.addDigit(token.toInt(), false)
			node.constructingInt = &added_int
		} else {
			var new_integer = newInteger().addDigit(token.toInt(), false)
			node.constructingInt = &new_integer
		}
	} else if token.isOperator() {
		if node.lastAddedOperator && node.constructingInt == nil {
			return errors.New("can't add two operators in a row")
		} else if node.function != nil && (len(node.values) > 0 || node.constructingInt != nil) {
			return errors.New("can't add operator to node with function")
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
	} else if token.isOpenParens() {
		if node.lastAddedValue {
			// TODO: auto-add multiplication
			return errors.New("can't add two values in a row")
		} else {
			var new_node = newAstNode()
			node.constructingNode = &new_node
			node.closedBracketType = takePtr(Token(token + 1))
		}
	} else if token.isCloseParens() {
		return errors.New("can't end parentheses if not constructing a node")
	} else if token.isFunction() {
		if node.lastAddedValue {
			// TODO: auto-add multiplication
			return errors.New("can't add two values in a row")
		} else {
			var new_node = newAstNode()
			new_node.function = &token
			node.constructingNode = &new_node
			node.closedBracketType = nil
		}
	} else {
		return errors.New("unrecognized token type")
	}
	return nil
}

func (node *AstNode) finishConstructingNode() {
	if node.constructingNode == nil {
		panic("can't finish constructing node when not constructing node")
	}
	node.values = append(node.values, nodeValue(*node.constructingNode))
	node.constructingNode = nil
	node.lastAddedValue = true
	node.lastAddedOperator = false
}

func (node *AstNode) endTokens() error {
	if node.constructingNode != nil {
		var recursive_error = node.constructingNode.endTokens()
		if recursive_error != nil {
			return recursive_error
		}
		node.finishConstructingNode()
	}
	if node.constructingInt != nil {
		var new_integer = node.constructingInt.construct()
		node.values = append(node.values, intValue(new_integer))
		node.lastAddedValue = true
		node.lastAddedOperator = false
		node.constructingInt = nil
	}
	if node.lastAddedOperator {
		return errors.New("can't end node on an operator")
	}
	var length_difference = len(node.values) - len(node.operators)
	if length_difference < 0 || length_difference > 1 {
		return errors.New("incorrect number of values and operators")
	}
	return nil
}
