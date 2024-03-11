package main

import (
	"fmt"
	"sort"
)

func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	ds := initDisjointSet(n)
	for i := range logs {
		log := logs[i]
		timestamp := log[0]
		f1 := log[1]
		f2 := log[2]

		f1up := ds.findUPar(f1)
		f2up := ds.findUPar(f2)

		if f1up != f2up {
			ds.unionBySize(f1, f2)
		}

		if ds.size[f1up] == n || ds.size[f2up] == n {
			return timestamp
		}
	}

	return -1
}

// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/description/
func earliestMomentFriends() {
	logs := [][]int{
		{20190101, 0, 1},
		{20190104, 3, 4},
		{20190107, 2, 3},
		{20190211, 1, 5},
		{20190224, 2, 4},
		{20190301, 0, 3},
		{20190312, 1, 2},
		{20190322, 4, 5},
	}

	fmt.Println(earliestAcq(logs, 6))

	logs = [][]int{
		{0, 2, 0},
		{1, 0, 1},
		{3, 0, 3},
		{4, 1, 2},
		{7, 3, 1},
	}

	fmt.Println(earliestAcq(logs, 4))

	logs = [][]int{
		{9, 3, 0},
		{0, 2, 1},
		{8, 0, 1},
		{1, 3, 2},
		{2, 2, 0},
		{3, 3, 1},
	}

	fmt.Println(earliestAcq(logs, 4))
}
