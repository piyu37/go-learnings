package main

import (
	"fmt"
	"strconv"
	"strings"
)

var numberAlbhabetMap = map[string]string{
	"1": "A", "2": "B", "3": "C", "4": "D", "5": "E", "6": "F", "7": "G", "8": "H", "9": "I",
	"10": "J", "11": "K", "12": "L", "13": "M", "14": "N", "15": "O", "16": "P", "17": "Q",
	"18": "R", "19": "S", "20": "T", "21": "U", "22": "V", "23": "W", "24": "X", "25": "Y", "26": "Z",
}

// https://www.geeksforgeeks.org/print-all-possible-decodings-of-a-given-digit-sequence/
/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY:
 * Tight & Accurate Bound: O(N * 1.618^N) (or O(N * Fibonacci(N)))
 * Standard Upper Bound: O(N * 2^N)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^N)
 *    - At every index in the string, the algorithm makes up to 2 choices: decode
 *      the next 1 character, or decode the next 2 characters.
 *    - In the absolute worst-case scenario (e.g., a string of all "1"s like "11111"),
 *      both the 1-char and 2-char branches are always valid. The number of leaf nodes
 *      generated perfectly matches the Fibonacci sequence, which grows exponentially
 *      at a rate of roughly O(1.618^N) (the Golden Ratio).
 *    - A looser, safer upper bound commonly used in interviews is O(2^N) because it
 *      is a binary decision tree with a maximum depth of N.
 *
 * 2. Work Per State -> O(N)
 *    - In Go, strings are immutable. Doing `tmp+char` inside the recursive call
 *      forces the program to allocate a brand new string and copy the existing
 *      characters over.
 *    - Because the decoded string `tmp` can grow to a maximum length of N, this
 *      copy operation takes up to O(N) time at every single recursive step.
 *    - (Note: The `for` loop itself only ever runs a maximum of 2 times, so the
 *      `num.WriteString` and the map lookup `numberAlbhabetMap[num.String()]`
 *      are resolved in O(1) constant time).
 *
 * Total Time = O(2^N) * O(N) = O(N * 2^N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case path (where the algorithm decodes exactly
 *    1 character at a time), the call stack will go exactly N levels deep.
 * 2. Auxiliary Space: String creation takes extra memory in each stack frame, but
 *    the active footprint at any given depth is bounded by O(N).
 *
 * Total Space = Stack O(N) + Strings O(N) = O(N)
 */
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

	var num strings.Builder
	for i := idx; i < idx+2 && i < len(s); i++ {
		num.WriteString(string(s[i]))
		char, ok := numberAlbhabetMap[num.String()]
		if ok {
			generateDecodedStrings(i+1, s, tmp+char, result)
		}
	}
}

// https://leetcode.com/problems/decode-ways/
/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N)
 * ----------------------------
 * Where N is the length of the string `s`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 *    - The state is defined solely by the current index `idx`.
 *    - Thanks to the `memo` map, the algorithm guarantees that it will never
 *      re-evaluate the number of decodings for the same index twice.
 *    - Since `idx` can only range from 0 to N, there are exactly N unique states
 *      that need to be computed.
 *
 * 2. Work Per State -> O(1)
 *    - Inside each recursive call (when the result is not already cached), the
 *      algorithm performs basic checks (`s[idx] == '0'`) and at most two recursive calls.
 *    - Slicing a strictly 2-character string (`s[idx : idx+2]`) and converting it
 *      to an integer with `strconv.Atoi` takes constant O(1) time, regardless of
 *      how large N gets.
 *    - Map insertions and lookups operate in O(1) average time.
 *
 * Total Time = O(N) unique states * O(1) work per state = O(N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case scenario (where the algorithm consistently
 *    evaluates the `idx+1` branch), the call stack goes exactly N levels deep before
 *    hitting the base case. This takes O(N) space.
 * 2. Auxiliary Space: The `memo` map stores exactly N key-value pairs (mapping an
 *    index to its decoded integer count). This consumes O(N) memory.
 *
 * Total Space = Stack O(N) + Memo Map O(N) = O(N)
 */
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
