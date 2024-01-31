package main

import "fmt"

func findNoOfWaysToMakeChange(denoms []int, n int) int {
	prvState := make([]int, n+1)
	prvState[0] = 1

	for i := 0; i < len(denoms); i++ {
		resultArr := make([]int, n+1)
		for j := 0; j <= n; j++ {
			var withCoinIncluded int
			if j-denoms[i] >= 0 {
				withCoinIncluded = resultArr[j-denoms[i]]
			}

			withoutCoinIncluded := prvState[j]

			resultArr[j] = withCoinIncluded + withoutCoinIncluded
		}

		prvState = resultArr
	}

	return prvState[n]
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/14_number_of_ways_to_make_change
func noOfWaysToMakeChange() {
	n := 5
	denoms := []int{1, 2, 5}

	fmt.Println(findNoOfWaysToMakeChange(denoms, n))
}
