package main

func shipWithinDays(weights []int, days int) int {
	sum, maxW := 0, 0
	for _, w := range weights {
		sum += w
		if w > maxW {
			maxW = w
		}
	}

	// avg is the ceiling of sum/days: the average load per day if the total
	// weight were split as evenly as possible across all days.
	// e.g. weights = [1..10], sum = 55, days = 5 -> avg = ceil(55/5) = 11.
	avg := (sum + days - 1) / days // ceil(sum / days)

	// low is the smallest capacity we even need to try.
	// Two facts limit how small the answer can be:
	//   1. The ship's capacity must fit the single heaviest package, so it
	//      can never be less than maxW (e.g. a capacity of 9 can't carry a
	//      package that weighs 10).
	//   2. Even in the best case (perfectly even split), capacity can't be
	//      less than avg (e.g. you can't ship 55 total weight in 5 days
	//      with a boat that only holds 10 per day - that's only 50 total).
	// So low starts at whichever of these two is bigger.
	// e.g. maxW = 10, avg = 11 -> low = max(10, 11) = 11.
	low := max(maxW, avg)

	// high is a capacity we're sure WILL work - it doesn't need to be tight,
	// just a safe starting point for the binary search.
	//
	// Example: weights = [1..10], sum = 55, days = 5, maxW = 10, avg = 11.
	// Here's how we build up to high = avg + maxW - 1 = 11 + 10 - 1 = 20:
	//
	//  1. We want each day to carry at least the average, 11, so that 5 days
	//     is enough to cover the whole 55.
	//  2. But packages are lumpy, not liquid - a day can stop early with
	//     unused room, if the next package just doesn't fit. How much room
	//     could be wasted like that? At most maxW - 1 = 9. (If 10 or more
	//     room were left, the next package - at most maxW = 10 - would still
	//     have fit, so the day wouldn't have stopped.)
	//  3. So, worst case, a day might carry its target of 11 but ALSO waste
	//     up to 9 of unused capacity on top of that. To guarantee it still
	//     carries at least 11 despite that waste, we pad the capacity by the
	//     worst-case waste: capacity = 11 (target) + 9 (worst-case waste)
	//     = 20.
	//  4. With capacity 20, every "full" day carries at least 11. 5 such
	//     days already cover 5 x 11 = 55, the entire shipment - so we never
	//     need more than 5 days.
	//
	// Simulating canShip with capacity 20 on this example actually only
	// needs 4 days, comfortably within the 5-day limit.
	// (An even simpler choice would just be high = sum, since a capacity
	// that big always ships everything in a single day - but that starts
	// the search from a much bigger range, e.g. 11..55 instead of 11..20,
	// so it takes more steps to narrow down to the answer.)
	high := avg + maxW - 1

	for low < high {
		mid := low + (high-low)/2
		if canShip(weights, days, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canShip(weights []int, days, capacity int) bool {
	load, used := 0, 1
	for _, w := range weights {
		if load+w > capacity {
			used++
			if used > days {
				return false
			}
			load = 0
		}
		load += w
	}
	return true
}

// https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/description/
func shipWithinDaysMain() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5

	shipWithinDays(weights, days)
}
