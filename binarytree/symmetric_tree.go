package main

import "fmt"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return isSymmetricTree(root1.Left, root2.Right) &&
		isSymmetricTree(root1.Right, root2.Left)
}

// https://leetcode.com/problems/symmetric-tree/?envType=study-plan-v2&envId=top-interview-150
func symmetricTree() {
	a := []int{1, 2, 2, 3, 4, 4, 3}
	bt := BT{}

	for _, n := range a {
		bt.Insert(n)
	}

	fmt.Println(isSymmetric(bt.root))
}
