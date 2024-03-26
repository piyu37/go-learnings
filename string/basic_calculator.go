package main

import (
	"fmt"
	"unicode"
)

func calculate(s string) int {
	result := 0
	lastNumber := 0
	currNumber := 0
	operator := '+'
	operatorMap := make(map[rune]bool)
	operatorMap['+'] = true
	operatorMap['-'] = true
	operatorMap['*'] = true
	operatorMap['/'] = true
	for i, r := range s {
		if !operatorMap[r] {
			currNumber = currNumber*10 + int(r-'0')
			if i < len(s)-1 {
				continue
			}
		}

		if operator == '+' || operator == '-' {
			result += lastNumber
			if operator == '+' {
				lastNumber = currNumber
			} else {
				lastNumber = -currNumber
			}
		} else if operator == '*' {
			lastNumber *= currNumber
		} else {
			lastNumber /= currNumber
		}

		operator = r
		currNumber = 0
	}

	result += lastNumber

	return result
}

// calculate evaluates a basic arithmetic expression given as a string
func calculate2(s string) int {
	if s == "" {
		return 0
	}
	length := len(s)
	currentNumber, lastNumber, result := 0, 0, 0
	operation := '+'

	for i := 0; i < length; i++ {
		currentChar := rune(s[i])
		if unicode.IsDigit(currentChar) {
			currentNumber = currentNumber*10 + int(currentChar-'0')
		}
		if !unicode.IsDigit(currentChar) && !unicode.IsSpace(currentChar) || i == length-1 {
			switch operation {
			case '+', '-':
				result += lastNumber
				if operation == '+' {
					lastNumber = currentNumber
				} else {
					lastNumber = -currentNumber
				}
			case '*':
				lastNumber *= currentNumber
			case '/':
				lastNumber /= currentNumber
			}
			operation = currentChar
			currentNumber = 0
		}
	}
	result += lastNumber
	return result
}

func basicCalculator() {
	exp := "3*2+2"
	fmt.Println(calculate(exp))
	fmt.Println(calculate2(exp))
}
