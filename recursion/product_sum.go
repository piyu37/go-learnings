package main

import "fmt"

type SpecialArray []any

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N)
 * ----------------------------
 * Where N is the total number of elements in the array, including all integers
 * and all nested subarrays.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 * - The algorithm processes a deeply nested array structure, which can be
 * visualized as a tree where arrays are branches and integers are leaf nodes.
 * - It traverses every single item inside the top-level array and all nested
 * subarrays exactly once.
 * - Because no element or subarray is evaluated more than once, the total
 * number of states (or nodes visited) scales linearly with the total count
 * of elements, N.
 *
 * 2. Work Per State -> O(1)
 * - Inside the `for` loop, the algorithm performs a basic type switch
 * (`switch v := data.(type)`), which evaluates in constant O(1) time.
 * - If the element is an integer, it performs a simple addition (`total += v`).
 * - If the element is a nested array, it triggers a recursive call. The
 * arithmetic operations (addition and final multiplication `total * mul`)
 * all execute in O(1) constant time regardless of the size of N.
 *
 * Total Time = O(N) elements visited * O(1) work per element = O(N)
 *
 * SPACE COMPLEXITY: O(D)
 * ----------------------------
 * Where D is the maximum depth (or level of nesting) of the `SpecialArray`.
 *
 * 1. Recursion Stack: Every time the algorithm encounters a nested `SpecialArray`,
 * it makes a recursive call, adding a new frame to the call stack. In the
 * worst-case scenario (e.g., an array like `[[[[[5]]]]]`), the call stack
 * goes exactly D levels deep before returning. This consumes O(D) space.
 * 2. Auxiliary Space: The algorithm only allocates primitive integer variables
 * (`total` and the type-asserted `v`) within each stack frame. It never
 * allocates new arrays, slices, or maps during the traversal. This takes
 * O(1) extra space.
 *
 * Total Space = Stack O(D) + Auxiliary O(1) = O(D)
 */
func ProductSum(array []any) int {
	if len(array) == 0 {
		return -1
	}

	return sumOfArray(array, 1)
}

func sumOfArray(array []any, mul int) int {
	total := 0

	for _, data := range array {
		switch v := data.(type) {
		case int:
			total += v

		case SpecialArray:
			total += sumOfArray(v, mul+1)

		default:
			return total * mul
		}
	}

	return total * mul
}

// https://github.com/lee-hen/Algoexpert/tree/master/easy/08_product_sum
func productSumMain() {
	arr := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{-13, 8}, 4}}

	fmt.Println(ProductSum(arr))
}
