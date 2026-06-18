package main

import (
	"fmt"
	"strings"
)

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(4^N * N)
 * ----------------------------
 * Where N is the length of the string `phoneNumber`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(4^N)
 * - The algorithm builds a combinations tree. The depth of the recursion tree is
 * exactly N (one level for each digit in the phone number).
 * - At each level, the algorithm branches out based on the number of letters
 * associated with the current digit. On a standard telephone keypad, digits
 * map to at most 4 letters (e.g., '7' is PQRS, '9' is WXYZ).
 * - In the absolute worst-case scenario (a phone number consisting entirely
 * of 7s and 9s), the tree will have a branching factor of 4. A tree with
 * depth N and a branching factor of 4 has 4^N leaf nodes.
 *
 * 2. Work Per State -> O(N)
 * - At intermediate states (non-leaf nodes), the work is O(1) as it just
 * assigns a character to `currentMnemonic[idx]`.
 * - However, at the base case (the leaf nodes where `idx == len(phoneNumber)`),
 * the algorithm calls `strings.Join(currentMnemonic, "")`. Since the length
 * of the mnemonic is always N, iterating through the slice and creating a
 * new string takes O(N) time.
 *
 * Total Time = O(4^N) maximum states * O(N) work to build the string = O(4^N * N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The depth of the recursion tree is strictly N. In the
 * worst case, the call stack goes N levels deep before reaching the base case
 * and returning. This requires O(N) space.
 * 2. Auxiliary Space: Standard space complexity analysis excludes the memory
 * required for the final `result` array (which would take O(4^N * N) space).
 * Excluding the output, the algorithm allocates a single `currentMnemonic`
 * slice of length N, which is reused across all recursive calls. This uses
 * O(N) memory.
 *
 * Total Space = Stack O(N) + Auxiliary O(N) (excluding output) = O(N)
 */
func PhoneNumberMnemonicsRecur(phoneNumber string) []string {
	currentMnemonic := make([]string, len(phoneNumber))

	for i := range currentMnemonic {
		currentMnemonic[i] = "0"
	}

	result := make([]string, 0)

	phoneNumMnemonicHelper(0, phoneNumber, currentMnemonic, &result)

	return result
}

func phoneNumMnemonicHelper(idx int, phoneNumber string, currentMnemonic []string, result *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := strings.Join(currentMnemonic, "")
		*result = append(*result, mnemonic)

		return
	}

	ch := string(phoneNumber[idx])

	allChars := phoneNemonicsMap[ch]

	for i := range allChars {
		c := allChars[i]

		currentMnemonic[idx] = c

		phoneNumMnemonicHelper(idx+1, phoneNumber, currentMnemonic, result)
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/41_phone_number_mnemonics
func mnemonicsRecur() {
	phoneNum := "1905"

	fmt.Println(PhoneNumberMnemonicsRecur(phoneNum))
}
