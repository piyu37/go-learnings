package main

import (
	"fmt"
	"sort"
)

func getMinConnectionCost(warehouseCapacity []int32, additionalHubs [][]int32) []int64 {
	n := len(warehouseCapacity)
	cap := warehouseCapacity
	result := make([]int64, len(additionalHubs))

	for q, hubs := range additionalHubs {
		// Convert hubs to 0-based
		h1 := int(hubs[0] - 1)
		h2 := int(hubs[1] - 1)
		hubPositions := []int{h1, h2, n - 1}
		sort.Ints(hubPositions)

		var total int64 = 0
		for i := 0; i < n; i++ {
			// Find the nearest hub at or after position i
			idx := sort.Search(len(hubPositions), func(j int) bool {
				return hubPositions[j] >= i
			})
			hub := hubPositions[idx]
			cost := int64(cap[hub]) - int64(cap[i])
			total += cost
		}
		result[q] = total
	}

	return result
}

// “You are given an array of non-decreasing integers. For each query, you are given two hub positions.
// Together with the last element, these form the hub positions. For every index in the array,
// find the nearest hub on or after that index. The cost is A[hub] - A[i]. Return the total cost for each query.”
// example:
// warehouseCapacity = [3, 6, 10, 15, 20]
// query hubs = (2, 4) => 6, 15 from warehouse capacity; does not follow index; start with 1
// Convert to 0-based indices: hubs = {1, 3, 4}. => it follows index value. 6 is at idx 1, 15 at idx 3 &
// always include last index from warehouseCapacity.

// Now check each warehouse:

// Warehouse 0 → nearest hub at/after 0 = index 1
// cost = cap[1] - cap[0] = 6 - 3 = 3

// Warehouse 1 → nearest hub at/after 1 = index 1
// cost = 6 - 6 = 0

// Warehouse 2 → nearest hub at/after 2 = index 3
// cost = 15 - 10 = 5

// Warehouse 3 → nearest hub at/after 3 = index 3
// cost = 15 - 15 = 0

// Warehouse 4 → nearest hub at/after 4 = index 4
// cost = 20 - 20 = 0

// 👉 Total = 3 + 0 + 5 + 0 + 0 = 8

// That’s why the program prints [8].
func getMinConnectionCostMain() {
	warehouseCapacity := []int32{3, 6, 10, 15, 20}
	additionalHubs := [][]int32{{2, 4}}
	fmt.Println(getMinConnectionCost(warehouseCapacity, additionalHubs)) // Output: [8]

	warehouseCapacity = []int32{0, 2, 5, 9, 12, 18}
	additionalHubs = [][]int32{{2, 5}, {1, 3}}
	fmt.Println(getMinConnectionCost(warehouseCapacity, additionalHubs))
}
