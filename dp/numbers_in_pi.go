package main

import (
	"fmt"
	"math"
)

func findSpacesInPi(pi string, numbers []string) int {
	hash := make(map[string]bool)

	for i := range numbers {
		hash[numbers[i]] = true
	}

	noOfSpace := findSpaces(0, pi, hash, map[int]int{})

	return noOfSpace
}

func findSpaces(idx int, pi string, hash map[string]bool, memo map[int]int) int {
	if len(pi) == idx {
		return -1
	} else if v, ok := memo[idx]; ok {
		return v
	}

	str := ""
	noOfSpace := math.MaxInt8
	for i := idx; i < len(pi); i++ {
		str += string(pi[i])

		if hash[str] {
			count := findSpaces(i+1, pi, hash, memo)
			if count < noOfSpace {
				noOfSpace = count + 1
			}

			memo[idx] = noOfSpace
		}
	}

	return noOfSpace
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/16_numbers_in_pi
func numbersInPi() {
	pi := "3141592653589793238462643383279"
	numbers := []string{"314159265358979323846", "3141592653589793238462", "2", "6433", "8", "3279", "83279", "314159265", "35897932384626433832", "79"}
	fmt.Println(findSpacesInPi(pi, numbers))
}
