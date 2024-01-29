package main

import "fmt"

func kthLargest(bst *BST, k int) int {
	return reverseInorder(bst, &k)
}

func reverseInorder(bst *BST, k *int) int {
	if bst == nil {
		return -1
	}

	val := reverseInorder(bst.Right, k)
	if val > -1 {
		return val
	}

	if *k == 1 {
		return bst.Value
	}

	*k--

	return reverseInorder(bst.Left, k)
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/35_find_kth_largest_value_in_bst
func kthLargestMain() {
	bst := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 2,
					Left: &BST{
						Value: 1,
					},
				},
			},
			Right: &BST{
				Value: 5,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left: &BST{
					Value: 11,
					Right: &BST{
						Value: 12,
					},
				},
				Right: &BST{
					Value: 14,
				},
			},
			Right: &BST{
				Value: 22,
			},
		},
	}

	fmt.Println(kthLargest(bst, 12))
}
