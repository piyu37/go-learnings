package main

import (
	"fmt"
	"sort"
)

func covid(people []int, d int) int {
	sort.Ints(people)

	left := 0
	right := 1

	minVal := 50000
	for i := 0; i < len(people); i++ {
		if right == i {
			right++
		}

		for left < i && abs(people[left]-people[i]) > d {
			left++
		}

		for right < len(people) && abs(people[right]-people[i]) <= d {
			right++
		}

		if right-left-1 < minVal {
			minVal = right - left - 1
		}
	}

	return minVal
}

// n people
// anyone covid - 1
// 1d array []
// d distance

// i = 0
// j = i+1
// map = 1:2
// nd = nlogn
// people = [1, 2, 5, 7, 3, 8]
// sort = [1, 2, 3, 5, 7, 8] ? right = 6
// i := 1, l = 0, r = 3; l < i; count = right-left-1 = 2; people[r] < d+people[i]
// map
// distance = 2
// 1 = 2, 3
// 2 = 1, 3
// 5 = 3, 7
// 8 = 7
// Need to find index/value of person on which covid infection would be minimum
// if that person infected.
// input => people = [1, 2, 5, 7, 3, 8]; d = 2 => 5th index/8 value will only infect 1(which is 7)
// whereas other no. will atleast infect more than 1
func mmtInterviewCovid() {
	arr := []int{1, 2, 5, 9, 10, 3, 8}
	fmt.Println(covid(arr, 2))
}
