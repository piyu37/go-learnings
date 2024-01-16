package main

import "fmt"

func RearrangeLinkedList2(head *LinkedList, k int) *LinkedList {
	var sh, st *LinkedList
	var eh, et *LinkedList
	var gh, gt *LinkedList

	node := head
	for node != nil {
		if node.Value < k {
			sh, st = growLinkedList(sh, st, node)
		} else if node.Value > k {
			gh, gt = growLinkedList(gh, gt, node)
		} else {
			eh, et = growLinkedList(eh, et, node)
		}

		prevNode := node
		node = node.Next
		prevNode.Next = nil
	}

	firstHead, firstTail := connectLLs(sh, st, eh, et)
	finalHead, _ := connectLLs(firstHead, firstTail, gh, gt)

	return finalHead
}

func connectLLs(h1, t1, h2, t2 *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := h1, t2

	if newHead == nil {
		newHead = h2
	}

	if newTail == nil {
		newTail = t1
	}

	if t1 != nil {
		t1.Next = h2
	}

	return newHead, newTail
}

func growLinkedList(head, tail, node *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := head, node
	if newHead == nil {
		newHead = node
	}

	if tail != nil {
		tail.Next = node
	}

	return newHead, newTail
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/25_rearrange_linked_list
func rearrangeLL2ndSoln() {
	ll := &LinkedList{
		Value: 3,
		Next: &LinkedList{
			Value: 0,
			Next: &LinkedList{
				Value: 5,
				Next: &LinkedList{
					Value: 2,
					Next: &LinkedList{
						Value: 1,
						Next: &LinkedList{
							Value: 1,
						},
					},
				},
			},
		},
	}

	fmt.Println(RearrangeLinkedList2(ll, 3))
}
