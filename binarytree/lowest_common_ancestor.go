package main

import "fmt"

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := &TreeNode{}
	isLCA(root, p, q, ans)

	return ans
}

func isLCA(root, p, q, ans *TreeNode) bool {
	if root == nil {
		return false
	}

	leftLCA := isLCA(root.Left, p, q, ans)
	rightLCA := isLCA(root.Right, p, q, ans)

	mid := root == p || root == q

	if leftLCA && rightLCA {
		*ans = *root
	} else if (leftLCA || rightLCA) && mid {
		*ans = *root
	}

	return mid || leftLCA || rightLCA
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func lowestCommonAncestorMain() {
	arr := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}

	root := buildTreeFromArray(arr, 0)

	p := root.Left.Left
	q := root.Left.Right.Right

	fmt.Println(lowestCommonAncestor(root, p, q))
}
