package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BT struct {
	root *Node
}

func binaryTree() {
	a := []int{10, 11, 9, 7, 8}
	bt := BT{}

	for _, n := range a {
		bt.Insert(n)
	}

	bt.PreOrderPrint(bt.root)
}

func (bt *BT) Insert(value int) {
	if bt.root == nil {
		bt.root = &Node{data: value}
	} else {
		arr := []*Node{bt.root}

		for len(arr) > 0 {
			temp := arr[0]
			arr = arr[1:]

			if temp.left == nil {
				temp.left = &Node{data: value}
				break
			} else {
				arr = append(arr, temp.left)
			}

			if temp.right == nil {
				temp.right = &Node{data: value}
				break
			} else {
				arr = append(arr, temp.right)
			}
		}
	}
}

func (bt *BT) PreOrderPrint(n *Node) {
	if n != nil {
		fmt.Println(n.data)
		bt.PreOrderPrint(n.left)
		bt.PreOrderPrint(n.right)
	}
}
