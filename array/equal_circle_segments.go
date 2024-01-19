package main

import (
	"fmt"
	"math"
	"slices"
)

// need to omit values in 4 digit
func findMaxAreaSegment(radius []int, nos int) float64 {
	slices.Sort(radius)

	var finalAreaSegment float64

	radMap := make(map[float64]bool)

	areaRadiusArr := make([]float64, len(radius))
	for i := range radius {
		area := 3.14159265359 * float64(radius[i]) * float64(radius[i])
		area = math.Round(area*10000) / 10000
		areaRadiusArr[i] = area
	}

	maxAreaSegment := areaRadiusArr[len(radius)-1]
	maxAreaSegmentIdx := len(radius) - 1
	radMap[areaRadiusArr[maxAreaSegmentIdx]] = true
	areaRadiusArrCopy := make([]float64, len(areaRadiusArr))
	segment := 0
	segmentCut := 1
	for {
		copy(areaRadiusArrCopy, areaRadiusArr)
		for i := len(radius) - 1; i >= 0; i-- {
			for maxAreaSegment <= areaRadiusArrCopy[i] {
				segment++
				areaRadiusArrCopy[i] -= maxAreaSegment
			}
		}

		if segment < nos {
			segmentCut++
			maxAreaSegment = areaRadiusArr[maxAreaSegmentIdx] / float64(segmentCut)
			maxAreaSegment = math.Floor(maxAreaSegment*10000) / 10000
			segment = 0
		} else if segment > nos {
			if maxAreaSegment > finalAreaSegment {
				finalAreaSegment = maxAreaSegment
			}

			maxAreaSegmentIdx--

			for maxAreaSegmentIdx > 0 && radMap[areaRadiusArr[maxAreaSegmentIdx]] {
				maxAreaSegmentIdx--
			}

			if maxAreaSegmentIdx == 0 {
				break
			}

			radMap[areaRadiusArr[maxAreaSegmentIdx]] = true
			maxAreaSegment = areaRadiusArr[maxAreaSegmentIdx]
			segment = 0
			segmentCut = 1
		} else {
			break
		}
	}

	return maxAreaSegment
}

// https://extra.codemotion.com/the-philips-coding-challenge/
func equal_circle_segments() {
	r := []int{1, 1, 1, 2, 2, 3}
	r1 := []int{2, 2, 2, 2, 2, 2, 3}

	noOfSegment := 6

	fmt.Println(findMaxAreaSegment(r, noOfSegment))
	fmt.Println(findMaxAreaSegment(r1, noOfSegment))

}
