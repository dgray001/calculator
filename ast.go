package main

import (
	"errors"
	"strings"
)

type AstNode struct {
	left     Integer // eventually convert to a generic number type
	operator Token   // guaranteed to be an operator token at construction
	right    Integer

	constructingLeftInt  bool // true if currently adding int tokens to left
	constructingRightInt bool // true if currently adding int tokens to right

	hasLeft     bool
	hasOperator bool // if true then node must have right
	hasRight    bool // if true then node must have operator
}

func newAstNode() AstNode {
	return AstNode{
		hasLeft:     false,
		hasOperator: false,
		hasRight:    false,
	}
}

func (node AstNode) toString() string {
	var return_string strings.Builder
	return_string.WriteString("AstNode {")
	if node.hasLeft {
		return_string.WriteString("\n  left: " + node.left.toString())
	}
	if node.hasOperator {
		return_string.WriteString("\n  operator: " + node.operator.toString())
	}
	if node.hasRight {
		return_string.WriteString("\n  right: " + node.right.toString())
	}
	return_string.WriteString("\n}")
	return return_string.String()
}

func (i AstNode) equals(j AstNode) bool {
	if !i.left.equals(j.left) {
		return false
	}
	if i.operator != j.operator {
		return false
	}
	if !i.right.equals(j.right) {
		return false
	}
	if i.constructingLeftInt != j.constructingLeftInt {
		return false
	}
	if i.constructingRightInt != j.constructingRightInt {
		return false
	}
	if i.hasLeft != j.hasLeft {
		return false
	}
	if i.hasOperator != j.hasOperator {
		return false
	}
	if i.hasRight != j.hasRight {
		return false
	}
	return true
}

// Used by the parser
func (node *AstNode) addToken(token Token) error {
	if token.isInt() {
		if node.hasLeft {
			if node.constructingLeftInt {
				return errors.New("Can't be constructing left int after node already has left int")
			} else if node.hasRight {
				return errors.New("Can't add int token to node that already has left and right ints")
			} else if !node.hasOperator {
				return errors.New("Can't add right int to node before operator")
			} else if node.constructingRightInt {
				node.right = node.right.addDigit(token.toInt(), false)
			} else {
				node.constructingRightInt = true
				node.right = newInteger().addDigit(token.toInt(), false)
			}
		} else if node.hasOperator {
			if node.constructingLeftInt {
				return errors.New("Can't be constructing left int after node already has operator")
			} else if node.hasRight {
				return errors.New("Can't add right int token to node with right int already")
			} else if node.constructingRightInt {
				node.right = node.right.addDigit(token.toInt(), false)
			} else {
				node.constructingRightInt = true
				node.right = newInteger().addDigit(token.toInt(), false)
			}
		} else {
			if node.constructingRightInt {
				return errors.New("Can't be constructing right int without operator or left int")
			} else if node.constructingLeftInt {
				node.left = node.left.addDigit(token.toInt(), false)
			} else {
				node.constructingLeftInt = true
				node.left = newInteger().addDigit(token.toInt(), false)
			}
		}
	} else {
		return errors.New("Only ints implemented so far")
	}
	return nil
}

func (node *AstNode) endTokens() error {
	if node.constructingLeftInt && node.constructingRightInt {
		return errors.New("Can't be constructing left and right ints")
	} else if node.constructingLeftInt {
		node.constructingLeftInt = false
		node.left = node.left.construct()
		node.hasLeft = true
	} else if node.constructingRightInt {
		node.constructingRightInt = false
		node.right = node.right.construct()
		node.hasRight = true
	}
	return nil
}
