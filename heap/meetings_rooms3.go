package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type roomMeetingInfo struct {
	end, room int
}

type minUsedRoomsHeap []roomMeetingInfo

func (mh minUsedRoomsHeap) Len() int {
	return len(mh)
}

func (mh minUsedRoomsHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh minUsedRoomsHeap) Less(i, j int) bool {
	if mh[i].end != mh[j].end {
		return mh[i].end < mh[j].end
	}

	return mh[i].room < mh[j].room
}

func (mh *minUsedRoomsHeap) Push(x any) {
	*mh = append(*mh, x.(roomMeetingInfo))
}

func (mh *minUsedRoomsHeap) Pop() any {
	top := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]

	return top
}

// sort.IntSlice already gives Len, Less and Swap
type minRoomHeap []int

func (rh minRoomHeap) Len() int {
	return len(rh)
}

func (rh minRoomHeap) Swap(i, j int) {
	rh[i], rh[j] = rh[j], rh[i]
}

func (rh minRoomHeap) Less(i, j int) bool {
	return rh[i] < rh[j]
}

func (rh *minRoomHeap) Push(x any) {
	*rh = append(*rh, x.(int))
}

func (rh *minRoomHeap) Pop() any {
	top := (*rh)[len(*rh)-1]
	*rh = (*rh)[:len(*rh)-1]

	return top
}

func mostBooked(n int, meetings [][]int) int {
	sortedUsedRoomsByEndHeap := &minUsedRoomsHeap{}
	sortedFreeRoomHeap := &minRoomHeap{}
	count := make([]int, n)

	// rooms appended in ascending order already form a valid min-heap
	for roomNo := range n {
		*sortedFreeRoomHeap = append(*sortedFreeRoomHeap, roomNo)
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, interval := range meetings {
		start := interval[0]
		end := interval[1]

		for sortedUsedRoomsByEndHeap.Len() != 0 && start >= (*sortedUsedRoomsByEndHeap)[0].end {
			meetingEndedInfo := heap.Pop(sortedUsedRoomsByEndHeap).(roomMeetingInfo)
			heap.Push(sortedFreeRoomHeap, meetingEndedInfo.room)
		}

		if sortedFreeRoomHeap.Len() == 0 {
			meetingEndedInfo := heap.Pop(sortedUsedRoomsByEndHeap).(roomMeetingInfo)
			meetingEndedInfo.end += end - start
			heap.Push(sortedUsedRoomsByEndHeap, meetingEndedInfo)
			count[meetingEndedInfo.room]++
		} else {
			roomNo := heap.Pop(sortedFreeRoomHeap).(int)
			heap.Push(sortedUsedRoomsByEndHeap, roomMeetingInfo{end, roomNo})
			count[roomNo]++
		}
	}

	ans := 0
	for roomNo := 1; roomNo < n; roomNo++ {
		if count[roomNo] > count[ans] {
			ans = roomNo
		}
	}

	return ans
}

// https://leetcode.com/problems/meeting-rooms-iii/
func meetingsRooms3() {
	meetings := [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}
	fmt.Println(mostBooked(3, meetings))
}
