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

type floatInt interface {
	~float32 | ~float64 | ~int | ~int16 | ~int32 | ~int64
}

// generic method
func abs[fi floatInt](v1 fi) fi {
	if v1 >= 0 {
		return v1
	}

	return -v1
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

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/05_line_through_points
func lineThroughPointsMain() {
	arr := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{0, 4},
		{-2, 6},
		{4, 0},
		{2, 1},
	}

	fmt.Println(LineThroughPoints(arr))
}
