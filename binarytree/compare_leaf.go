package main

import "fmt"

type B1Info struct {
	head, prev *BinaryTree
}

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N1 + N2)
 * ----------------------------
 * Where N1 and N2 are the total number of nodes in binary trees `b1` and `b2`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N1 + N2)
 * - The algorithm fully traverses the first tree (visiting its N1 nodes) and
 * then fully traverses the second tree (visiting its N2 nodes) using the
 * recursive `preOrder1` function to connect the target nodes.
 * - After building the linked lists out of the leaves, it iterates through
 * them using a `for` loop. The maximum length of these linked lists is bounded
 * by the number of leaf nodes, which is strictly less than or equal to N1 and N2.
 * - Because each node across both trees and lists is visited a constant number
 * of times, the total states processed scales linearly with N1 + N2.
 *
 * 2. Work Per State -> O(1)
 * - Inside the recursive `preOrder1` traversal, the algorithm performs basic,
 * constant-time operations: nil checks and pointer reassignments
 * (`b1Info.prev.Right = b1`) to weave the nodes into an in-place linked list.
 * - During the `for` loop comparison, checking values (`b1Node.Value != b2Node.Value`)
 * and advancing the pointers (`b1Node = b1Node.Right`) also takes exactly
 * O(1) constant time per iteration.
 *
 * Total Time = O(N1 + N2) states * O(1) work per state = O(N1 + N2)
 *
 * SPACE COMPLEXITY: O(H1 + H2)
 * ----------------------------
 * Where H1 and H2 are the maximum heights of binary trees `b1` and `b2`.
 *
 * 1. Recursion Stack: The algorithm performs a depth-first search on both trees
 * (sequentially, not simultaneously). The maximum depth of the call stack for
 * the first tree is H1, and for the second tree is H2. In the worst-case
 * scenario (highly skewed trees), this consumes O(H1 + H2) space. If the
 * trees are perfectly balanced, this would be O(log N1 + log N2).
 * 2. Auxiliary Space: The algorithm cleverly reuses the existing `Right` pointers
 * of the tree's leaf nodes to form the linked list in-place. It only allocates
 * two minor `B1Info` struct tracking pointers. Because no dynamically sized
 * data structures (like slices or arrays) are allocated to store the leaves,
 * this requires O(1) extra space.
 *
 * Total Space = Stack O(H1 + H2) + Auxiliary O(1) = O(H1 + H2)
 */
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
