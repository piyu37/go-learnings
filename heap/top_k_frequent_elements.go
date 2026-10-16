package main

import (
	"container/heap"
	"fmt"
)

type minHeapFrequency [][]int

func (mh minHeapFrequency) Len() int {
	return len(mh)
}

func (mh *minHeapFrequency) Push(x any) {
	*mh = append(*mh, x.([]int))
}

func (mh *minHeapFrequency) Pop() any {
	val := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]

	return val
}

func (mh minHeapFrequency) Less(i, j int) bool {
	return mh[i][1] < mh[j][1]
}

func (mh minHeapFrequency) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func topKFrequent(nums []int, k int) []int {
	valMap := make(map[int]int)

	for i := range nums {
		valMap[nums[i]]++
	}

	mhf := &minHeapFrequency{}

	for key, val := range valMap {
		heap.Push(mhf, []int{key, val})
		if mhf.Len() > k {
			heap.Pop(mhf)
		}
	}

	result := make([]int, k)
	for i := range k {
		r := heap.Pop(mhf).([]int)
		result[i] = r[0]
	}

	return result
}

func topKFrequentUsingBucketSort(nums []int, k int) []int {
	valMap := make(map[int]int)

	for i := range nums {
		valMap[nums[i]]++
	}

	// buckets[f] holds all numbers that appear exactly f times
	buckets := make([][]int, len(nums)+1)
	for key, freq := range valMap {
		buckets[freq] = append(buckets[freq], key)
	}

	result := make([]int, 0, k)
	for freq := len(buckets) - 1; freq > 0 && len(result) < k; freq-- {
		for _, num := range buckets[freq] {
			result = append(result, num)
			if len(result) == k {
				break
			}
		}
	}

	return result
}

// https://leetcode.com/problems/top-k-frequent-elements/description/
func topKFrequentElements() {
	nums := []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}
	k := 2

	fmt.Println(topKFrequent(nums, k))
	fmt.Println(topKFrequentUsingBucketSort(nums, k))
}
