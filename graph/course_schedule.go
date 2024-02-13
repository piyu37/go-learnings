package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	courseGraph := make(map[int][]int)

	for i := range prerequisites {
		edge := prerequisites[i]

		courseGraph[edge[1]] = append(courseGraph[edge[1]], edge[0])
	}

	visited := make([]bool, numCourses)
	inStack := make([]bool, numCourses)

	for key := range courseGraph {
		if checkCycleDFS(key, courseGraph, inStack, visited) {
			return false
		}
	}

	return true
}

func checkCycleDFS(node int, courseGraph map[int][]int, inStack, visited []bool) bool {
	if inStack[node] {
		return true
	}

	if visited[node] {
		return false
	}

	neighbors := courseGraph[node]

	visited[node] = true
	inStack[node] = true

	for _, nie := range neighbors {
		if checkCycleDFS(nie, courseGraph, inStack, visited) {
			return true
		}
	}

	inStack[node] = false

	return false
}

// https://leetcode.com/problems/course-schedule/description/?envType=study-plan-v2&envId=top-interview-150
func courseSchedule() {
	numCourse := 2
	prequisites := [][]int{{1, 0}}

	fmt.Println(canFinish(numCourse, prequisites))

	numCourse = 2
	prequisites = [][]int{{1, 0}, {0, 1}}

	fmt.Println(canFinish(numCourse, prequisites))
}
