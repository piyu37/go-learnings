package main

import "fmt"

func LinkedListPalindrome(head *LinkedList) bool {
	node := head

	llLength := 0

	for node != nil {
		llLength++
		node = node.Next
	}

	if llLength == 1 {
		return true
	}

	isEven := false
	if llLength%2 == 0 {
		isEven = true
	}

	mid := llLength / 2

	node = head
	count := 1
	for count != mid {
		node = node.Next
		count++
	}

	var prev *LinkedList
	ptr := node
	if !isEven {
		node = node.Next.Next
	} else {
		node = node.Next
	}

	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	ptr.Next = prev

	node = head

	for prev != nil {
		if prev.Value != node.Value {
			return false
		}

		node = node.Next
		prev = prev.Next
	}

	return true
}

func llPalindrome() {
	ll := &LinkedList{
		Value: 0,
		Next: &LinkedList{
			Value: 0,
		},
	}

	fmt.Println(LinkedListPalindrome(ll))
}
