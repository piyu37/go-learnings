package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	i := 0
	for i < len(intervals) {
		min := intervals[i][0]
		max := intervals[i][1]

		i++
		for i < len(intervals) && intervals[i][0] <= max {
			if intervals[i][1] > max {
				max = intervals[i][1]
			}

			i++
		}

		result = append(result, []int{min, max})
	}

	return result
}

// https://leetcode.com/problems/merge-intervals/description/
func mergeMain() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))
}
