package main

import "fmt"

func mapParentNode(node, parentNode *BinaryTreeWithParent, target int, targetNode *BinaryTreeWithParent) {
	if node == nil {
		return
	}

	if node.Value == target {
		*targetNode = *node
		targetNode.Parent = parentNode
	}

	node.Parent = parentNode
	parentNode = node

	mapParentNode(node.Left, parentNode, target, targetNode)

	mapParentNode(node.Right, parentNode, target, targetNode)
}

func FindNodesDistanceK(tree *BinaryTreeWithParent, target int, k int) []int {
	root := tree

	var targetNode BinaryTreeWithParent

	mapParentNode(root, nil, target, &targetNode)

	result := make([]int, 0)

	findKDistanceNodes(&targetNode, 0, &result, k)
	var side string = "left"
	if targetNode.Parent != nil && targetNode.Parent.Right != nil &&
		targetNode.Parent.Right.Value == targetNode.Value {
		side = "right"
	}
	findParentKDistanceNodes(targetNode.Parent, 1, &result, k, side)

	return result
}

func findParentKDistanceNodes(tree *BinaryTreeWithParent, counter int, result *[]int, k int, side string) {
	if tree == nil {
		return
	}

	if counter == k {

		*result = append(*result, tree.Value)

		return
	}

	if tree.Left != nil && side != "left" {
		findKDistanceNodes(tree.Left, counter+1, result, k)
	}

	if tree.Right != nil && side != "right" {
		findKDistanceNodes(tree.Right, counter+1, result, k)
	}

	findParentKDistanceNodes(tree.Parent, counter+1, result, k, side)
}

func findKDistanceNodes(tree *BinaryTreeWithParent, counter int, result *[]int, k int) {
	if tree == nil {
		return
	}

	if counter == k {
		*result = append(*result, tree.Value)

		return
	}

	findKDistanceNodes(tree.Left, counter+1, result, k)
	findKDistanceNodes(tree.Right, counter+1, result, k)
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	graph := make(map[int][]int)

	buildGraph(root, nil, graph)

	visited := make(map[int]bool)
	visited[target.Val] = true

	result := make([]int, 0)

	dfsTree(graph, target.Val, 0, k, visited, &result)

	return result
}

func buildGraph(curr, parent *TreeNode, graph map[int][]int) {
	if curr != nil && parent != nil {
		graph[curr.Val] = append(graph[curr.Val], parent.Val)
		graph[parent.Val] = append(graph[parent.Val], curr.Val)
	}

	if curr.Left != nil {
		buildGraph(curr.Left, curr, graph)
	}

	if curr.Right != nil {
		buildGraph(curr.Right, curr, graph)
	}
}

func dfsTree(graph map[int][]int, target, distance, k int, visited map[int]bool, result *[]int) {
	if distance == k {
		*result = append(*result, target)
		return
	}

	for _, n := range graph[target] {
		if !visited[n] {
			visited[n] = true
			dfsTree(graph, n, distance+1, k, visited, result)
		}
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/09_find_nodes_distance_k
func findNodeDistanceKMain() {
	target := &BinaryTreeWithParent{
		Value: 3,
		Right: &BinaryTreeWithParent{
			Value: 6,
			Left: &BinaryTreeWithParent{
				Value: 7,
			},
			Right: &BinaryTreeWithParent{
				Value: 8,
			},
		},
	}

	tree := &BinaryTreeWithParent{
		Value: 1,
		Left: &BinaryTreeWithParent{
			Value: 2,
			Left: &BinaryTreeWithParent{
				Value: 4,
			},
			Right: &BinaryTreeWithParent{
				Value: 5,
			},
		},
		Right: target,
	}

	fmt.Println(FindNodesDistanceK(tree, target.Value, 2))

	tree1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}

	fmt.Println(distanceK(tree1, tree1.Left, 2))
}
