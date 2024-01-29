package main

func findNode(bst *BST, v int) (*BST, *BST) {
	var parent *BST

	for bst != nil {
		if bst.Value == v {
			return parent, bst
		}

		parent = bst
		if v < bst.Value {
			bst = bst.Left
		} else {
			bst = bst.Right
		}
	}

	return parent, nil
}

func findNextGreaterNode(bst *BST) (*BST, *BST) {
	var parent *BST

	for bst.Left != nil {
		parent = bst
		bst = bst.Left
	}

	return parent, bst
}

func delete(bst *BST, v int) {
	root := bst
	_, toBeDeletedNode := findNode(bst, v)

	if toBeDeletedNode.Left == nil && toBeDeletedNode.Right == nil {
		// leaf node deletion
	} else if toBeDeletedNode.Left == nil || toBeDeletedNode.Right == nil {
		// 1 child deletion
	} else {
		parent, nextGreaterNode := findNextGreaterNode(toBeDeletedNode.Right)
		parent.Left = nextGreaterNode.Right
		toBeDeletedNode.Value = nextGreaterNode.Value
		nextGreaterNode = nil
	}

	println(root, toBeDeletedNode)
}

// program to delete node from BST
func deletionMain() {
	bst := &BST{
		Value: 10,
		Left: &BST{
			Value: 5,
			Left: &BST{
				Value: 2,
				Left: &BST{
					Value: 1,
				},
			},
			Right: &BST{
				Value: 5,
			},
		},
		Right: &BST{
			Value: 15,
			Left: &BST{
				Value: 13,
				Left: &BST{
					Value: 11,
					Right: &BST{
						Value: 12,
					},
				},
				Right: &BST{
					Value: 14,
				},
			},
			Right: &BST{
				Value: 22,
			},
		},
	}

	delete(bst, 10)
}
