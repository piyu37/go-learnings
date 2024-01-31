package main

import (
	"fmt"
	"math"
)

func getMinNoOfWaysToMakeChange(denoms []int, n int) int {
	return findMin(denoms, n, map[int]int{}, math.MaxInt)
}

func findMin(denoms []int, n int, memoizn map[int]int, minCoinChanges int) int {
	if n == 0 {
		return 0
	}

	for _, d := range denoms {
		change := n - d
		if change < 0 {
			continue
		}

		var minCount int

		if v, ok := memoizn[change]; ok {
			minCount = v
		} else {
			minCount = findMin(denoms, change, memoizn, minCoinChanges)
		}
		if minCount != math.MaxInt && minCount+1 < minCoinChanges {
			minCoinChanges = minCount + 1
		}
	}

	memoizn[n] = minCoinChanges

	return minCoinChanges
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/15_min_number_of_coins_for_change
func minNoOfWaysToMakeChange() {
	n := 18
	denoms := []int{7, 5, 1}

	fmt.Println(getMinNoOfWaysToMakeChange(denoms, n))
}
