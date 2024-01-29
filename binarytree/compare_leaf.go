package main

import "fmt"

type B1Info struct {
	head, prev *BinaryTree
}

func compareLeaf(b1, b2 *BinaryTree) bool {
	if b1 == nil && b2 != nil {
		return false
	}

	if b1 != nil && b2 == nil {
		return false
	}

	if b1 == nil && b2 == nil {
		return true
	}

	b1Info := &B1Info{}
	b1Info.preOrder1(b1)

	b2Info := &B1Info{}
	b2Info.preOrder1(b2)

	b1Node := b1Info.head
	b2Node := b2Info.head

	for b1Node != nil && b2Node != nil {
		if b1Node.Value != b2Node.Value {
			return false
		}

		b1Node = b1Node.Right
		b2Node = b2Node.Right
	}

	if (b1Node == nil && b2Node != nil) || (b2Node == nil && b1Node != nil) {
		return false
	}

	return true
}

func (b1Info *B1Info) preOrder1(b1 *BinaryTree) {
	if b1.Left != nil {
		b1Info.preOrder1(b1.Left)
	}

	if b1.Right != nil {
		b1Info.preOrder1(b1.Right)
		return
	}

	if b1Info.head == nil {
		b1Info.head = b1
	}

	if b1Info.prev != nil {
		b1Info.prev.Right = b1
		b1Info.prev = b1Info.prev.Right
	} else {
		b1Info.prev = b1
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/11_compare_leaf_traversal
func compareLeafMain() {
	bt1 := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
			},
			Right: &BinaryTree{
				Value: 5,
				Left: &BinaryTree{
					Value: 7,
				},
				Right: &BinaryTree{
					Value: 8,
				},
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Right: &BinaryTree{
				Value: 6,
			},
		},
	}

	bt2 := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Right: &BinaryTree{
				Value: 5,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 6,
				},
			},
		},
	}

	fmt.Println(compareLeaf(bt1, bt2))
}
