package main

import "fmt"

func AmbiguousMeasurements(measuringCups [][]int, low int, high int) bool {
	// Write your code here.
	memoization := make(map[string]bool)

	return isValidForMeasurement(measuringCups, low, high, 0, 0, len(measuringCups)-1, memoization)
}

func isValidForMeasurement(measuringCups [][]int, low, high, countLow, countHigh,
	index int, memoization map[string]bool) bool {
	memoizeKey := fmt.Sprintf("%d:%d", countLow, countHigh)

	if v, exist := memoization[memoizeKey]; exist {
		return v
	}

	if countLow >= low {
		if countLow > high {
			return false
		}

		if countHigh <= high {
			return true
		} else {
			return false
		}
	}

	canMeasure := false

	for i := index; i >= 0; i-- {
		lowTemp, highTemp := measuringCups[i][0], measuringCups[i][1]

		countLow += lowTemp
		countHigh += highTemp

		canMeasure = isValidForMeasurement(measuringCups, low, high, countLow, countHigh, i, memoization)

		if canMeasure {
			break
		} else {
			countLow -= lowTemp
			countHigh -= highTemp
		}
	}

	memoization[memoizeKey] = canMeasure

	return canMeasure
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/33_ambiguous_measurements
func ambiguousMeasurementsMain() {
	mc := [][]int{
		{1, 3},
		{2, 4},
		{5, 6},
	}
	low := 100
	high := 101

	fmt.Println(AmbiguousMeasurements(mc, low, high))
}
