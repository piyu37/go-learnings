package main

import "fmt"

func rotateArrayRight(arr [][]int) [][]int {
	originalRowStart, originalRowEnd := 0, len(arr)
	originalColStart, originalColEnd := 0, len(arr[0])

	for originalRowStart < originalRowEnd && originalColStart < originalColEnd {
		rowStart, rowEnd := originalRowStart, originalRowEnd
		colStart, colEnd := originalColStart, originalColEnd

		prevValue := arr[rowStart][colStart]
		for colStart < colEnd-1 {
			nextValue := arr[rowStart][colStart+1]
			arr[rowStart][colStart+1] = prevValue
			prevValue = nextValue
			colStart++
		}

		for rowStart < rowEnd-1 {
			nextValue := arr[rowStart+1][colStart]
			arr[rowStart+1][colStart] = prevValue
			prevValue = nextValue
			rowStart++
		}

		colStart, colEnd = originalColEnd-1, originalColStart

		for colStart > colEnd {
			nextValue := arr[rowStart][colStart-1]
			arr[rowStart][colStart-1] = prevValue
			prevValue = nextValue
			colStart--
		}

		rowStart, rowEnd = originalRowEnd-1, originalRowStart

		for rowStart > rowEnd {
			nextValue := arr[rowStart-1][colStart]
			arr[rowStart-1][colStart] = prevValue
			prevValue = nextValue
			rowStart--
		}

		originalRowStart, originalRowEnd = originalRowStart+1, originalRowEnd-1
		originalColStart, originalColEnd = originalColStart+1, originalColEnd-1
	}

	return arr
}

// rotate clockwise array by one in right
// [1 2 3 4			[5 1 2 3
//  5 6 7 8		=>   9 10 6 4
//  9 10 11 12		 13 11 7 8
//  13 14 15 16]	 14 15 16 12]
func arrayRotate() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(rotateArrayRight(arr))
}
