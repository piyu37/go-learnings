package main

import (
	"fmt"
	"math"
)

func mainRect(points [][]int) int {
	result := 0

	maxArea := math.MaxInt

	yCoordinates := make(map[int][]int)

	for _, point := range points {
		yCoordinates[point[0]] = append(yCoordinates[point[0]], point[1])
	}

	for i, point := range points {
		x1 := point[0]
		y1 := point[1]

		for j := i + 1; j < len(points); j++ {
			x2 := points[j][0]
			y2 := points[j][1]

			if x1 != x2 && y1 != y2 {
				if contains(yCoordinates[x1], y2) && contains(yCoordinates[x2], y1) {
					area := abs(x2-x1) * abs(y2-y1)
					if area < maxArea {
						result = area
						maxArea = area
					}
				}
			}
		}
	}

	return result
}

func contains(arr []int, v int) bool {
	for i := range arr {
		if arr[i] == v {
			return true
		}
	}

	return false
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/19_rectangle_mania
// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/04_minimum_area_rectangle
func minRectangle() {
	points := [][]int{
		{36219, 4673},
		{26311, 36047},
		{26311, 4673},
		{36219, 16024},
		{17010, 16024},
		{26311, 6287},
		{22367, 6287}, {17010, 36047}, {17010, 6287},
		{22367, 16024}, {36219, 6287}, {22367, 4673}, {17010, 4673}, {36219, 36047},
	}

	points2 := [][]int{
		{1, 5},
		{5, 1},
		{4, 2},
		{2, 4},
		{2, 2},
		{1, 2},
		{4, 5},
		{2, 5},
		{-1, -2},
	}

	fmt.Println(mainRect(points))
	fmt.Println(mainRect(points2))
}
