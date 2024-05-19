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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy
	// Advances first pointer so that the gap between first and second is n nodes apart
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	// Move first to the end, maintaining the gap
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/23_remove_kth_node_from_end
func removeLastKNode() {
	ll := &LinkedList{Value: 1, Next: &LinkedList{Value: 2, Next: &LinkedList{Value: 3}}}
	RemoveKthNodeFromEnd(ll, 3)
	fmt.Println(ll)

	ll2 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	removeNthFromEnd(ll2, 2)
	fmt.Println(ll2)
}
