package main

import "fmt"

func SunsetViews(buildings []int, direction string) []int {
	result := make([]int, 0)
	buildingsLen := len(buildings)
	top := -1
	if direction == "EAST" {
		for i := buildingsLen - 1; i >= 0; i-- {
			if buildings[i] > top {
				result = append(result, i)
				top = buildings[i]
			}
		}

		left := 0
		right := len(result) - 1
		for left < right {
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}

		return result
	}

	for i := 0; i < buildingsLen; i++ {
		if buildings[i] > top {
			result = append(result, i)
			top = buildings[i]
		}
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/47_sunset_views
func sunset() {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "WEST"
	fmt.Println(SunsetViews(buildings, direction))
}
