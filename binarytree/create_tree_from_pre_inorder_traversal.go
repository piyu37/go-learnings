package main

import "fmt"

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)

	for i, v := range inorder {
		inorderMap[v] = i
	}

	return arrayToTree(preorder, new(int), inorderMap, 0, len(preorder)-1)
}

func arrayToTree(preorder []int, preorderIdx *int, inorderMap map[int]int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := preorder[*preorderIdx]
	*preorderIdx++
	root := &TreeNode{
		Val: rootVal,
	}

	root.Left = arrayToTree(preorder, preorderIdx, inorderMap, left, inorderMap[rootVal]-1)
	root.Right = arrayToTree(preorder, preorderIdx, inorderMap, inorderMap[rootVal]+1, right)

	return root
}

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/?envType=study-plan-v2&envId=top-interview-150
func createTreeFromPreorderInorderTraversal() {
	po := []int{3, 9, 20, 15, 7}
	io := []int{9, 3, 15, 20, 7}

	fmt.Println(buildTree(po, io))

}
