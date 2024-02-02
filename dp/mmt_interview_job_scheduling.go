package main

import (
	"fmt"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	jobsWithProfit := make([][3]int, len(profit))
	for i := range profit {
		jobsWithProfit[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}

	sort.Slice(jobsWithProfit, func(i, j int) bool {
		return jobsWithProfit[i][0] < jobsWithProfit[j][0]
	})

	memoizatn := make(map[int]int)

	return scheduleJobs(jobsWithProfit, 0, memoizatn)
}

func scheduleJobs(jobsWithProfit [][3]int, current int, memoizatn map[int]int) int {
	var profitWithCurrent int

	if current >= len(jobsWithProfit) {
		return 0
	}

	if profit, ok := memoizatn[current]; ok {
		return profit
	}

	profitWithoutCurrent := scheduleJobs(jobsWithProfit, current+1, memoizatn)
	if current+1 < len(jobsWithProfit) && jobsWithProfit[current+1][0] >= jobsWithProfit[current][1] {
		profitWithCurrent = jobsWithProfit[current][2] + profitWithoutCurrent
	} else {
		idx := findNext(jobsWithProfit, jobsWithProfit[current][1], current+1, len(jobsWithProfit)-1)
		profitWithCurrent = jobsWithProfit[current][2] + scheduleJobs(jobsWithProfit, idx, memoizatn)
	}

	if profitWithCurrent > profitWithoutCurrent {
		memoizatn[current] = profitWithCurrent

		return profitWithCurrent
	}

	memoizatn[current] = profitWithoutCurrent

	return profitWithoutCurrent
}

func findNext(jobsWithProfit [][3]int, currentJobEnd, start, end int) int {
	result := len(jobsWithProfit)
	for start <= end {
		mid := start + (end-start)/2

		if jobsWithProfit[mid][0] >= currentJobEnd {
			result = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return result
}

// https://leetcode.com/problems/maximum-profit-in-job-scheduling/description/
func mmtInterviewJobScheduling() {
	startTime := []int{6, 15, 7, 11, 1, 3, 16, 2}
	endTime := []int{19, 18, 19, 16, 10, 8, 19, 8}
	profit := []int{2, 9, 1, 19, 5, 7, 3, 19}

	fmt.Println(jobScheduling(startTime, endTime, profit))
}
