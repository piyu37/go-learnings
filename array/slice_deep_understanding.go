package main

import "fmt"

func sliceUnderstanding() {
	arr := [5]string{"A", "B", "C", "D"}
	// len tells no. of elemnts in slice
	// cap tells block size of slice in memory
	fmt.Println(arr, len(arr), cap(arr))
	// can't give end index more than len of array
	sl := arr[2:5]
	fmt.Println(sl, len(sl), cap(sl))
	// sl = append(sl, "E")
	// fmt.Println(sl, len(sl), cap(sl))
	check(sl)
	fmt.Println(sl, len(sl), cap(sl))
}

func check(sl []string) {
	sl[0] = "F"
	fmt.Println(sl, len(sl), cap(sl))
	sl = append(sl, "G")
	fmt.Println(sl, len(sl), cap(sl))
	sl[1] = "H"
	fmt.Println(sl, len(sl), cap(sl))
}
