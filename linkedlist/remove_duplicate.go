package main

import "fmt"

// This is an input struct. Do not edit.
// type LinkedList struct {
// 	Value int
// 	Next  *LinkedList
// }

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	start := linkedList
	for start != nil && start.Next != nil {
		if start.Value == start.Next.Value {
			start.Next = start.Next.Next
		} else {
			start = start.Next
		}
	}

	return linkedList
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/?envType=study-plan-v2&envId=top-interview-150
func deleteDuplicates(head *ListNode) *ListNode {
	start := &ListNode{0, head}
	dummy := start
	for start.Next != nil {
		flag := false
		for start.Next.Next != nil {
			if start.Next.Val == start.Next.Next.Val {
				start.Next = start.Next.Next
				flag = true
			} else {
				break
			}
		}

		if flag {
			start.Next = start.Next.Next
		} else {
			start = start.Next
		}
	}

	return dummy.Next
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/20_remove_duplicates_from_linkedlist
func removeDuplicate() {
	ll := LinkedList{
		Value: 1,
	}

	ll.Next = &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 3,
			Next: &LinkedList{
				Value: 4,
				Next: &LinkedList{
					Value: 4,
					Next: &LinkedList{
						Value: 4,
						Next: &LinkedList{
							Value: 5,
						},
					},
				},
			},
		},
	}

	fmt.Println(RemoveDuplicatesFromLinkedList(&ll))

	ll2 := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	fmt.Println(deleteDuplicates(&ll2))

	ll2 = ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	fmt.Println(deleteDuplicates(&ll2))
}
