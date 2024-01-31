package main

func MaxSubsetSumNoAdjacentWithoutModifyingInput(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}

	first := max(array[0], array[1]) // array[i-1]
	second := array[0]               // array[i-2]
	for i := 2; i < len(array); i++ {
		// ex.  [7 10 12 14]
		current := second + array[i]    // 10 + 14
		previous := max(first, current) // max(array[i-1], array[i-2] + array[i])
		second = first                  // 10
		first = previous                // 19
	}
	return first
}
