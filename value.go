package main

import (
	"strings"
)

type ValueType int8

const (
	ERROR_ValueType ValueType = iota
	INTEGER
	AST_NODE
)

func (value_type ValueType) toString() string {
	switch value_type {
	case INTEGER:
		return "Integer"
	case AST_NODE:
		return "AstNode"
	default:
		return ""
	}
}

type Value struct {
	value_type ValueType

	integer *Integer
	// rational number
	// irrational number
	ast_node *AstNode
}

func intValue(integer Integer) Value {
	return Value{
		value_type: INTEGER,
		integer:    &integer,
		ast_node:   nil,
	}
}

func nodeValue(node AstNode) Value {
	return Value{
		value_type: AST_NODE,
		integer:    nil,
		ast_node:   &node,
	}
}

func (i Value) equals(untyped interface{}) bool {
	var j = untyped.(Value)
	if i.value_type != j.value_type {
		return false
	}
	switch i.value_type {
	case ERROR_ValueType:
		return true
	case INTEGER:
		return i.integer.equals(*j.integer)
	case AST_NODE:
		return i.ast_node.equals(*j.ast_node)
	default:
		panic("Unknown value type")
	}
}

func (value Value) toString(shallow bool) string {
	var return_string strings.Builder
	return_string.WriteString("Value {")
	return_string.WriteString("\n  value_type: " + value.value_type.toString())
	switch value.value_type {
	case INTEGER:
		return_string.WriteString("\n  integer: " + value.integer.toString())
	case AST_NODE:
		if shallow {
			return_string.WriteString("\n  ast_node: ...")
		} else {
			return_string.WriteString("\n  ast_node: " + value.ast_node.toString(shallow))
		}
	default:
		panic("Unknown value type")
	}
	return_string.WriteString("\n}")
	return return_string.String()
}

func (value Value) toResultString() string {
	switch value.value_type {
	case INTEGER:
		return value.integer.toString()
	default:
		panic("Unknown value type")
	}
}
