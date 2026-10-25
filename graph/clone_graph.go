package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraphRecur(node *Node) *Node {
	visited := map[*Node]*Node{}

	var dfs func(n *Node) *Node
	dfs = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if clone, ok := visited[n]; ok {
			return clone
		}

		clone := &Node{Val: n.Val, Neighbors: make([]*Node, 0, len(n.Neighbors))}
		visited[n] = clone

		for _, nb := range n.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(nb))
		}

		return clone
	}

	return dfs(node)
}

func cloneGraphRecurBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	clone := &Node{
		Val: node.Val,
	}

	visited := make(map[int]*Node)

	visited[node.Val] = clone

	queue := []*Node{node}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nb := range curr.Neighbors {
			if _, ok := visited[nb.Val]; !ok {
				queue = append(queue, nb)
				visited[nb.Val] = &Node{Val: nb.Val}
			}

			visited[curr.Val].Neighbors = append(visited[curr.Val].Neighbors, visited[nb.Val])
		}
	}

	return visited[node.Val]
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
