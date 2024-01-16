package main

import "fmt"

func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	h1 := headOne
	h2 := headTwo

	for h1 != nil && h2 != nil {
		v1 := h1.Value
		v2 := h2.Value

		if v2 < v1 {
			temp := *h1
			h1.Value = v2
			h1.Next = &temp
			h2 = h2.Next
		}

		if h1.Next == nil {
			h1.Next = h2
			break
		}

		h1 = h1.Next
	}

	return headOne
}

// mergeLL to merge 2 sorted LL in one sorted LL
func mergeLL() {
	h2 := &LinkedList{
		Value: 2,
		Next: &LinkedList{
			Value: 6,
			Next: &LinkedList{
				Value: 7,
				Next: &LinkedList{
					Value: 8,
				},
			},
		},
	}

	h1 := &LinkedList{
		Value: 1,
		Next: &LinkedList{
			Value: 3,
			Next: &LinkedList{
				Value: 4,
				Next: &LinkedList{
					Value: 5,
					Next: &LinkedList{
						Value: 9,
						Next: &LinkedList{
							Value: 10,
						},
					},
				},
			},
		},
	}

	fmt.Println(MergeLinkedLists(h1, h2))
}
