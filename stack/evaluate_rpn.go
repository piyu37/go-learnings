package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}

	for i := range tokens {
		token := tokens[i]

		if !isOperator(token) {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			top2 := stack[len(stack)-2]
			top1 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			output := evaluate(top1, top2, token)
			stack = append(stack, output)
		}
	}

	return stack[0]
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

func evaluate(val1, val2 int, operator string) int {
	switch operator {
	case "+":
		return val2 + val1
	case "-":
		return val2 - val1
	case "*":
		return val2 * val1
	case "/":
		return val2 / val1
	}

	return 0
}

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/?envType=study-plan-v2&envId=top-interview-150
func evaluateRPN() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
