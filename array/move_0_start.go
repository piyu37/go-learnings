package main

import "fmt"

func moveZeroesToStart(input []int) {
	i := 0
	ptr := 0
	for i < len(input) {
		if input[ptr] == 0 {
			i++
			ptr++
			continue
		}

		input = append(input, input[ptr])

		tempArr := input[ptr+1:]

		input = input[:ptr]

		input = append(input, tempArr...)

		i++
	}

	fmt.Print(input)
}

func moveZeroesToStart2(arr []int) {
	n := len(arr)
	writeIdx := n - 1

	// Iterate through the array from left to right
	for i := n - 1; i >= 0; i-- {
		if arr[i] != 0 {
			if writeIdx != i {
				// Swap non-zero element with the element at writeIdx
				arr[writeIdx] = arr[i]
				arr[i] = 0
			}
			writeIdx--
		}
	}
}

// i/p: {3,5,0,4,0,1}
// o/p:{0,0,3,5,4,1}
func move0StartMain() {
	input := []int{3, 5, 0, 4, 0, 1}
	input1 := []int{3, 5, 0, 4, 0, 1}

	moveZeroesToStart(input)
	moveZeroesToStart2(input1)
	fmt.Println(input1)
}
