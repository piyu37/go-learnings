package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	count := 0
	i := 1

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	lastIntervalValue := intervals[i-1][1]
	for i < len(intervals) {
		if intervals[i][0] < lastIntervalValue {
			if intervals[i][1] < lastIntervalValue {
				lastIntervalValue = intervals[i][1]
			}

			count++
		} else {
			lastIntervalValue = intervals[i][1]
		}

		i++
	}

	return count
}

// https://leetcode.com/problems/non-overlapping-intervals/
func nonOverlapingIntervals() {
	arr := [][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95},
		{-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}
	fmt.Println(eraseOverlapIntervals(arr))
}
