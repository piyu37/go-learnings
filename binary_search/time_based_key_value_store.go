package main

import "fmt"

type tsArray struct {
	ts  int
	val string
}

type TimeMap struct {
	store map[string][]tsArray
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]tsArray),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if tsArr, ok := this.store[key]; ok {
		tsArr = append(tsArr, tsArray{timestamp, value})
		this.store[key] = tsArr
		return
	}

	this.store[key] = []tsArray{{timestamp, value}}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if tsArr, ok := this.store[key]; ok {
		return findPrev(tsArr, timestamp)
	}

	return ""
}

func findPrev(tsArr []tsArray, timestamp int) string {
	first := 0
	last := len(tsArr) - 1

	prev := ""
	for first <= last {
		mid := first + (last-first)/2

		if tsArr[mid].ts <= timestamp {
			prev = tsArr[mid].val
			first = mid + 1
		} else {
			last = mid - 1
		}
	}

	return prev
}

// same as above but uses internal func
// func findPrev(tsArr []tsArray, timestamp int) string {
// 	idx := sort.Search(len(tsArr), func(i int) bool {
// 		return tsArr[i].ts > timestamp
// 	})
// 	if idx == 0 {
// 		return ""
// 	}
// 	return tsArr[idx-1].val
// }

// https://leetcode.com/problems/time-based-key-value-store/description/
func timeBasedKeyValueStore() {
	tkvStore := Constructor()
	tkvStore.Set("foo", "bar", 1)
	fmt.Println(tkvStore.Get("foo", 1))
	fmt.Println(tkvStore.Get("foo", 3))
	tkvStore.Set("foo", "bar2", 4)
	fmt.Println(tkvStore.Get("foo", 4))
	fmt.Println(tkvStore.Get("foo", 5))
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
