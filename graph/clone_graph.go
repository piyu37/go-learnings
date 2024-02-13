package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraphRecur(node *Node) *Node {
	return createGraphRecur(node, map[int]*Node{})
}

func createGraphRecur(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	clone := &Node{
		Val: node.Val,
	}

	visited[node.Val] = clone

	for i := range node.Neighbors {
		if neighbor, ok := visited[node.Neighbors[i].Val]; ok {
			clone.Neighbors = append(clone.Neighbors, neighbor)
			continue
		}

		neighbor := createGraphRecur(node.Neighbors[i], visited)
		clone.Neighbors = append(clone.Neighbors, neighbor)
	}

	return clone
}

// https://leetcode.com/problems/clone-graph/description/?envType=study-plan-v2&envId=top-interview-150
func cloneGraphMain() {
	node1 := &Node{
		Val: 1,
	}
	node2 := &Node{
		Val: 2,
	}
	node3 := &Node{
		Val: 3,
	}
	node4 := &Node{
		Val: 4,
	}

	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)

	fmt.Println(cloneGraphRecur(node1))
}
