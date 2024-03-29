package main

import "fmt"

type GraphNode struct {
	Name     string
	Children []*GraphNode
}

func (n *GraphNode) AddChildren(names ...string) *GraphNode {
	for _, name := range names {
		child := GraphNode{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func (n *GraphNode) bfs() []string {
	result := make([]string, 0)

	stack := make([]*GraphNode, 0)

	stack = append(stack, n)

	for len(stack) > 0 {
		v := stack[0]

		result = append(result, v.Name)

		stack = stack[1:]

		stack = append(stack, v.Children...)
	}

	return result
}

func bfsMain() {
	g := &GraphNode{
		Name: "A",
	}

	g.AddChildren("B", "C", "D")
	g.Children[0].AddChildren("E", "F")
	g.Children[2].AddChildren("G", "H")
	g.Children[0].Children[1].AddChildren("I", "J")
	g.Children[2].Children[0].AddChildren("K")
	fmt.Println(g.bfs())
}
