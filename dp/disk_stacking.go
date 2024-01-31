package main

import (
	"fmt"
	"sort"
)

func diskStacking(disks [][]int) [][]int {
	sort.Slice(disks, func(i, j int) bool {
		return disks[i][2] < disks[j][2]
	})

	result := make([][]int, 0)
	msis := make([]int, len(disks))
	maxSum := 0

	for i := range disks {
		disk := disks[i]

		msis[i] = disk[2]
	}

	for i := 1; i < len(disks); i++ {
		iDisk := disks[i]
		for j := 0; j < i; j++ {
			jDisk := disks[j]

			if (iDisk[0] > jDisk[0] && iDisk[1] > jDisk[1] && iDisk[2] > jDisk[2]) && msis[i] < iDisk[2]+msis[j] {
				msis[i] = iDisk[2] + msis[j]
				if maxSum < iDisk[2]+msis[j] {
					maxSum = iDisk[2] + msis[j]
				}
			}
		}
	}

	temp := maxSum
	for i := len(msis) - 1; i >= 0; i-- {
		if msis[i] == temp {
			result = append(result, disks[i])
			temp -= disks[i][2]
		}
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/15_disk_stacking
func diskStackingMain() {
	disks := [][]int{
		{2, 1, 2},
		{3, 2, 3},
		{2, 2, 8},
		{2, 3, 4},
		{1, 2, 1},
		{4, 4, 5},
		{1, 1, 4},
	}

	fmt.Println(diskStacking(disks))
}
