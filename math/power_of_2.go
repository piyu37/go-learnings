package main

import (
	"fmt"
	"math"
)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	return math.Ceil(math.Log2(float64(n))) == math.Floor(math.Log2(float64(n)))
}

// -n = ~n + 1 (2's complement)
// hence when we do n bitwise (-n) for any number which is power of 2, it again gives n
// but for others it doesn't
// n = 4 (00000100) ~n(11111011)+1=11111100 => n & (~n+1)=00000100 = 4
func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}

	return n&(-n) == n
}

// To subtract 1 means to change the rightmost 1-bit to 0 and to set all the lower bits to 1.
// e.g. 8(00001000) => 7(00000111)
// hence when we do 8 & 7 => it will give 0
// whereas for other which is not power of 2; it will not result in 0
func isPowerOfTwo3(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}

// https://leetcode.com/problems/power-of-two/description/
func powerOf2() {
	fmt.Println(isPowerOfTwo(64))
	fmt.Println(isPowerOfTwo2(64))
	fmt.Println(isPowerOfTwo3(64))
}
