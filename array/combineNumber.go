package main

import (
	"strconv"
)

func Solution(list []int) string {
	result := ""

	listLen := len(list)
	count := 1
	i := 0
	for i = 0; i < listLen-1; i++ {
		if list[i] == list[i+1]-1 {
			count = count + 1
			continue
		}

		if count > 2 {
			val1 := list[i]
			val2 := list[i] - count + 1
			result = result + strconv.Itoa(val2) + "-" + strconv.Itoa(val1) + ","
		} else if count == 2 {
			result = result + strconv.Itoa(list[i-1]) + "," + strconv.Itoa(list[i]) + ","
		} else {
			s := strconv.Itoa(list[i])
			result = result + s + ","
		}

		count = 1
	}

	if count > 2 {
		val1 := list[i]
		val2 := list[i] - count + 1
		result = result + strconv.Itoa(val2) + "-" + strconv.Itoa(val1) + ","
	} else if count == 2 {
		result = result + strconv.Itoa(list[i-1]) + "," + strconv.Itoa(list[i]) + ","
	} else {
		s := strconv.Itoa(list[i])
		result = result + s + ","
	}

	b := len(result) - 1
	result = result[:b]

	return result
}
