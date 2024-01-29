package main

import "fmt"

var prev *TreeNode

func flattend_bt(root *TreeNode) {
	if root == nil {
		return
	}

	flattend_bt(root.Right)

	root.Right = prev
	prev = root

	flattend_bt(root.Left)

	root.Left = nil
}

func flattendBTMain() {
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

	flattend_bt(tree)
	fmt.Println(tree)
}
