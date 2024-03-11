package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type friendNode struct {
	timestamp int
	from, to  string
}

func minTimeToBecomeFriendsUsingPrimms(userList, events []string) int {
	friendMap := make(map[string][]friendNode)

	minFlag := false
	start := ""

	for i := range events {
		event := events[i]

		eventInSlice := strings.Split(event, " ")

		timestamp, _ := strconv.Atoi(eventInSlice[0])

		if !minFlag {
			start = eventInSlice[1]
			minFlag = true
		}

		friendMap[eventInSlice[1]] = append(friendMap[eventInSlice[1]],
			friendNode{timestamp, eventInSlice[1], eventInSlice[3]})

		friendMap[eventInSlice[3]] = append(friendMap[eventInSlice[3]],
			friendNode{timestamp, eventInSlice[3], eventInSlice[1]})
	}

	minTime := findMinTime(friendMap, len(userList), start)

	return minTime
}

type minHeap []friendNode

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].timestamp < h[j].timestamp
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v any) {
	*h = append(*h, v.(friendNode))
}

func (h *minHeap) Pop() any {
	arr := *h
	val := arr[len(arr)-1]
	*h = arr[:len(arr)-1]

	return val
}

func findMinTime(friendMap map[string][]friendNode, userCount int, start string) int {
	visited := make(map[string]bool)
	var pq minHeap
	minTime := -1

	heap.Init(&pq)

	heap.Push(&pq, friendNode{0, "", start})

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(friendNode)

		if visited[curr.to] {
			continue
		}

		visited[curr.to] = true
		if curr.timestamp > minTime {
			minTime = curr.timestamp
		}

		friends := friendMap[curr.to]
		for i := range friends {
			friend := friends[i]
			if !visited[friend.to] {
				heap.Push(&pq, friend)
			}
		}
	}

	if userCount != len(visited) {
		return -1
	}

	return minTime
}

func minTimeToBecomeFriendsUsingKruskal(userList, events []string) int {
	ds := initDisjointSet(len(userList))

	userlistIdxMap := make(map[string]int)

	for i := range userList {
		userlistIdxMap[userList[i]] = i
	}

	for i := range events {
		event := events[i]
		eventSlice := strings.Split(event, " ")
		t, _ := strconv.Atoi(eventSlice[0])
		f1 := eventSlice[1]
		f2 := eventSlice[3]
		f1Idx := userlistIdxMap[f1]
		f2Idx := userlistIdxMap[f2]
		f1up := ds.findUPar(f1Idx)
		f2up := ds.findUPar(f2Idx)

		if f1up != f2up {
			ds.unionBySize(f1Idx, f2Idx)
		}

		if ds.size[f1up] == len(userList) || ds.size[f2up] == len(userList) {
			return t
		}
	}

	return -1
}

// user_list = [a, b, c, d, e]

// // modified DFS & min spanning tree
// events = [
// 	10 a & b become friends
// 	11 b & c become friends
// 	12 c & d become friends
// 	13 b & e become friends
// ]

// find the min time in which they will become friends(variant of below leetcode ques)
// https://leetcode.com/problems/the-earliest-moment-when-everyone-become-friends/description/
func googleMinTimeToBecomeFriends() {
	user_list := []string{"a", "b", "c", "d", "e"}

	// modified DFS & min spanning tree
	events := []string{
		"10 a & b become friends",
		"11 b & e become friends",
		"12 c & d become friends",
		"13 a & c become friends",
		"14 b & c become friends",
	}

	fmt.Println(minTimeToBecomeFriendsUsingPrimms(user_list, events))

	fmt.Println(minTimeToBecomeFriendsUsingKruskal(user_list, events))
}
