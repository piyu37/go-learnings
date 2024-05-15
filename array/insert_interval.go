package main

import "fmt"

// check start range with intervals end range
// once it's lesser or equal than end range; have min start from 2
// check end range with start range of intervals
func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	i := 0
	for ; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[0] <= interval[1] {
			break
		}

		result = append(result, interval)
	}

	start := newInterval[0]
	if len(intervals) > 0 && i < len(intervals) {
		start = min(intervals[i][0], newInterval[0])
	}

	for ; i < len(intervals); i++ {
		interval := intervals[i]
		if newInterval[1] < interval[0] {
			break
		}
	}

	end := newInterval[1]
	if i-1 >= 0 {
		end = max(intervals[i-1][1], newInterval[1])
	}

	remaining := intervals[i:]
	result = append(result, []int{start, end})
	result = append(result, remaining...)

	return result
}

// https://leetcode.com/problems/insert-interval/?envType=study-plan-v2&envId=top-interview-150
func insertInterval() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}

	fmt.Println(insert(intervals, newInterval))

	intervals = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	newInterval = []int{4, 8}

	fmt.Println(insert(intervals, newInterval))

	intervals = [][]int{{1, 5}}
	newInterval = []int{6, 8}

	fmt.Println(insert(intervals, newInterval))
}
