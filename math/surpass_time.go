package main

import "fmt"

func minNum(samDaily int32, kellyDaily int32, difference int32) int32 {
	// fmt.Println(samDaily, kellyDaily, difference)
	if samDaily >= kellyDaily {
		return -1
	}

	samTotalGap := samDaily + difference
	kellyTotal := kellyDaily

	var surpassDays int32 = 1
	for samTotalGap >= kellyTotal {
		samTotalGap += samDaily
		kellyTotal += kellyDaily
		surpassDays++
	}

	return surpassDays
}

// calculate the number of days it takes for Sam's total productivity to
// surpass Kelly's total productivity based on their daily productivity and an
// initial difference in productivity. The function takes three parameters:

// samDaily: Sam's daily productivity.
// kellyDaily: Kelly's daily productivity.
// difference: The initial difference in productivity, where Sam's total productivity
// is samDaily + difference and Kelly's total productivity is kellyDaily.
func surpassTime() {
	samDaily := 45
	kellyDaily := 46
	difference := 10

	fmt.Println(minNum(int32(samDaily), int32(kellyDaily), int32(difference)))
}
