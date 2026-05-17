package main

import (
	"fmt"
	"math"
)

func maxScoreSightseeingPair(values []int) int {
	maxScore := math.MinInt32

	leftMax := values[0]

	for i := 1; i < len(values); i++ {
		leftMax = max(leftMax, values[i-1]+i-1)

		if leftMax+(values[i]-i) > maxScore {
			maxScore = leftMax + (values[i] - i)
		}
	}

	return maxScore
}

// https://leetcode.com/problems/best-sightseeing-pair/description/
func bestSightSeeingPair() {
	values := []int{8, 1, 5, 2, 6}

	fmt.Println(maxScoreSightseeingPair(values))
}
