package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(K * N^K) where K = Target / Min(Candidates)
 * Large Array, Small Target N^K is better => This approach is beter than below approach
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N^K)
 *    - Let N be the number of candidates. At each node, we can make up to N choices.
 *    - Because candidates can be reused, the recursion tree stops only when the target
 *      reaches 0 or less. The maximum possible depth is K = (Target / Minimum_Candidate).
 *    - Max states in the worst-case tree = N * N * N ... (K times) = N^K.
 *
 * 2. Work Per State -> O(K)
 *    - The standard recursive branching (appending and slicing `currComb`) takes O(1) time.
 *    - Additionally, when a valid combination is found (`target == 0`), copying
 *      the `currComb` slice into `temp` takes up to O(K) time, since the combination
 *      can be at most K elements long.
 *
 * Total Time = O(N^K) * O(K) = O(K * N^K)
 *
 * SPACE COMPLEXITY: O(K)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most K levels deep in the worst case.
 * 2. Auxiliary Space: The `currComb` slice holds a maximum of K integers.
 *
 * Total Space = O(K) + O(K) = O(K)
 */
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	findCombinations(0, candidates, target, []int{}, &result)

	return result
}

func findCombinations(idx int, candidates []int, target int, currComb []int, result *[][]int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(currComb))
		copy(temp, currComb)
		*result = append(*result, temp)
		return
	}

	for i := idx; i < len(candidates); i++ {
		currComb = append(currComb, candidates[i])
		findCombinations(i, candidates, target-candidates[i], currComb, result)
		currComb = currComb[:len(currComb)-1]
	}
}

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(K * 2^(K+N)) where K = Target / Min(Candidates)
 * Small Array, Huge Target 2^(K+N) is better)
 * ----------------------------
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(2^(K+N))
 *    - Unlike the loop-based approach, this code uses a "Pick or Skip" strategy.
 *      At every recursive step, it makes exactly 2 choices: pick the candidate
 *      (decreases target) or skip it (increases index).
 *    - The maximum number of times you can "pick" is K = Target / Min(Candidates).
 *    - The maximum number of times you can "skip" is N (length of candidates).
 *    - The maximum depth of the recursion tree is K + N. Since there are 2 choices
 *      at each level, the worst-case number of states is 2^(K+N).
 *
 * 2. Work Per State -> O(K)
 *    - Standard recursive branching takes O(1) time.
 *    - When a valid combination is found (`target == 0`), copying the `temp`
 *      slice into the new `arr` takes up to O(K) time, as the combination
 *      can be at most K elements long.
 *
 * Total Time = O(2^(K+N)) * O(K) = O(K * 2^(K+N))
 *
 * SPACE COMPLEXITY: O(K + N)
 * ----------------------------
 * 1. Recursion Stack: The call stack goes at most K + N levels deep.
 * 2. Auxiliary Space: The `temp` slice holds a maximum of K integers.
 *
 * Total Space = O(K+N) + O(K) = O(K + N)
 */
func combinationSumDiffApproach(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	calcSum(candidates, target, 0, []int{}, &result)

	return result
}

func calcSum(candidates []int, target int, currIdx int, temp []int, result *[][]int) {
	if target < 0 || currIdx >= len(candidates) {
		return
	}

	if target == 0 {
		arr := make([]int, len(temp))
		copy(arr, temp)

		*result = append(*result, arr)
		return
	}

	calcSum(candidates, target-candidates[currIdx], currIdx, append(temp, candidates[currIdx]), result)
	calcSum(candidates, target, currIdx+1, temp, result)
}

// https://leetcode.com/problems/combination-sum/description/?envType=study-plan-v2&envId=top-interview-150
func combinationSumMain() {
	candidates := []int{2, 3, 5}
	target := 8

	fmt.Println(combinationSum(candidates, target))

	candidates = []int{186, 419, 83, 408}
	target = 6249

	fmt.Println(combinationSum(candidates, target))
	fmt.Println(combinationSumDiffApproach(candidates, target))
}
