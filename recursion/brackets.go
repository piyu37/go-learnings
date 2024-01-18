package main

import "fmt"

func printProperBrackets(n int) []string {
	result := make([]string, 0)

	printBrackets("", n, 0, 0, &result)

	return result
}

func printBrackets(str string, n, open, close int, result *[]string) {
	if open == n && close == n {
		*result = append(*result, str)
		return
	}

	if open < n {
		printBrackets(str+"(", n, open+1, close, result)
	}

	if close < open {
		printBrackets(str+")", n, open, close+1, result)
	}
}

// https://www.geeksforgeeks.org/print-all-combinations-of-balanced-parentheses/
func brackets() {
	n := 3
	fmt.Println(printProperBrackets(n))
}
