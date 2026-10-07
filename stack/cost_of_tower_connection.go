package main

import "fmt"

func towerConnection(N, K int, H []int) (int, int) {
	stack := make([]int, 0) // stores idx
	degrees := make([]int, N)
	maxDegree := 0
	maxDegreeIdx := N - 1

	towerConnectionCost := 0

	for i := N - 1; i >= 0; i-- {
		for len(stack) > 0 && H[i] >= H[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			towerConnectionCost += K
		} else {
			towerConnectionCost += H[i] + H[stack[len(stack)-1]]
			degrees[stack[len(stack)-1]] += 1

			if degrees[stack[len(stack)-1]] >= degrees[maxDegreeIdx] {
				maxDegreeIdx = stack[len(stack)-1]
				maxDegree = degrees[stack[len(stack)-1]]
			}
		}

		stack = append(stack, i)
	}

	return towerConnectionCost, maxDegree
}

// Cost of Tower Connection (Golang Coding)Problem Statement:
// In a city, there are N towers in a row, and each tower has some height. A tower can be connected to the closest tower
// to its right that has a height greater than its height. The cost of this connection is the sum of the two heights.
// If the tower cannot be connected to any other tower, an additional cost K is added. The degree of the whole connection
// is defined as the maximum number of connections to a particular tower. Print the total cost of the tower connection and
// the degree.
// Constraints:1 <= N <= 1000000   0 <= K <= 1000000   1 <= Height <= 1000000
// Sample:Input: 4 10 (denotes N and K) \n 4 7 3 8 (denotes height of towers)
// Output: 47 2   Explanation: Tower 1 (4) connects to Tower 2 (7). Towers 2 (7) and 3 (3) connect to Tower 4 (8).
// Tower 4 does not connect to anyone. Cost = (4+7) + (7+8) + (3+8) + 10 = 47. Tower 4 receives 2 connections, making the degree 2.

// Similar LeetCode Problem: Finding the "closest tower to its right that is greater" is a textbook use case for a Monotonic Stack.
// See LeetCode 496: Next Greater Element I and LeetCode 503: Next Greater Element II.
// Ericsson Hackerrank test
func costOfTowerConnectionMain() {
	N, K := 4, 10
	H := []int{4, 7, 3, 8}

	fmt.Println(towerConnection(N, K, H))
}
