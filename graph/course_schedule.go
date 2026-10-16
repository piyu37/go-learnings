package main

import "fmt"

// DFS cycle detection.
// Time: O(V + E) - each course and each prerequisite edge is visited once.
// Space: O(V + E) - adjacency list O(V + E), state O(V), recursion stack O(V) in the worst case.
// Same complexity as canFinishKahn, but a very long prerequisite chain means deep recursion.
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	state := make([]int, numCourses) // 0 = unvisited, 1 = in progress, 2 = done

	var hasCycle func(int) bool
	hasCycle = func(node int) bool {
		if state[node] != 0 {
			return state[node] == 1
		}

		state[node] = 1
		for _, next := range graph[node] {
			if hasCycle(next) {
				return true
			}
		}
		state[node] = 2

		return false
	}

	for i := range numCourses {
		if hasCycle(i) {
			return false
		}
	}

	return true
}

// Kahn's algorithm: repeatedly take courses with no remaining prerequisites.
// If every course gets taken, there is no cycle.
// Time: O(V + E) - each course is queued once and each edge is relaxed once.
// Space: O(V + E) - adjacency list O(V + E), inDegree O(V), queue O(V).
// Same complexity as canFinish, but iterative, so no recursion-depth risk on long chains.
func canFinishKahn(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		inDegree[p[0]]++
	}

	queue := []int{}
	for i := range numCourses {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	taken := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		taken++

		for _, next := range graph[node] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return taken == numCourses
}

// https://leetcode.com/problems/course-schedule/description/?envType=study-plan-v2&envId=top-interview-150
func courseSchedule() {
	numCourse := 2
	prequisites := [][]int{{1, 0}}

	fmt.Println(canFinish(numCourse, prequisites))
	fmt.Println(canFinishKahn(numCourse, prequisites))

	numCourse = 2
	prequisites = [][]int{{1, 0}, {0, 1}}

	fmt.Println(canFinish(numCourse, prequisites))
	fmt.Println(canFinishKahn(numCourse, prequisites))
}
