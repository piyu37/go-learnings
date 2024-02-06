package main

import "fmt"

func sumNumbers(root *TreeNode) int {
	return calculateSumTillLeaf(root, 0)
}

func calculateSumTillLeaf(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	sum = (sum * 10) + root.Val

	leftSum := calculateSumTillLeaf(root.Left, sum)
	rightSum := calculateSumTillLeaf(root.Right, sum)

	if leftSum+rightSum > 0 {
		return leftSum + rightSum
	}

	return sum
}

// https://leetcode.com/problems/sum-root-to-leaf-numbers/description/?envType=study-plan-v2&envId=top-interview-150
func sumRootToLeaf() {
	arr := []interface{}{4, 9, 0, 5, 1}

	root := buildTreeFromArray(arr, 0)

	fmt.Println(sumNumbers(root))
}
