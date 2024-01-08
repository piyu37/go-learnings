package main

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
