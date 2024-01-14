// Given 2 epoch timestamps 1610819400 and 1610945400
// Find out how many days lie between these 2 timestamps

package main

import "fmt"

func main() {
	t1 := 1610819400

	t2 := 1610945400
	timeDiff := t2 - t1
	// fmt.Println(timeDiff)
	t := timeDiff / (24 * 60 * 60)
	fmt.Println(t)
}
