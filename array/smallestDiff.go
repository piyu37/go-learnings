package main

import (
	"fmt"
)

func LineThroughPoints(points [][]int) int {
	maxPointsOnLine := 1

	for i, p1 := range points {
		slopeMap := make(map[string]int)

		for j := i + 1; j < len(points); j++ {
			p2 := points[j]

			y := p1[1] - p2[1]
			x := p1[0] - p2[0]
			gcd := getGCD(abs(x), abs(y))
			y /= gcd
			x /= gcd
			if x < 0 {
				y *= -1
				x *= -1
			}

			key := fmt.Sprintf("%d/%d", y, x)
			if _, ok := slopeMap[key]; ok {
				slopeMap[key] += 1
			} else {
				slopeMap[key] = 2
			}

			if slopeMap[key] > maxPointsOnLine {
				maxPointsOnLine = slopeMap[key]
			}
		}
	}

	return maxPointsOnLine
}

func getGCD(x, y int) int {
	for {
		if x == 0 {
			return y
		}

		if y == 0 {
			return x
		}

		x, y = y, x%y
	}
}
