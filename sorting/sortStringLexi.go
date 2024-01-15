package main

import "fmt"

func smallestString(s string) string {
	result := ""

	b := 0

	for _, r := range s {
		if r == 'b' {
			b++
		}
	}

	used := false

	for _, r := range s {
		if r == 'c' && !used {
			used = true

			for j := 0; j < b; j++ {
				result += "b"
			}
		}

		if r != 'b' {
			result += string(r)
		}
	}

	if !used {
		for j := 0; j < b; j++ {
			result += "b"
		}
	}

	return result
}

// the function aims to produce the smallest lexicographically possible string by replacing the
// first 'c' with the appropriate number of 'b' characters.
func sortStringLexi() {
	str := "baacba"
	fmt.Println(smallestString(str))
}
