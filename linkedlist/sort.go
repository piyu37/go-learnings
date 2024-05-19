package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)

	left := head
	right := mid.Next

	mid.Next = nil

	left = sortList(left)
	right = sortList(right)
	return merge(left, right)
}

func merge(list1, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummyHead.Next
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

var (
	tail        *ListNode
	nextSubList *ListNode
)

func sortListWithConstantSpace(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := getCount(head)
	start := head
	dummyHead := &ListNode{}
	for size := 1; size < n; size = size * 2 {
		tail = dummyHead
		for start != nil {
			if start.Next == nil {
				tail.Next = start
				break
			}
			mid := split(start, size)
			merge2(start, mid)
			start = nextSubList
		}
		start = dummyHead.Next
	}
	return dummyHead.Next
}

func split(start *ListNode, size int) *ListNode {
	midPrev := start
	end := start.Next
	for index := 1; index < size && (midPrev.Next != nil || end.Next != nil); index++ {
		if end.Next != nil {
			if end.Next.Next != nil {
				end = end.Next.Next
			} else {
				end = end.Next
			}
		}
		if midPrev.Next != nil {
			midPrev = midPrev.Next
		}
	}
	mid := midPrev.Next
	nextSubList = end.Next
	midPrev.Next = nil
	if end != nil {
		end.Next = nil
	}
	return mid
}

func getCount(head *ListNode) int {
	cnt := 0
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		cnt++
	}
	return cnt
}

func merge2(list1 *ListNode, list2 *ListNode) {
	dummyHead := &ListNode{}
	newTail := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newTail.Next = list1
			list1 = list1.Next
		} else {
			newTail.Next = list2
			list2 = list2.Next
		}
		newTail = newTail.Next
	}
	if list1 != nil {
		newTail.Next = list1
	} else {
		newTail.Next = list2
	}
	for newTail.Next != nil {
		newTail = newTail.Next
	}
	tail.Next = dummyHead.Next
	tail = newTail
}

// https://leetcode.com/problems/sort-list/?envType=study-plan-v2&envId=top-interview-150
func sort() {
	first := ListNode{Val: 4}
	sec := ListNode{Val: 2}
	third := ListNode{Val: 1}
	forth := ListNode{Val: 3}

	first.Next = &sec
	sec.Next = &third
	third.Next = &forth

	fmt.Println(sortList(&first))

	first = ListNode{Val: 4}
	sec = ListNode{Val: 2}
	third = ListNode{Val: 1}
	forth = ListNode{Val: 3}

	first.Next = &sec
	sec.Next = &third
	third.Next = &forth

	fmt.Println(sortListWithConstantSpace(&first))
}
