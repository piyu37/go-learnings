package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	nodeMap := make(map[*Node]bool)

	connectLeftToRight(root.Left, root, nodeMap)
	connectLeftToRight(root.Right, root, nodeMap)

	return root
}

func connectLeftToRight(root, parent *Node, nodeMap map[*Node]bool) {
	if root == nil || nodeMap[root] {
		return
	}

	nodeMap[root] = true

	if parent.Right != nil && parent.Right != root {
		root.Next = parent.Right
	} else {
		parentNext := parent.Next
		for parentNext != nil && parentNext.Left == nil && parentNext.Right == nil {
			parentNext = parentNext.Next
		}

		if parentNext != nil {
			if parentNext.Left != nil {
				root.Next = parentNext.Left
			} else {
				root.Next = parentNext.Right
			}
		}
	}

	connectLeftToRight(root.Next, parent, nodeMap)
	connectLeftToRight(root.Left, root, nodeMap)
	connectLeftToRight(root.Right, root, nodeMap)
}

func connectUsingQueue(root *Node) *Node {
	if root == nil {
		return root
	}

	nq := make([]*Node, 0)

	nq = append(nq, root)

	for len(nq) > 0 {
		size := len(nq)
		for i := 0; i < size; i++ {
			node := nq[0]
			nq = nq[1:]
			if i < size-1 {
				node.Next = nq[0]
			}

			if node.Left != nil {
				nq = append(nq, node.Left)
			}
			if node.Right != nil {
				nq = append(nq, node.Right)
			}
		}
	}

	return root
}

func connectWithoutQueue(root *Node) *Node {
	if root == nil {
		return root
	}

	head := root

	for head != nil {
		dummy := &Node{}
		temp := dummy

		for head != nil {
			if head.Left != nil {
				temp.Next = head.Left
				temp = temp.Next
			}

			if head.Right != nil {
				temp.Next = head.Right
				temp = temp.Next
			}

			head = head.Next
		}

		head = dummy.Next
	}

	return root
}

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii/description/?envType=study-plan-v2&envId=top-interview-150
func rightSiblingInNextPointer() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 7,
			},
		},
	}
	connect(root)
	connectUsingQueue(root)
	connectWithoutQueue(root)
}
