package main

import (
	"fmt"
	"math"
)

func getMinNoOfWaysToMakeChangeTopDown(denoms []int, n int) int {
	memoizn := make([]int, n)
	minCoinChanges := findMin(denoms, n, memoizn)
	return minCoinChanges
}

func findMin(denoms []int, n int, memoizn []int) int {
	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	}

	if memoizn[n-1] != 0 {
		return memoizn[n-1]
	}

	minCount := math.MaxInt
	for _, d := range denoms {
		res := findMin(denoms, n-d, memoizn)

		if res >= 0 && res < minCount {
			minCount = res + 1
		}
	}

	if minCount == math.MaxInt {
		memoizn[n-1] = -1
	} else {
		memoizn[n-1] = minCount
	}

	return memoizn[n-1]
}

func getMinNoOfWaysToMakeChangeBottomUp(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for _, c := range coins {
		for a := c; a <= amount; a++ {
			dp[a] = min(dp[a], 1+dp[a-c])
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/15_min_number_of_coins_for_change
// https://leetcode.com/problems/coin-change/description/?envType=study-plan-v2&envId=top-interview-150
func minNoOfWaysToMakeChange() {
	n := 7
	denoms := []int{1, 5, 10}

	fmt.Println(getMinNoOfWaysToMakeChangeTopDown(denoms, n))

	n = 11
	denoms = []int{1, 2, 5}

	fmt.Println(getMinNoOfWaysToMakeChangeTopDown(denoms, n))

	// output: 186,186,186,419,419,419,419,419,83,83,83,83,408,408,408,408,408,408,408,408
	n = 6249
	denoms = []int{186, 419, 83, 408}

	fmt.Println(getMinNoOfWaysToMakeChangeBottomUp(denoms, n))
}
