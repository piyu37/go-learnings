package main

import "fmt"

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(L * H * N)
 * ----------------------------
 * Where L is `low`, H is `high`, and N is the number of `measuringCups`.
 * (Note: If L and H scale together to represent a general "Volume", this is
 * sometimes generalized as O(V^2 * N)).
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(L * H)
 *    - Because you are caching results using a memoization map with the key
 *      `fmt.Sprintf("%d:%d", countLow, countHigh)`, the algorithm guarantees
 *      that it never re-evaluates the same state twice.
 *    - `countLow` ranges from 0 up to approximately `low` (plus the max cup size).
 *    - `countHigh` ranges from 0 up to approximately `high` (plus the max cup size).
 *    - Therefore, the theoretical maximum number of unique cached states is
 *      proportional to L * H.
 *
 * 2. Work Per State -> O(N)
 *    - When a state is NOT in the cache, the algorithm executes the `for` loop,
 *      iterating backward from `index` to 0. In the worst case, this iterates
 *      N times (where N is the total number of cups).
 *    - Creating the `memoizeKey` via `fmt.Sprintf` takes O(1) time relative
 *      to the input size.
 *
 * Total Time = O(L * H) * O(N) = O(L * H * N)
 *
 * SPACE COMPLEXITY: O(L * H)
 * ----------------------------
 * 1. Recursion Stack: In the worst-case scenario (e.g., using a cup of size 1),
 *    the call stack will go L levels deep to reach the target `low`. This takes
 *    O(L) space.
 * 2. Auxiliary Space: The `memoization` map stores a boolean for every unique
 *    (countLow, countHigh) pair evaluated. This takes up to O(L * H) space.
 *
 * Total Space = O(L) + O(L * H) = O(L * H)
 */
func AmbiguousMeasurements(measuringCups [][]int, low int, high int) bool {
	// A struct is extremely fast for map keys in Go (no string allocation overhead)
	type state struct {
		low  int
		high int
	}
	memo := make(map[state]bool)

	// Define our recursive function
	var canMeasure func(l, h int) bool
	canMeasure = func(l, h int) bool {
		// If our high boundary drops below 0, it means the current combination
		// of cups has exceeded the maximum allowed volume. This is a dead end.
		if h < 0 {
			return false
		}

		// If we've subtracted enough to satisfy the minimum requirement (l <= 0)
		// and we haven't exceeded the maximum requirement (h >= 0), we succeeded!
		if l <= 0 && h >= 0 {
			return true
		}

		// Check memoization table
		key := state{l, h}
		if val, exists := memo[key]; exists {
			return val
		}

		// Try every cup
		for _, cup := range measuringCups {
			cupLow, cupHigh := cup[0], cup[1]

			// Clamp newLow to 0. Once we meet the minimum volume,
			// we only care about not exceeding the maximum volume.
			newLow := l - cupLow
			if newLow < 0 {
				newLow = 0
			}

			// Recursively check if the new state can lead to a valid measurement
			if canMeasure(newLow, h-cupHigh) {
				memo[key] = true
				return true
			}
		}

		// If no cups lead to a valid measurement, mark as false
		memo[key] = false
		return false
	}

	return canMeasure(low, high)
}

// Same complexity as above
func AmbiguousMeasurementsBFS(measuringCups [][]int, low int, high int) bool {
	type state struct {
		l int
		h int
	}

	// Queue for BFS
	queue := []state{{l: low, h: high}}

	// Visited array to prevent duplicate work
	visited := make([][]bool, low+1)
	for i := range visited {
		visited[i] = make([]bool, high+1)
	}
	visited[low][high] = true

	for len(queue) > 0 {
		// Pop the front of the queue
		curr := queue[0]
		queue = queue[1:]

		// If we've satisfied the requirements, we are done
		if curr.l <= 0 && curr.h >= 0 {
			return true
		}

		// Try every cup for the current state
		for _, cup := range measuringCups {
			cupLow, cupHigh := cup[0], cup[1]

			newLow := max(curr.l-cupLow, 0)
			newHigh := curr.h - cupHigh

			// If the new high is below 0, this path is dead.
			// If we haven't visited this state yet, add it to the queue.
			if newHigh >= 0 && !visited[newLow][newHigh] {
				visited[newLow][newHigh] = true
				queue = append(queue, state{l: newLow, h: newHigh})
			}
		}
	}

	return false
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/33_ambiguous_measurements
func ambiguousMeasurementsMain() {
	mc := [][]int{
		{1, 3},
		{2, 4},
		{5, 6},
	}
	low := 100
	high := 101

	fmt.Println(AmbiguousMeasurements(mc, low, high))
	fmt.Println(AmbiguousMeasurementsBFS(mc, low, high))
}
