package main

import "fmt"

func dfs(start string, graph map[string][]string, output []string) []string {
	auxList := []string{start}

	for len(auxList) > 0 {
		v := auxList[0]
		children := graph[v]

		temp := auxList[1:]

		auxList = children

		auxList = append(auxList, temp...)

		output = append(output, v)
	}

	return output
}

func dfs2(start string, graph map[string][]string, array []string) []string {
	array = append(array, start)
	for _, chidNode := range graph[start] {
		array = dfs2(chidNode, graph, array)
	}

	return array
}

func dfsMain() {
	graph := map[string][]string{
		"A": {"B", "C", "D"},
		"B": {"E", "F"},
		"F": {"I", "J"},
		"D": {"G", "H"},
		"G": {"K"},
	}

	fmt.Println(dfs("A", graph, []string{}))
	fmt.Println(dfs2("A", graph, []string{}))
}
