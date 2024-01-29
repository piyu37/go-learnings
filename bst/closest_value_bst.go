package main

import "fmt"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func closestValueInBST(bst *BST, target int) int {
	closest := bst.Value
	mainDiff := absDiff(bst.Value, target)

	for bst != nil {
		if target >= bst.Value {
			if bst.Right != nil {
				if absDiff(bst.Right.Value, target) < mainDiff {
					closest = bst.Right.Value
					mainDiff = absDiff(bst.Right.Value, target)
				}

				bst = bst.Right
			} else {
				break
			}
		} else {
			if bst.Left != nil {
				if absDiff(bst.Left.Value, target) < mainDiff {
					closest = bst.Left.Value
					mainDiff = absDiff(bst.Left.Value, target)
				}

				bst = bst.Left
			} else {
				break
			}
		}

	}

	return closest
}

func absDiff(v1, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	}

	return v2 - v1
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/04_find_closest_value
func closestValueBSTMain() {
	bst := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
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
				Right: &BST{
					Value: 14,
				},
			},
			Right: &BST{
				Value: 22,
			},
		},
	}

	fmt.Println(closestValueInBST(bst, 12))
}
