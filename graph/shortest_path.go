package main

import (
	"container/list"
	"fmt"
)

// finds shortest path between 2 nodes of a graph using BFS
func bfsShortestPath(graph map[string][]string, start, goal string) interface{} {
	// keep track of explored nodes
	explored := make(map[string]bool)
	// keep track of all the paths to be checked
	queue := list.New()
	queue.PushBack([]string{start})

	// return path if start is goal
	if start == goal {
		return "That was easy! Start = goal"
	}

	// keeps looping until all possible paths have been checked
	for queue.Len() > 0 {
		// pop the first path from the queue
		path := queue.Remove(queue.Front()).([]string)
		// get the last node from the path
		node := path[len(path)-1]

		if !explored[node] {
			neighbours := graph[node]
			// go through all neighbour nodes, construct a new path and
			// push it into the queue
			for _, neighbour := range neighbours {
				newPath := append([]string(nil), path...)
				newPath = append(newPath, neighbour)
				queue.PushBack(newPath)
				// return path if neighbour is goal
				if neighbour == goal {
					return newPath
				}
			}

			// mark node as explored
			explored[node] = true
		}
	}

	// in case there's no path between the 2 nodes
	return "So sorry, but a connecting path doesn't exist :("
}

func addEdge(graph map[string][]string, u, v string) {
	graph[u] = append(graph[u], v)
}

func shortestPath() {
	graph := make(map[string][]string)
	addEdge(graph, "G", "A")
	addEdge(graph, "G", "B")
	addEdge(graph, "G", "C")
	addEdge(graph, "A", "G")
	addEdge(graph, "A", "D")
	addEdge(graph, "A", "E")
	addEdge(graph, "B", "G")
	addEdge(graph, "B", "C")
	addEdge(graph, "C", "G")
	addEdge(graph, "C", "B")
	addEdge(graph, "C", "D")
	addEdge(graph, "D", "A")
	addEdge(graph, "D", "C")
	addEdge(graph, "E", "A")

	result := bfsShortestPath(graph, "G", "D") // returns ["G", "C", "A", "D"]
	fmt.Println(result)
}
