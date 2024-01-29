package main

func RightSiblingTree(root *BinaryTree) *BinaryTree {
	mutate(root, nil, false)

	return root
}

func mutate(node, parent *BinaryTree, isLeftNode bool) {
	if node == nil {
		return
	}

	left, right := node.Left, node.Right

	mutate(left, node, true)
	if parent == nil {
		node.Right = nil
	} else if isLeftNode {
		node.Right = parent.Right
	} else {
		if parent.Right == nil {
			node.Right = nil
		} else {
			node.Right = parent.Right.Left
		}
	}

	mutate(right, node, false)
}

// https://github.com/lee-hen/Algoexpert/tree/master/very_hard/09_right_sibling_tree
func rightSiblingTreeMain() {
	bt := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
				Right: &BinaryTree{
					Value: 10,
				},
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
				Left: &BinaryTree{
					Value: 11,
					Left: &BinaryTree{
						Value: 14,
					},
				},
			},
			Right: &BinaryTree{
				Value: 7,
				Left: &BinaryTree{
					Value: 12,
				},
				Right: &BinaryTree{
					Value: 13,
				},
			},
		},
	}

	RightSiblingTree(bt)
}
