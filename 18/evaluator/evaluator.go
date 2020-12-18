package evaluator

import (
	"../ast"
)

func Eval(node ast.Node) int64 {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return node.Value
	case *ast.InfixExpression:
		left := Eval(node.Left)
		right := Eval(node.Right)
		return evalInfixExpression(node.Operator, left, right)
	}
	return -1
}

func evalInfixExpression(operator string, left, right int64) int64 {
	switch {
	case operator == "+":
		return left + right
	case operator == "*":
		return left * right
	}
	return -1
}
