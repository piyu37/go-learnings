package main

import "fmt"

func main() {
	arrayRotate()
	//-----------------------------------
	bananaEating()
	// -----------------------------------
	a := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println(Solution(a))
	// -----------------------------------
	duplicateNumber()
	//-----------------------------------
	fmt.Println(BinarySearch([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 0))
	//-----------------------------------
	input := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, []interface{}{[]interface{}{100}}, interface{}(nil), []interface{}{[]interface{}{interface{}(nil)}}}, -2}
	fmt.Println(Flatten(input))
	//-----------------------------------
	getTheGroupsMain()
	//-----------------------------------
	ar := []int{-12, 1, 2, 3, 12}
	fmt.Println(IndexEqualsValue(ar))
	//-----------------------------------
	lineThroughPointsMain()
	//-----------------------------------
	moveNegativeToEnd()
	//-----------------------------------
	nonOverlapingIntervals()
	//-----------------------------------
	pascalTriangle()
	//-----------------------------------
	reachTheEnd()
	//-----------------------------------
	removeDuplicatesMain()
	//-----------------------------------
	sortedMatrixSearch()
	//-----------------------------------
	stockMaxProfit()
	//-----------------------------------
	sunset()
	//-----------------------------------
	taskManagementMain()
	//-----------------------------------
	zigzag()
}
