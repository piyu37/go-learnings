package main

import "fmt"

func areInterweaven(one, two, three string, i, j int) bool {
	k := i + j

	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		if areInterweaven(one, two, three, i+1, j) {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		return areInterweaven(one, two, three, i, j+1)
	}

	return false
}

func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}

	return areInterweaven(one, two, three, 0, 0)
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/30_interweaving_strings
func interweavingStringsMain() {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"

	fmt.Println(InterweavingStrings(one, two, three))
}
