package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N)
 * ----------------------------
 * Where N is the total number of nodes in the Binary Tree.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 * - The algorithm performs a standard Depth-First Search (DFS) traversal
 * of the binary tree.
 * - It visits every single node in the tree exactly once to accumulate the
 * depths and evaluate its left and right children.
 * - Because no node is evaluated more than once, there are exactly N unique
 * states (or nodes) that need to be processed.
 *
 * 2. Work Per State -> O(1)
 * - Inside each recursive call, the algorithm performs a few basic,
 * constant-time operations: a nil check (`bt == nil`), integer additions
 * (`depthSum += depth` and the final return addition), and basic variable
 * assignments.
 * - Regardless of how large the tree gets, the processing overhead at any
 * single individual node takes O(1) constant time.
 *
 * Total Time = O(N) nodes visited * O(1) work per node = O(N)
 *
 * SPACE COMPLEXITY: O(H)
 * ----------------------------
 * Where H is the maximum height of the Binary Tree.
 *
 * 1. Recursion Stack: The algorithm explores the tree branch by branch. The
 * maximum depth of the recursion tree corresponds exactly to the height
 * of the longest branch in the binary tree.
 * - In a relatively balanced tree, this height is O(log N).
 * - In the absolute worst-case scenario (a highly skewed tree where every
 * node only has one child, acting like a linked list), the height is O(N).
 * - Stating O(H) is the most theoretically accurate and accepted answer.
 * 2. Auxiliary Space: The algorithm only uses primitive integer variables
 * (`depthSum`, `depth`, `left`, `right`) passed by value within each stack
 * frame. It does not allocate any new arrays, slices, or hash maps.
 * This requires O(1) extra space.
 *
 * Total Space = Stack O(H) + Auxiliary O(1) = O(H)
 */
func calculateSumOfSubtrees(bt *BinaryTree, depthSum, depth int) int {
	if bt == nil {
		return 0
	}

	depthSum += depth

	left := calculateSumOfSubtrees(bt.Left, depthSum, depth+1)
	right := calculateSumOfSubtrees(bt.Right, depthSum, depth+1)

	return depthSum + left + right
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/06_node_depths
func allKindOfNodeDepth() {
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

	fmt.Println(calculateSumOfSubtrees(bt, 0, 0))
}
