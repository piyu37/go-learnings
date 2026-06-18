package main

import "fmt"

var phoneNemonicsMap = map[string][]string{
	"1": {"1"},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
	"0": {"0"},
}

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
 * - Instead of using a recursive depth-first approach, this algorithm builds
 * combinations iteratively in a breadth-first manner.
 * - For each digit in the phone number, it takes all previously built strings
 * and branches them out by adding up to 4 new characters.
 * - Across all N iterations, the total number of intermediate strings generated
 * scales exponentially, maxing out at roughly 4^N combinations at the final step.
 *
 * 2. Work Per State -> O(N)
 * - Inside the innermost loop, the algorithm performs string concatenation
 * (`tempResult[v] + ch`).
 * - Because strings are immutable in Go, concatenating creates an entirely new
 * string in memory. The length of these strings grows with each iteration,
 * eventually reaching length N. Copying these characters takes O(N) time.
 * - Reassigning and copying the arrays (`append([]string{}, result...)`) also
 * adds linear work proportional to the number of combinations.
 *
 * Total Time = O(4^N) intermediate states generated * O(N) work per state = O(4^N * N)
 *
 * SPACE COMPLEXITY: O(4^N * N)
 * ----------------------------
 * 1. Recursion Stack: Because this is a purely iterative solution, there are no
 * recursive function calls. The call stack depth remains completely flat,
 * using O(1) space.
 * 2. Auxiliary Space: Unlike the recursive approach (which only needs to hold a
 * single O(N) string in memory at any given time along the stack), this
 * iterative approach must hold *all* intermediate combinations in memory
 * simultaneously. During the final iteration, the `tempResult` array holds
 * O(4^(N-1)) strings while the `result` array builds O(4^N) strings, all of
 * which approach length N.
 *
 * Total Space = Stack O(1) + Auxiliary Arrays O(4^N * N) = O(4^N * N)
 */
func PhoneNumberMnemonics(phoneNumber string) []string {
	result := []string{}

	for _, ru := range phoneNumber {
		char := string(ru)

		if len(result) == 0 {
			allChars := append([]string{}, phoneNemonicsMap[char]...)

			result = append(result, allChars...)

			continue
		}

		tempResult := append([]string{}, result...)
		result = []string{}

		for v := range tempResult {
			allChars := append([]string{}, phoneNemonicsMap[char]...)

			for i := range allChars {
				ch := allChars[i]

				result = append(result, tempResult[v]+ch)
			}
		}
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/41_phone_number_mnemonics
func mnemonics() {
	phoneNum := "1905"

	fmt.Println(PhoneNumberMnemonics(phoneNum))
}
