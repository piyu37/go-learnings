package main

import "math"

func validateBST1(bst *BST) bool {
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

func validateBST2(root *BST) bool {
	inorderArr := []int{}

	inorderTraversal(root, &inorderArr)

	for i := 1; i < len(inorderArr); i++ {
		if inorderArr[i] <= inorderArr[i-1] {
			return false
		}
	}

	return true
}

func validateBST3(root *BST) bool {
	prev := math.MinInt64
	return inorderBST3(root, &prev)
}

func inorderBST3(root *BST, prev *int) bool {
	if root == nil {
		return true
	}

	if !inorderBST3(root.Left, prev) {
		return false
	}

	if root.Value <= *prev {
		return false
	}

	*prev = root.Value

	return inorderBST3(root.Right, prev)
}

func inorderTraversal(root *BST, inorderArr *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, inorderArr)
	*inorderArr = append(*inorderArr, root.Value)
	inorderTraversal(root.Right, inorderArr)
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

	println(validateBST1(bst))
	println(validateBST2(bst))
	println(validateBST3(bst))
}
