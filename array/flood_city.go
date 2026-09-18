package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	lakeMap := make(map[int]int)
	dryDays := make([]int, 0) // sorted, currently-unused dry day indices
	ans := make([]int, len(rains))

	for i, river := range rains {
		if river != 0 {
			lastRiverIdx, exist := lakeMap[river]
			if exist {
				// sort.SearchInts binary-searches the sorted dryDays slice and
				// returns the smallest index pos such that dryDays[pos] >= lastRiverIdx
				// (or len(dryDays) if every element is smaller).
				// It repeatedly halves the search range, moving low/high toward
				// the boundary between "< lastRiverIdx" and ">= lastRiverIdx".
				//
				// e.g. dryDays = [2, 5, 9, 14], lastRiverIdx = 6
				//   lo=0 hi=4 mid=2 dryDays[2]=9  >=6 -> hi=2
				//   lo=0 hi=2 mid=1 dryDays[1]=5  <6  -> lo=2
				//   lo=2 hi=2 loop ends -> pos=2, dryDays[2]=9 (first value >= 6)
				pos := sort.SearchInts(dryDays, lastRiverIdx) // smallest dryDays[pos] >= lastRiverIdx
				if pos == len(dryDays) {
					return []int{}
				}
				dryDayIdx := dryDays[pos]
				ans[dryDayIdx] = river
				dryDays = append(dryDays[:pos], dryDays[pos+1:]...)
			}
			lakeMap[river] = i
			ans[i] = -1
			continue
		}

		dryDays = append(dryDays, i)
		ans[i] = 1 // default filler if this day never ends up needed
	}

	return ans
}

// segTree is a sum segment tree over dry-day indices [0, n-1], used to
// answer "smallest unused dry day >= x" in true O(log n), including the
// removal — unlike a sorted slice (avoidFlood), whose removal is O(n) and
// makes the overall algorithm O(n^2) in the worst case.
type segTree struct {
	n    int
	tree []int
}

func newSegTree(n int) *segTree {
	return &segTree{n: n, tree: make([]int, 4*n)}
}

func (st *segTree) update(node, l, r, idx, val int) {
	if l == r {
		st.tree[node] = val
		return
	}
	mid := (l + r) / 2
	if idx <= mid {
		st.update(node*2, l, mid, idx, val)
	} else {
		st.update(node*2+1, mid+1, r, idx, val)
	}
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

func (st *segTree) query(node, l, r, x int) int {
	if r < x || st.tree[node] == 0 {
		return -1
	}
	if l == r {
		return l
	}
	mid := (l + r) / 2
	if left := st.query(node*2, l, mid, x); left != -1 {
		return left
	}
	return st.query(node*2+1, mid+1, r, x)
}

func (st *segTree) add(idx int)    { st.update(1, 0, st.n-1, idx, 1) }
func (st *segTree) remove(idx int) { st.update(1, 0, st.n-1, idx, 0) }

func (st *segTree) firstAvailableFrom(x int) int {
	if x > st.n-1 {
		return -1
	}
	return st.query(1, 0, st.n-1, x)
}

// avoidFloodOptimal is O(n log n) worst case: both finding the smallest
// unused dry day >= lastRiverIdx and removing it cost O(log n).
func avoidFloodOptimal(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	lakeMap := make(map[int]int)
	dryDays := newSegTree(n)

	for i, river := range rains {
		if river != 0 {
			if lastRiverIdx, exist := lakeMap[river]; exist {
				dryDayIdx := dryDays.firstAvailableFrom(lastRiverIdx)
				if dryDayIdx == -1 {
					return []int{}
				}
				ans[dryDayIdx] = river
				dryDays.remove(dryDayIdx)
			}
			lakeMap[river] = i
			ans[i] = -1
			continue
		}

		dryDays.add(i)
		ans[i] = 1 // default filler if this day never ends up needed
	}

	return ans
}

// https://leetcode.com/problems/avoid-flood-in-the-city/description/
func floodCity() {
	rains := []int{1, 0, 2, 0, 2, 1}
	// this cods is fine. No need to write optimal code. Below optimal code is just for knowledge
	fmt.Println(avoidFlood(rains))

	rains = []int{1, 0, 2, 0, 2, 1}
	fmt.Println(avoidFloodOptimal(rains))

}
