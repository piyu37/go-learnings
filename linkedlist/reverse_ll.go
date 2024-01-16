package main

import "fmt"

func ReverseLinkedList(head *LinkedList) *LinkedList {
	current := head

	var prev, next *LinkedList

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func reverseLL() {
	ll1 := &LinkedList{
		Value: 2,
		Next: &LinkedList{
			Value: 4,
			Next: &LinkedList{
				Value: 7,
				Next: &LinkedList{
					Value: 1,
				},
			},
		},
	}

	fmt.Println(ReverseLinkedList(ll1))
}
