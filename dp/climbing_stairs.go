package main

import "fmt"

func climbStairs(n int) int {
	memoization := make([]int, n+1)
	return stairsClimb(n, memoization)
}

func stairsClimb(n int, memoization []int) int {
	if n == 0 {
		memoization[n] = 1
		return 1
	}

	if memoization[n] > 0 {
		return memoization[n]
	}

	count1 := stairsClimb(n-1, memoization)
	count2 := 0
	if n-2 >= 0 {
		count2 = stairsClimb(n-2, memoization)
	}

	memoization[n] = count1 + count2

	return count1 + count2
}

func climbStairs2(n int) int {
	a, b := 0, 1

	for i := n; i > 1; i-- {
		a, b = b, a+b
	}

	return a + b
}

// https://leetcode.com/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-interview-150
func climbingStairsMain() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs2(3))
}
