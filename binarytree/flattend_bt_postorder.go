package main

import "fmt"

func flattend_bt_postorder(root *TreeNode) {
	if root == nil {
		return
	}

	nextRight := root.Right
	nextLeft := root.Left
	root.Right = prev
	prev = root
	root.Left = nil

	flattend_bt_postorder(nextRight)

	flattend_bt_postorder(nextLeft)
}

func postorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Println(root.Val)
}

func flattendBTPostorderMain() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 8,
				},
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}

	postorderTraversal(tree)
	flattend_bt_postorder(tree)
	fmt.Println(tree)
}
