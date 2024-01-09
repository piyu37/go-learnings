package main

import "fmt"

func getNum(row, col int) int {
	if row == 0 || col == 0 || row == col {
		return 1
	}

	return getNum(row-1, col-1) + getNum(row-1, col)
}

func getRow(rowIndex int) []int {
	var ans []int

	for i := 0; i <= rowIndex; i++ {
		ans = append(ans, getNum(rowIndex, i))
	}

	return ans
}

func getRowMath(n int) []int {
	ans := []int{1}

	for k := 1; k <= n; k++ {
		right := (int64(n - k + 1))
		op := int(int64(ans[len(ans)-1]) * right)
		result := op / k
		ans = append(ans, result)
	}

	return ans
}

// https://leetcode.com/problems/pascals-triangle/
// https://leetcode.com/problems/pascals-triangle-ii/
func pascalTriangle() {
	n := 3 // Adjust n as needed
	result := getRow(n)
	fmt.Println(result)
	result = getRowMath(n)
	fmt.Println(result)
}
