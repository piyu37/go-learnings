package main

import (
	"fmt"
	"math"
)

func taskManagement(N, T int, input []int) int {
	minNum := input[0]

	for i := 1; i < N; i++ {
		if input[i] < minNum {
			minNum = input[i]
		}
	}

	maxNum := N
	var k int
	minK := math.MaxInt16
	flag := false

	for minNum <= maxNum {
		k = minNum + (maxNum-minNum)/2
		tasksCache := 0
		inputCounter := 0
		flag = false
		arr := make([]int, T+1)
		arr[0] = -1

		for i := 0; i <= T; i++ {
			if arr[i] > 0 {
				tasksCache -= arr[i]
				arr[i] = 0
			}

			for tasksCache < k && inputCounter < N {
				if i+input[inputCounter] > T {
					minNum = k + 1
					flag = true
					break
				}

				arr[i+input[inputCounter]] += 1
				tasksCache++
				inputCounter++
			}

			if flag {
				break
			}
		}

		if !flag {
			if k < minK {
				minK = k
			}

			maxNum = k - 1
		}
	}

	return minK
}

// WAP to find minimum cache window/size to complete N tasks in given T time.
// A[i] represents time taken to complete each task.
func taskManagementMain() {
	N := 6
	T := 7
	A := []int{2, 3, 2, 3, 3, 1}
	fmt.Println(taskManagement(N, T, A))
}
