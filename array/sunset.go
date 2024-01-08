package main

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
