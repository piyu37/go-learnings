package main

import "fmt"

func NodeSwapR(head *LinkedList) *LinkedList {
	node := head

	var prev *LinkedList

	for node != nil && node.Next != nil {
		first := node
		second := node.Next

		first.Next = second.Next
		second.Next = first
		if prev == nil {
			head = second
		} else {
			prev.Next = second
		}

		prev = first
		node = first.Next
	}

	return head
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/28_node_swap
func nodeSwapLLRecur() {
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

	fmt.Println(NodeSwapR(ll))
}
