package main

func (tree *BinaryTree) InvertBinaryTree() {
	invertLeftAndRight(tree)
}

func invertLeftAndRight(root *BinaryTree) *BinaryTree {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil && root.Right != nil &&
		root.Right.Left == nil && root.Right.Right == nil {
		root.Left, root.Right = root.Right, root.Left

		return root
	}

	if root.Right == nil && root.Left != nil &&
		root.Left.Left == nil && root.Left.Right == nil {
		root.Left, root.Right = root.Right, root.Left

		return root
	}

	if root.Left.Left == nil && root.Left.Right == nil &&
		root.Right.Left == nil && root.Right.Right == nil {
		root.Left, root.Right = root.Right, root.Left

		return root
	}

	left := invertLeftAndRight(root.Left)
	right := invertLeftAndRight(root.Right)

	root.Left, root.Right = right, left

	return root
}

func invertTreeOptimal(root *BinaryTree) *BinaryTree {
	if root == nil {
		return root
	}

	right := root.Right
	root.Right = invertTreeOptimal(root.Left)
	root.Left = invertTreeOptimal(right)

	return root
}

// https://leetcode.com/problems/invert-binary-tree/description/?envType=study-plan-v2&envId=top-interview-150
func invertTree() {
	tree := &BinaryTree{
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
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}

	tree.InvertBinaryTree()
	invertTreeOptimal(tree)
}
