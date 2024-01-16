package main

import (
	"fmt"
	"math"
)

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	var result *LinkedList

	ll1 := linkedListOne
	ll2 := linkedListTwo
	ll1Value := 0
	unitPlace := 0

	for ll1 != nil {
		ll1Value += ll1.Value * int(math.Pow10(unitPlace))
		unitPlace++
		ll1 = ll1.Next
	}

	ll2Value := 0
	unitPlace = 0

	for ll2 != nil {
		ll2Value += ll2.Value * int(math.Pow10(unitPlace))
		unitPlace++
		ll2 = ll2.Next
	}

	total := ll1Value + ll2Value

	if total == 0 {
		return &LinkedList{
			Value: 0,
		}
	}

	var prevNode *LinkedList

	for total > 0 {
		unitValue := total % 10
		node := &LinkedList{
			Value: unitValue,
		}

		if result == nil {
			result = node
			prevNode = result
			total = total / 10
			continue
		}

		prevNode.Next = node
		prevNode = prevNode.Next
		total = total / 10
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/46_sum_of_linked_lists
func llSum() {
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

	ll2 := &LinkedList{
		Value: 9,
		Next: &LinkedList{
			Value: 4,
			Next: &LinkedList{
				Value: 5,
			},
		},
	}
	fmt.Println(SumOfLinkedLists(ll1, ll2))
}
