package main

func findMinPath(row, col, i, j int, matrix [][]int, pathCount int) int {
	currentValue := matrix[i][j]

	if i == row-1 && j == col-1 {
		return pathCount + 1
	}

	nextRowVal, nextColVal := i, j

	var rowTraverse, colTraverse int

	if i+currentValue < row {
		nextRowVal = i + currentValue
		rowTraverse = findMinPath(row, col, nextRowVal, j, matrix, pathCount+1)
	} else {
		rowTraverse = 50000
	}

	if j+currentValue < col {
		nextColVal = j + currentValue
		colTraverse = findMinPath(row, col, i, nextColVal, matrix, pathCount+1)
	} else {
		colTraverse = 50000
	}

	min := min(rowTraverse, colTraverse)

	return min
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}

	return v1
}

// func main() {
// 	var N int

// 	fmt.Scanf("%d", &N)

// 	for i := 0; i < N; i++ {
// 		var row, col int
// 		fmt.Scanf("%d%d", &row, &col)

// 		arr2D := make([][]int, 0)

// 		for r := 0; r < row; r++ {
// 			arr := make([]int, col)
// 			for c := 0; c < col; c++ {
// 				fmt.Scanf("%d", &arr[c])
// 			}

// 			arr2D = append(arr2D, arr)
// 		}

// 		fmt.Println(findMinPath(row, col, 0, 0, arr2D, -1))

// 	}
// }
