package main

import "fmt"

type BT struct {
	root *TreeNode
}

func binaryTree() {
	a := []int{10, 11, 9, 7, 8}
	bt := BT{}

	for _, n := range a {
		bt.Insert(n)
	}

	bt.PreOrderPrint(bt.root)

	treeArray := []interface{}{1, 2, 3, 4, nil, 5, nil}

	// Build the tree from the array
	root := buildTreeFromArray(treeArray, 0)

	bt.root = root
	bt.PreOrderPrint(root)
}

func (bt *BT) Insert(value int) {
	if bt.root == nil {
		bt.root = &TreeNode{Val: value}
	} else {
		arr := []*TreeNode{bt.root}

		for len(arr) > 0 {
			temp := arr[0]
			arr = arr[1:]

			if temp.Left == nil {
				temp.Left = &TreeNode{Val: value}
				break
			} else {
				arr = append(arr, temp.Left)
			}

			if temp.Right == nil {
				temp.Right = &TreeNode{Val: value}
				break
			} else {
				arr = append(arr, temp.Right)
			}
		}
	}
}

func (bt *BT) PreOrderPrint(n *TreeNode) {
	if n != nil {
		fmt.Println(n.Val)
		bt.PreOrderPrint(n.Left)
		bt.PreOrderPrint(n.Right)
	}
}

func buildTreeFromArray(arr []interface{}, index int) *TreeNode {
	if index >= len(arr) || arr[index] == nil {
		return nil
	}

	root := &TreeNode{Val: arr[index].(int)}
	root.Left = buildTreeFromArray(arr, 2*index+1)
	root.Right = buildTreeFromArray(arr, 2*index+2)

	return root
}
