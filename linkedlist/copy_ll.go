package main

import "fmt"

type NodeWithRandom struct {
	Val          int
	Next, Random *NodeWithRandom
}

func copyRandomList(head *NodeWithRandom) *NodeWithRandom {
	if head == nil {
		return head
	}

	oldNode := head
	for oldNode != nil {
		newNode := &NodeWithRandom{
			Val:    oldNode.Val,
			Next:   oldNode.Next,
			Random: oldNode.Random,
		}

		oldNext := oldNode.Next
		oldNode.Next = newNode
		newNode.Next = oldNext
		oldNode = oldNode.Next.Next
	}

	oldHead := head
	for oldHead != nil {
		if oldHead.Random != nil {
			oldHead.Next.Random = oldHead.Random.Next
		}

		oldHead = oldHead.Next.Next
	}

	var newhead *NodeWithRandom
	oldHead = head
	finalHead := oldHead.Next
	for oldHead != nil {
		if newhead == nil {
			newhead = oldHead.Next
		} else {
			newhead.Next = oldHead.Next
			newhead = newhead.Next
		}

		oldHead.Next = newhead.Next
		oldHead = oldHead.Next
	}

	return finalHead
}

// https://leetcode.com/problems/copy-list-with-random-pointer/?envType=study-plan-v2&envId=top-interview-150
func copyLL() {
	node1 := &NodeWithRandom{
		Val: 7,
	}
	node2 := &NodeWithRandom{
		Val: 13,
	}
	node3 := &NodeWithRandom{
		Val: 11,
	}
	node4 := &NodeWithRandom{
		Val: 10,
	}

	node5 := &NodeWithRandom{
		Val: 1,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	fmt.Println(copyRandomList(node1))
}
