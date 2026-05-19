package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(4^N * N)
 * ----------------------------
 * Where N is `numberOfTags`.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(4^N)
 *    - To build a valid combination, the algorithm must place exactly N opening tags
 *      and N closing tags, meaning the recursion tree goes exactly 2N levels deep.
 *    - At each level, there are at most 2 choices: add an opening tag OR a closing tag.
 *    - An unpruned tree with depth 2N and branching factor 2 has 2^(2N) leaf nodes,
 *      which mathematically simplifies to 4^N.
 *    - Note for interviewer: The algorithm prunes invalid paths (e.g., never adding
 *      a closing tag when it exceeds opening tags). This means the actual number of
 *      valid states is the N-th Catalan number, which is strictly less than 4^N.
 *      However, O(4^N) is the standard, accepted upper bound.
 *
 * 2. Work Per State -> O(N)
 *    - Inside each recursive call, the algorithm concatenates a string
 *      (`partialString + "<div>"` or `partialString + "</div>"`).
 *    - Because strings are immutable in Go, concatenating creates a completely
 *      new string. The maximum length of this string is proportional to 2N,
 *      which takes O(N) time to allocate and copy.
 *    - The `append(result, ...)` operation also takes linear time based on the
 *      number of elements being bubbled up the call stack.
 *
 * Total Time = O(4^N) states * O(N) work per state = O(4^N * N)
 *
 * SPACE COMPLEXITY: O(N)
 * ----------------------------
 * 1. Recursion Stack: The algorithm goes exactly 2N levels deep (N opening tags
 *    and N closing tags) before hitting the base case. This requires O(N) space
 *    for the call stack frames.
 * 2. Auxiliary Space: Standard space complexity analysis excludes the memory
 *    required to hold the final output array. Excluding the result array, the
 *    only extra memory used is for the immutable strings created at each step,
 *    tracked along the O(N) deep stack.
 *
 * Total Space = Stack O(N) + Auxiliary O(1) (excluding output) = O(N)
 */
func GenerateDivTags(numberOfTags int) []string {
	result := createDivtags("", numberOfTags, numberOfTags)
	return result
}

func createDivtags(partialString string, openingTag, closingTag int) []string {
	result := make([]string, 0)

	if closingTag == 0 {
		result = append(result, partialString)
		return result
	}

	if openingTag > 0 {
		result = append(result, createDivtags(partialString+"<div>", openingTag-1, closingTag)...)
	}

	if closingTag > openingTag {
		result = append(result, createDivtags(partialString+"</div>", openingTag, closingTag-1)...)
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/32_generate_div_tags
func generateDivTagsMain() {
	fmt.Println(GenerateDivTags(3))
}
