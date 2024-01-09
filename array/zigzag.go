package main

import "fmt"

func convert(s string, numRows int) string {
	result := ""

	if numRows == 1 {
		return s
	}

	sLen := len(s)
	charsInSec := (numRows - 1) * 2

	for r := 0; r < numRows; r++ {
		i := r

		for i < sLen {
			result += string(s[i])

			if r != 0 && r != numRows-1 {
				charInBetween := charsInSec - (2 * r)
				secondIdx := i + charInBetween

				if secondIdx < sLen {
					result += string(s[secondIdx])
				}
			}

			i += charsInSec
		}
	}

	return result
}

// https://leetcode.com/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150
func zigzag() {
	s := "PAYPALISHIRING"
	row := 4

	fmt.Println(convert(s, row))
}
