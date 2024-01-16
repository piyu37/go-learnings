package main

import "fmt"

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	currentNode := head

	llLength := 0

	for currentNode != nil {
		llLength++
		currentNode = currentNode.Next
	}

	positionNodeToBeDeleted := llLength - k

	currentNode = head
	if positionNodeToBeDeleted == 0 {
		currentNode.Value = currentNode.Next.Value
		currentNode.Next = currentNode.Next.Next
		return
	}

	count := 1
	for currentNode != nil && positionNodeToBeDeleted != count {
		count++
		currentNode = currentNode.Next
	}

	currentNode.Next = currentNode.Next.Next
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/23_remove_kth_node_from_end
func removeLastKNode() {
	ll := &LinkedList{Value: 1, Next: &LinkedList{Value: 2, Next: &LinkedList{Value: 3}}}
	RemoveKthNodeFromEnd(ll, 3)
	fmt.Println(ll)
}
