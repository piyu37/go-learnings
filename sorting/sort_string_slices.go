package main

import (
	"fmt"
)

func sortSliceOfString() {
	// Create a slice of strings
	strings := []string{"cab", "bb", "bad"}

	// Define a less function
	less := func(i, j int) bool {
		return strings[i] < strings[j]
	}

	// Sort the slice using the less function
	for i := 0; i < len(strings)-1; i++ {
		for j := i + 1; j < len(strings); j++ {
			if less(j, i) {
				strings[i], strings[j] = strings[j], strings[i]
			}
		}
	}

	// Print the sorted slice
	fmt.Println(strings)
}

func sortString() {
	// Create a string
	str := "hello world"

	// Convert the string to a slice of runes
	runes := []rune(str)

	// Sort the slice of runes
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}

	// Convert the slice of runes back to a string
	sortedStr := string(runes)

	// Print the sorted string
	fmt.Println(sortedStr)
}
