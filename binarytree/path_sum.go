package main

import "fmt"

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return calculatePathSum(root, targetSum, 0)
}

func calculatePathSum(root *TreeNode, targetSum, sum int) bool {
	if root.Left == nil && root.Right == nil {
		if root.Val+sum == targetSum {
			return true
		} else {
			return false
		}
	}

	if root.Left != nil && calculatePathSum(root.Left, targetSum, sum+root.Val) {
		return true
	}

	return root.Right != nil && calculatePathSum(root.Right, targetSum, sum+root.Val)
}

func hasPathSumIter(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	de := make([]queue, 0)

	de = append(de, queue{targetSum - root.Val, root})

	for len(de) > 0 {
		currSum, node := de[0].value, de[0].node
		de = de[1:]

		if node.Left == nil && node.Right == nil && currSum == 0 {
			return true
		}

		if node.Right != nil {
			de = append(de, queue{currSum - node.Right.Val, node.Right})
		}
		if node.Left != nil {
			de = append(de, queue{currSum - node.Left.Val, node.Left})
		}
	}

	return false
}

// https://leetcode.com/problems/path-sum/description/?envType=study-plan-v2&envId=top-interview-150
func pathSum() {
	arr := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}

	root := buildTreeFromArray(arr, 0)

	fmt.Println(hasPathSum(root, 22))

	arr = []interface{}{1, 2, 3}

	root = buildTreeFromArray(arr, 0)

	fmt.Println(hasPathSum(root, 5))
	fmt.Println(hasPathSumIter(root, 5))
}
