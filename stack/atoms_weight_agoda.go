package main

import (
	"fmt"
	"unicode"
)

func atomsWeightAgoda() {
	formula := "CH4"
	fmt.Println(calculateWeight(formula))

	formula = "((CH4))"
	fmt.Println(calculateWeightWithBrackets(formula))

	formula = "H(CH4)2"
	fmt.Println(calculateWeightWithBrackets(formula))
}

// Description

// A molecule is composed of atoms. Each atom is represented by a single uppercase letter.
// We only consider C, H, and O for this problem. The weight of each element can be found below:

// C = 12
// H = 1
// O = 8

// We are interested in calculating the mass of the molecule. It does not matter if the specified
// formula is not a valid molecule, we just want to get the total weight.

// You need to write a function calculateWeight which takes in a single parameter: a string
// called formula. We can calculate the total weight by adding all the weights of each atom,
// multiplied by the appropriate multiplier. A number in the input string represents a multiplier
// and it is applied to the element preceding it.

// We also guarantee that all the numbers are single-digit and larger than 0.

// Examples

// CH4: The output should be 16 as it is 12 (C is 12) and 4 times H (which is 1), so it is 12 + 4 = 16.

// Function Description

// Complete the function calculateWeight in the editor below.

// Function Parameters: string formula, the chemical formula
func calculateWeight(formula string) int {
	atomMap := map[byte]int{
		'C': 12,
		'H': 1,
		'O': 8,
	}

	result := 0
	for i := 0; i < len(formula); i++ {
		if i+1 < len(formula) && unicode.IsDigit(rune(formula[i+1])) {
			result += atomMap[formula[i]] * int(formula[i+1]-'0')
			i++
		} else {
			result += atomMap[formula[i]]
		}
	}

	return result
}

// Description

// (Same setup as Part A: elements C = 12, H = 1, O = 8.)

// As we know, chemical formulas can have groups of elements having the same multiplier,
// with the group being enclosed in a pair of brackets.

// You need to write a function calculateWeight which takes in a single parameter:
// a string called formula. Now, if a number is after a close bracket ), then it is applied to all inside the preceding
// bracket group. You can see the examples below for more explanation.

// We also guarantee that all the numbers are single-digit and larger than 0. Note that a formula with an opening bracket
// at the start and a closing bracket at the end is a valid input, e.g. ((CH4)) will still be valid with a weight of 16.

// Examples

// CH4: The output should be 16 as it is 12 (C is 12) and 4 times H (which is 1), so it is 12 + 4 = 16.

// H(CH4)2: The output should be 33 as it is H, which is 1, plus (CH4)2. In the latter part, the multiplier "2" is applied to
// everything inside the preceding brackets. Since CH4 is 16, we multiply this by 2, so we get 32. As a result, we get 1 + 32,
// which is 33.

// We guarantee that all the inputs are valid and there are no unmatched brackets, meaning that all opening brackets have a
// corresponding closing bracket. All elements in the input are guaranteed to be either C, H, or O.
// All numbers in the input are single-digit numbers.

// Function Description

// Complete the function calculateWeight in the editor below.

// Function Parameters: string formula, the chemical formula
func calculateWeightWithBrackets(formula string) int {
	atomMap := map[byte]int{
		'C': 12,
		'H': 1,
		'O': 8,
	}

	stack := make([]int, 0)
	stack = append(stack, 0)

	for i := 0; i < len(formula); i++ {
		if formula[i] == '(' {
			stack = append(stack, 0)
			continue
		}

		if formula[i] == ')' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if i+1 < len(formula) && unicode.IsDigit(rune(formula[i+1])) {
				top *= int(formula[i+1] - '0')
				i++
			}

			stack[len(stack)-1] += top
			continue
		}

		if i+1 < len(formula) && unicode.IsDigit(rune(formula[i+1])) {
			stack[len(stack)-1] += atomMap[formula[i]] * int(formula[i+1]-'0')
			i++
		} else {
			stack[len(stack)-1] += atomMap[formula[i]]
		}
	}

	return stack[0]
}
