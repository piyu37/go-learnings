package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	courseGraph := make(map[int][]int)

	for i := range prerequisites {
		edge := prerequisites[i]

		courseGraph[edge[1]] = append(courseGraph[edge[1]], edge[0])
	}

	visited := make([]bool, numCourses)
	inStack := make([]bool, numCourses)
	result := make([]int, 0)

	for i := 0; i < numCourses; i++ {
		if checkCycleDFS2(i, courseGraph, inStack, visited, &result) {
			return []int{}
		}
	}

	i := 0
	j := len(result) - 1
	for i < j {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}

	return result
}

func checkCycleDFS2(node int, courseGraph map[int][]int, inStack, visited []bool, result *[]int) bool {
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
		if checkCycleDFS2(nie, courseGraph, inStack, visited, result) {
			return true
		}
	}

	*result = append(*result, node)

	inStack[node] = false

	return false
}

// https://leetcode.com/problems/course-schedule-ii/description/?envType=study-plan-v2&envId=top-interview-150
func courseSchedule2() {
	numCourse := 4
	prequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}

	fmt.Println(findOrder(numCourse, prequisites))

	numCourse = 4
	prequisites = [][]int{{1, 0}, {3, 1}, {3, 2}}

	fmt.Println(findOrder(numCourse, prequisites))

	numCourse = 4
	prequisites = [][]int{{3, 2}, {1, 0}, {2, 1}}

	fmt.Println(findOrder(numCourse, prequisites))
}
