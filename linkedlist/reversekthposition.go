package main

import "fmt"

type Node2 struct {
	data int
	next *Node2
}

func reverse(head *Node2, k int) *Node2 {
	current := head
	var next, prev *Node2
	count := 0
	for current != nil && count < k {
		next = current.next
		current.next = prev
		prev = current
		current = next
		count++
	}
	if next != nil {
		head.next = reverse(next, k)
	}
	return prev
}

// e.g. 1 -> 2 -> 3 -> 4; k = 2 => 2 -> 1 -> 4 -> 3
func reverseKthPosition() {
	head := &Node2{data: 1}
	head.next = &Node2{data: 2}
	head.next.next = &Node2{data: 3}
	head.next.next.next = &Node2{data: 4}
	head.next.next.next.next = &Node2{data: 5}

	ll := reverse(head, 2)

	// Print the reversed linked list
	current := ll
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
