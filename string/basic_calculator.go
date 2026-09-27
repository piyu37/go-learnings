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

		switch operator {
		case '+', '-':
			result += lastNumber
			if operator == '+' {
				lastNumber = currNumber
			} else {
				lastNumber = -currNumber
			}
		case '*':
			lastNumber *= currNumber
		default:
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

	for i := range length {
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

type frame struct {
	result int  // sum of completed +/- terms at this level
	last   int  // the term currently being built (carries its sign)
	op     byte // operator pending for the next operand
}

// calculateWithBrackets evaluates an expression with + - * / and parentheses.
// e.g. "(3-2*(4+5+2)-3)+(6/2)" -> -19
func calculateWithBrackets(s string) int {
	var stack []frame
	result, last, curr := 0, 0, 0
	op := byte('+')

	// fold one operand(i.e. number) into the current level using the pending operator
	apply := func(operand int) {
		switch op {
		case '+':
			result += last
			last = operand
		case '-':
			result += last
			last = -operand
		case '*':
			last *= operand
		case '/':
			last /= operand
		}
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			curr = curr*10 + int(c-'0')

		case c == ' ':
			// ignore

		case c == '(':
			// save this level, start a fresh one
			stack = append(stack, frame{result, last, op})
			result, last, op, curr = 0, 0, '+', 0

		case c == ')':
			apply(curr)            // fold the last operand of the inner level
			inner := result + last // flush the inner level to a single value

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result, last, op = top.result, top.last, top.op

			// the inner value becomes the pending operand of the outer level;
			// do NOT apply it yet -- let the next operator (or EOF) do that
			curr = inner

		default: // + - * /
			apply(curr)
			op = c
			curr = 0
		}
	}

	apply(curr)

	return result + last
}

// somwwhat similar to https://leetcode.com/problems/basic-calculator/description/
func basicCalculator() {
	exp := "3*2+2"
	fmt.Println(calculate(exp))
	fmt.Println(calculate2(exp))
	exp = "-3+2*2"
	fmt.Println(calculate(exp))
	exp = "(3-2*(4+5+2)-3)+(6/2)"
	fmt.Println(calculateWithBrackets(exp))
}
