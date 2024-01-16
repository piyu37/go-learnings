package main

import "fmt"

func NodeSwap(head *LinkedList) *LinkedList {
	node := head

	var newHead, tail *LinkedList

	k := 2

	for node != nil {
		join := node
		count := 0
		var prev *LinkedList

		for node != nil && k != count {
			next := node.Next
			node.Next = prev
			prev = node
			node = next
			count++
		}

		if newHead == nil {
			newHead = prev
		}

		if tail != nil {
			tail.Next = prev
		}

		tail = join

	}

	return newHead
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/28_node_swap
func nodeSwapLL() {
	ll := &LinkedList{
		Value: 0,
		Next: &LinkedList{
			Value: 1,
			Next: &LinkedList{
				Value: 2,
				Next: &LinkedList{
					Value: 3,
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

	fmt.Println(NodeSwap(ll))
}
