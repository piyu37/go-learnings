package main

import "math"

func validateBST(bst *BST) bool {
	min := math.MinInt
	return inorder(bst, &min, true)
}

func inorder(bst *BST, min *int, leftFlag bool) bool {
	if bst != nil {
		valid := inorder(bst.Left, min, true)
		if !valid {
			return false
		}
		if leftFlag {
			if *min >= bst.Value {
				return false
			}
		} else {
			if *min > bst.Value {
				return false
			}
		}

		*min = bst.Value

		return inorder(bst.Right, min, false)
	}

	return true
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/09_validate_bst
func validateBSTMain() {
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
				Value: 15,
			},
		},
	}

	println(validateBST(bst))
}
