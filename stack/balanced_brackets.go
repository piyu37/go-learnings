package main

import "fmt"

func BalancedBrackets(s string) bool {
	stack := make([]string, 0)
	stackLen := 0

	for _, r := range s {
		str := string(r)

		if str == "(" || str == "[" || str == "{" {
			stack = append(stack, str)
			stackLen++
		} else if stackLen > 0 {
			if str == ")" {
				if stack[stackLen-1] == "(" {
					stack = stack[:stackLen-1]
					stackLen--
				} else {
					return false
				}
			} else if str == "]" {
				if stack[stackLen-1] == "[" {
					stack = stack[:stackLen-1]
					stackLen--
				} else {
					return false
				}
			} else {
				if stack[stackLen-1] == "{" {
					stack = stack[:stackLen-1]
					stackLen--
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}

	return stackLen == 0
}

func balancedBracketsMain() {
	val := "(a)"
	fmt.Println(BalancedBrackets(val))
}
