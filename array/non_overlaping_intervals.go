package main

import (
	"fmt"
	"math"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	removedCount := 0
	currIdx := 1

	lastKeptEnd := intervals[currIdx-1][1]
	for currIdx < len(intervals) {
		currStart, currEnd := intervals[currIdx][0], intervals[currIdx][1]

		if currStart < lastKeptEnd {
			if currEnd < lastKeptEnd {
				lastKeptEnd = currEnd
			}

			removedCount++
		} else {
			lastKeptEnd = currEnd
		}

		currIdx++
	}

	return removedCount
}

// Greedy by end time: sort by end, keep every interval that starts at or after
// the last kept interval's end, and remove the rest.
func eraseOverlapIntervalsSortByEnd(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	removedCount := 0
	lastKeptEnd := math.MinInt

	for _, interval := range intervals {
		currStart, currEnd := interval[0], interval[1]

		if currStart >= lastKeptEnd {
			// no overlap: keep it, its end becomes the new boundary
			lastKeptEnd = currEnd
		} else {
			// overlaps the last kept interval: remove it
			removedCount++
		}
	}

	return removedCount
}

// https://leetcode.com/problems/non-overlapping-intervals/
func nonOverlapingIntervals() {
	intervals := [][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95},
		{-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}
	fmt.Println(eraseOverlapIntervals(intervals))
}
