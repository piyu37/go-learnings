package main

import "fmt"

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

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
 * - The algorithm performs a Depth-First Search (DFS) traversal to find
 * all root-to-leaf paths.
 * - It visits every single node in the binary tree exactly once, stopping
 * only when it hits the nil children of leaf nodes.
 * - Since no node is evaluated more than once, there are exactly N unique
 * states (or nodes) that need to be processed.
 *
 * 2. Work Per State -> O(1)
 * - Inside each recursive call, the algorithm performs constant-time
 * operations: nil checks, evaluating if a node is a leaf, and basic integer
 * addition (`count + root.Value`).
 * - At leaf nodes, it appends the sum to the `result` slice. In Go, slice
 * appending operates in O(1) amortized time.
 * - Because these operations do not scale with the size of the tree, the
 * work done per node is O(1).
 *
 * Total Time = O(N) nodes visited * O(1) work per node = O(N)
 *
 * SPACE COMPLEXITY: O(H)
 * ----------------------------
 * Where H is the maximum height of the Binary Tree.
 *
 * 1. Recursion Stack: The algorithm explores the tree branch by branch down
 * to the leaf nodes. The maximum depth of the call stack corresponds to the
 * longest root-to-leaf path in the tree.
 * - In a perfectly balanced tree, this takes O(log N) space.
 * - In the worst-case scenario (a highly skewed tree resembling a linked list),
 * it takes O(N) space.
 * - O(H) is the most accurate and standard answer.
 * 2. Auxiliary Space: Standard space complexity analysis excludes the memory
 * required to hold the final output array (which would take O(L) space, where
 * L is the number of leaf nodes, bounded by O(N)). Excluding the `result`
 * slice, the algorithm only uses primitive integer variables and pointers
 * within each stack frame, allocating no extra dynamic data structures.
 * This requires O(1) auxiliary space.
 *
 * Total Space = Stack O(H) + Auxiliary O(1) (excluding output) = O(H)
 */
func BranchSums(root *BinaryTree) []int {

	result := make([]int, 0)

	calculateSum(root, 0, &result)

	return result
}

func calculateSum(root *BinaryTree, count int, result *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, count+root.Value)

		return
	}

	calculateSum(root.Left, count+root.Value, result)

	calculateSum(root.Right, count+root.Value, result)
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/05_branch_sums
func branchSumMain() {
	tree := &BinaryTree{
		Value: 0,
		Right: &BinaryTree{
			Value: 1,
			Right: &BinaryTree{
				Value: 10,
				Right: &BinaryTree{
					Value: 100,
				},
			},
		},
	}

	fmt.Println(BranchSums(tree))
}
