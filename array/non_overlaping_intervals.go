package main

import (
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
