package main

import "fmt"

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := findLeftTreeHeight(root)
	rh := findRightTreeHeight(root)

	if lh == rh {
		// 1<<2 is same to 2^2; 1<<3=2^3; 1<<4 = 2^4
		// shilfting 1 in left by 2 giving 100 in binary; converts 4 in integer
		// left bit shift operation(1 << 2); more efficient than doing 2^2
		return (1 << lh) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func findLeftTreeHeight(root *TreeNode) int {
	count := 0

	for root != nil {
		count++
		root = root.Left
	}

	return count
}

func findRightTreeHeight(root *TreeNode) int {
	count := 0

	for root != nil {
		count++
		root = root.Right
	}

	return count
}

// https://leetcode.com/problems/count-complete-tree-nodes/description/?envType=study-plan-v2&envId=top-interview-150
func countCompleteTreenNodes() {
	arr := []interface{}{1, 2, 3, 4, 5, 6}

	root := buildTreeFromArray(arr, 0)
	fmt.Println(countNodes(root))
}
