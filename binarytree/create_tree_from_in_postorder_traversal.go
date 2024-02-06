package main

import "fmt"

var postorderIdx int
var inorderMap map[int]int

func buildTree2(inorder []int, postorder []int) *TreeNode {
	inorderMap = map[int]int{}

	for i, v := range inorder {
		inorderMap[v] = i
	}

	postorderIdx = len(postorder) - 1
	return arrayToTree2(postorder, 0, len(postorder)-1)
}

func arrayToTree2(postorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := postorder[postorderIdx]
	root := &TreeNode{
		Val: rootVal,
	}

	postorderIdx--
	root.Right = arrayToTree2(postorder, inorderMap[rootVal]+1, right)
	root.Left = arrayToTree2(postorder, left, inorderMap[rootVal]-1)

	return root
}

// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/?envType=study-plan-v2&envId=top-interview-150
func createTreeFromInorderPostorderTraversal() {
	io := []int{9, 3, 15, 20, 7}
	po := []int{9, 15, 7, 20, 3}

	fmt.Println(buildTree2(io, po))
}
