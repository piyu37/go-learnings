package main

import (
	"math"
)

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas.
// The guards have gone and will come back in h hours.
// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile
// of bananas and eats k bananas from that pile. If the pile has less than k bananas, she
// eats all of them instead and will not eat any more bananas during this hour.
// Koko likes to eat slowly but still wants to finish eating all the bananas before the
// guards return.
// Return the minimum integer k such that she can eat all the bananas within h hours.

// Input: piles = [3,6,7,11], h = 8
// o/p = 4

// Input: piles = [30,11,23,4,20], h = 5
// = 30
// 4 11 20 23 30

// Input: piles = [30,11,23,4,20], h = 6
// o/p = 23

// Input: piles = [30,11,23,4,20], h = 7
// o/p = 20

// Input: piles = [30,11,23,4,20], h = 8
// o/p = 15

// Input: piles = [30,11,23,4,20], h = 9
// o/p = 12

// Input: piles = [30,11,23,4,20], h = 10
// o/p = 11

func minEatingSpeed(piles []int, h int) int {
	maxInPile := findMax(piles)

	left := 1
	right := maxInPile

	for left < right {
		mid := left + (right-left)/2
		totalHours := 0
		for _, pile := range piles {
			totalHours += int(math.Ceil(float64(pile) / float64(mid)))
		}

		if totalHours <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

func findMax(values []int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}
