package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {

	result := make([]int, 0)

	calculateSum(root, 0, &result)

	return result
}

func calculateSum(root *BinaryTree, count int, result *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, count+root.Value)

		return
	}

	calculateSum(root.Left, count+root.Value, result)

	calculateSum(root.Right, count+root.Value, result)
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/05_branch_sums
func branchSumMain() {
	tree := &BinaryTree{
		Value: 0,
		Right: &BinaryTree{
			Value: 1,
			Right: &BinaryTree{
				Value: 10,
				Right: &BinaryTree{
					Value: 100,
				},
			},
		},
	}

	fmt.Println(BranchSums(tree))
}
