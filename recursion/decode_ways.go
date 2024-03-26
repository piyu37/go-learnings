package main

import (
	"fmt"
	"strconv"
)

var numberAlbhabetMap = map[string]string{
	"1": "A", "2": "B", "3": "C", "4": "D", "5": "E", "6": "F", "7": "G", "8": "H", "9": "I",
	"10": "J", "11": "K", "12": "L", "13": "M", "14": "N", "15": "O", "16": "P", "17": "Q",
	"18": "R", "19": "S", "20": "T", "21": "U", "22": "V", "23": "W", "24": "X", "25": "Y", "26": "Z",
}

// https://www.geeksforgeeks.org/print-all-possible-decodings-of-a-given-digit-sequence/
func generateAllnumDecodings(s string) []string {
	result := make([]string, 0)

	generateDecodedStrings(0, s, "", &result)

	return result
}

func generateDecodedStrings(idx int, s, tmp string, result *[]string) {
	if idx == len(s) {
		*result = append(*result, tmp)
		return
	}

	num := ""
	for i := idx; i < idx+2 && i < len(s); i++ {
		num += string(s[i])
		char, ok := numberAlbhabetMap[num]
		if ok {
			generateDecodedStrings(i+1, s, tmp+char, result)
		}
	}
}

// https://leetcode.com/problems/decode-ways/
func numDecodings(s string) int {
	memo := make(map[int]int)

	return recursiveWithMemo(0, s, memo)
}

func recursiveWithMemo(idx int, s string, memo map[int]int) int {
	if v, ok := memo[idx]; ok {
		return v
	}

	if idx == len(s) {
		return 1
	}

	if s[idx] == '0' {
		return 0
	}

	if idx == len(s)-1 {
		return 1
	}

	ans := recursiveWithMemo(idx+1, s, memo)
	doubleChar, _ := strconv.Atoi(s[idx : idx+2])
	if doubleChar <= 26 {
		ans += recursiveWithMemo(idx+2, s, memo)
	}

	memo[idx] = ans

	return ans
}

func decodeWays() {
	num := "1123"
	fmt.Println(generateAllnumDecodings(num))
	fmt.Println(numDecodings(num))
	num = "11106"
	fmt.Println(generateAllnumDecodings(num))
	fmt.Println(numDecodings(num))
	num = "101"
	fmt.Println(generateAllnumDecodings(num))
}
