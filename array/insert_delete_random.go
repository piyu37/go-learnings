package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	SetMap map[int]int
	SetArr []int
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		SetMap: make(map[int]int),
		SetArr: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.SetMap[val]; ok {
		return false
	}

	rs.SetArr = append(rs.SetArr, val)
	rs.SetMap[val] = len(rs.SetArr) - 1

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	v, ok := rs.SetMap[val]
	if !ok {
		return false
	}

	rs.SetArr[len(rs.SetArr)-1] = v
	rs.SetArr[v], rs.SetArr[len(rs.SetArr)-1] = rs.SetArr[len(rs.SetArr)-1], rs.SetArr[v]
	rs.SetArr = rs.SetArr[:len(rs.SetArr)-1]
	delete(rs.SetMap, v)

	return true
}

func (rs *RandomizedSet) GetRandom() int {
	idx := randInt(0, len(rs.SetArr))

	return rs.SetArr[idx]
}

// https://leetcode.com/problems/insert-delete-getrandom-o1/description/?envType=study-plan-v2&envId=top-interview-150
func insertDeleteRandom() {
	rs := Constructor()
	rs.Insert(1)
	rs.Insert(2)
	rs.Insert(3)
	rs.Insert(4)
	rs.Insert(5)

	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())

}

/**
* Your RandomizedSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
 */
