package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(rotateArrayRight(arr))
	//-----------------------------------
	piles := []int{30, 11, 23, 4, 20}
	h := 10
	fmt.Println(minEatingSpeed(piles, h))
	// -----------------------------------
	a := []int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}
	fmt.Println(Solution(a))
	// -----------------------------------
	nums := []int{0, 4, 3, 2, 7, 8, 2, 3, 1}
	duplicates := findDuplicates(nums)
	fmt.Println("The repeating elements are:")
	for _, num := range duplicates {
		fmt.Println(num)
	}
	//-----------------------------------
	fmt.Println(BinarySearch([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 0))
	//-----------------------------------
	input := []interface{}{0, 2, []interface{}{[]interface{}{2, 3}, 8, []interface{}{[]interface{}{100}}, interface{}(nil), []interface{}{[]interface{}{interface{}(nil)}}}, -2}
	fmt.Println(Flatten(input))
	//-----------------------------------
	var n int32 = 100
	qt := []string{"Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total"}
	s1 := []int32{30, 49, 33, 79, 61, 13, 73, 1, 23, 14, 22, 1, 38, 82, 80, 75, 45, 64, 3, 19, 67, 68, 40, 10, 53, 37, 2, 47, 75, 49, 83, 63, 43, 8, 70, 37, 44, 41, 3, 34, 36, 23, 75, 13, 29, 53, 11, 60, 55, 22, 4, 48, 44, 21, 82, 69, 18, 77, 64, 19, 36, 61, 77, 65, 69, 4, 21, 78, 59, 43, 32, 30, 84, 70, 77, 17, 28, 21, 60, 67, 5, 44, 12, 57}
	s2 := []int32{62, 19, 5, 83, 56, 78, 18, 35, 72, 33, 21, 8, 54, 6, 57, 46, 40, 10, 61, 51, 42, 51, 12, 47, 52, 67, 26, 23, 17, 52, 3, 81, 8, 68, 41, 52, 83, 48, 31, 62, 44, 38, 1, 71, 56, 33, 8, 37, 24, 26, 57, 20, 22, 14, 51, 40, 47, 20, 23, 27, 20, 13, 32, 78, 82, 51, 77, 65, 52, 5, 15, 69, 69, 7, 58, 20, 12, 28, 31, 76, 82, 64, 4, 29}
	fmt.Println(getTheGroups(n, qt, s1, s2))
	//-----------------------------------
	ar := []int{-12, 1, 2, 3, 12}
	fmt.Println(IndexEqualsValue(ar))
	//-----------------------------------
	N := 100
	K := 100
	ar = []int{60, 83, 82, 16, 17, 7, 89, 6, 83, 100, 85, 41, 72, 44, 23, 28, 64, 84, 3, 23, 33, 52, 93,
		30, 81, 38, 67, 25, 26, 97, 94, 78, 41, 74, 74, 17, 53, 51, 54, 17, 20, 81, 95, 76, 42, 16, 16, 56,
		74, 69, 30, 9, 82, 91, 32, 13, 47, 45, 97, 40, 56, 57, 27, 28, 84, 98, 91, 5, 61, 20, 3, 43, 42, 26,
		83, 40, 34, 100, 5, 63, 62, 61, 72, 5, 32, 58, 93, 79, 7, 18, 50, 43, 17, 24, 77, 73, 87, 74, 98, 2}
	var out_ int = Max(N, K, ar)
	fmt.Println(out_)
	//-----------------------------------
	ar = []int{1, 2, 3, -4, -5, -6, -7}
	fmt.Println("Original array:", ar)
	moveNegativeElements(ar)
	fmt.Println("Modified array:", ar)
	//-----------------------------------
	arr = [][]int{
		{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95},
		{-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26},
	}
	fmt.Println(eraseOverlapIntervals(arr))
	//-----------------------------------
	pascalTriangle()
	//-----------------------------------
	permutation()
	//-----------------------------------
	arr2D := [][]int{
		{2, 1, 1},
		{1, 2, 1},
	}

	fmt.Println(findMinPath(2, 3, 0, 0, arr2D, -1))
	//-----------------------------------
	arr = [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{0, 4},
		{-2, 6},
		{4, 0},
		{2, 1},
	}

	removeDuplicatesMain()
	//-----------------------------------
	fmt.Println(LineThroughPoints(arr))
	//-----------------------------------
	matrix := [][]int{
		{1, 4, 7, 12, 15, 1000},
		{2, 5, 19, 31, 32, 1001},
		{3, 8, 24, 33, 35, 1002},
		{40, 41, 42, 44, 45, 1003},
		{99, 100, 103, 106, 128, 1004},
	}
	fmt.Println(SearchInSortedMatrix(matrix, 4))
	//-----------------------------------
	stockMaxProfit()
	//-----------------------------------
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "WEST"
	fmt.Println(SunsetViews(buildings, direction))
	//-----------------------------------
	N = 6
	T := 7
	A := []int{2, 3, 2, 3, 3, 1}
	fmt.Println(taskManagement(N, T, A))
	//-----------------------------------
	zigzag()
}
