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
}
