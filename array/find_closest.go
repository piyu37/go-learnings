package main

import "fmt"

// method which will take 2 parameters []float and float value

func findClosestValue(values []float32, value float32) float32 {
	if len(values) == 0 {
		return -1
	}

	if len(values) == 1 {
		return values[0]
	}

	closestValue := abs(value - values[0])
	closestArrayValue := values[0]

	for i := 1; i < len(values); i++ {
		tempClosestValue := abs(value - values[i])

		if tempClosestValue < closestValue {
			closestValue = tempClosestValue
			closestArrayValue = values[i]
		}
	}

	return closestArrayValue
}

func findClosestValueMain() {
	values := []float32{1.1, 1.72, 1.75, 1.68, 1.23, -3.7, -1.5}
	value := 1.73

	fmt.Println(findClosestValue(values, float32(value)))
}
