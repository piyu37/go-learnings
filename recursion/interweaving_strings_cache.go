package main

import "fmt"

func areInterweavenCache(one, two, three string, i, j int, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	k := i + j

	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterweavenCache(one, two, three, i+1, j, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterweavenCache(one, two, three, i, j+1, cache)
		cache[i][j] = &result

		return result
	}

	result := false
	cache[i][j] = &result

	return result
}

func InterweavingStringsCache(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}

	cache := make([][]*bool, len(one)+1)

	for i := 0; i < len(one)+1; i++ {
		cache[i] = make([]*bool, len(two)+1)
	}

	return areInterweavenCache(one, two, three, 0, 0, cache)
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/30_interweaving_strings
func interweavingStringsCacheMain() {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"

	fmt.Println(InterweavingStringsCache(one, two, three))
}
