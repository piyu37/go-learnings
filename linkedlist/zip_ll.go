package main

import "fmt"

func ZipLinkedList(head *LinkedList) *LinkedList {
	node := head

	llLength := 0

	for node != nil {
		llLength++
		node = node.Next
	}

	if llLength == 2 {
		return head
	}

	mid := llLength / 2

	node = head
	count := 1
	for count != mid {
		node = node.Next
		count++
	}

	isEven := false
	if llLength%2 == 0 {
		isEven = true
	}

	var prev *LinkedList
	ptr := node
	if isEven {
		node = node.Next
	} else {
		node = node.Next.Next
		ptr = ptr.Next
	}

	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}

	ptr.Next = nil

	node = head

	for prev != nil && node != nil {
		temp := node.Next
		node.Next = prev
		prev = prev.Next
		node.Next.Next = temp
		node = node.Next.Next
	}

	return head
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/27_zip_linked_list
func zipLL() {
	ll := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 2,
			Next: &LinkedList{
				Value: 3,
				Next: &LinkedList{
					Value: 4,
					// 	Next: &LinkedList{
					// 		Value: 5,
					// 		Next: &LinkedList{
					// 			Value: 6,
					// 			// Next: &LinkedList{
					// 			// 	Value: 7,
					// 			// },
					// 		},
					// 	},
				},
			},
		},
	}

	fmt.Println(ZipLinkedList(ll))
}
