package main

import "fmt"

type requestCounter struct {
	count     int
	timestamp int
}

func droppedRequests(requestTime_count int, requestTime []int) int {
	rc := [60]requestCounter{}

	droppedReqCount := 0

	for i := range requestTime_count {
		droppedFlag := false
		ts := requestTime[i] % 60

		if rc[ts].timestamp == requestTime[i] {
			rc[ts].count++

			if rc[ts].count > 3 {
				droppedFlag = true
			}
		} else {
			rc[ts] = requestCounter{1, requestTime[i]}
		}

		c10, c60 := windowCounts(requestTime[i], &rc)
		droppedFlag = droppedFlag || c10 > 20 || c60 > 60

		if droppedFlag {
			droppedReqCount++
		}

	}

	return droppedReqCount
}

func windowCounts(currTime int, rc *[60]requestCounter) (last10, last60 int) {
	for _, slot := range rc {
		age := currTime - slot.timestamp
		if age < 60 {
			last60 += slot.count
			if age < 10 {
				last10 += slot.count
			}
		}
	}
	return
}

func droppedRequestsOptimized(n int, t []int) int {
	dropped := 0
	for i := range n {
		if (i >= 3 && t[i-3] == t[i]) ||
			(i >= 20 && t[i]-t[i-20] < 10) ||
			(i >= 60 && t[i]-t[i-60] < 60) {
			dropped++
		}
	}

	return dropped
}

// Rate Limiter Proxy

// User requests for a service are routed through a rate-limiting proxy to ensure that the network is not overloaded by
// non-critical requests.

// The proxy has the following limits:

// The number of requests in any given second cannot exceed 3.
// The number of requests in any given 10-second period cannot exceed 20. A ten-second period includes all requests arriving
// from any time max(1, T-9) to T (inclusive of both) for any valid time T.
// The number of requests in any given minute cannot exceed 60. Similar to above, 1 minute is from max(1, T-59) to T.

// Any request that exceeds any of the above limits will be dropped by the proxy. Given the times at which different requests
// arrive, sorted ascending, find how many requests will be dropped.

// Note: Even if a request is dropped, it is still considered for future calculations. However, if a request is dropped due to
// multiple violations, it is still counted only once.

// Example

// n = 27
// requestTime = [1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 11, 11, 11, 11]
// Request 1: Not dropped.
// Request 1: Not dropped.
// Request 1: Not dropped.
// Request 1: Dropped. At most 3 requests are allowed in one second.
// No request will be dropped until time 6, since all arrive at an allowed rate of 3 per second and the 10-second limit is not violated.
// Request 7: Not dropped. The total number of requests has reached 20 now.
// Request 7: Dropped. At most 20 requests are allowed in ten seconds.
// Request 7: Dropped. At most 20 requests are allowed in ten seconds.
// Request 7: Dropped. At most 20 requests are allowed in ten seconds. The 1-second limit is also violated here.
// Request 11: Not dropped. The 10-second window is now 2 to 11, so the total number of requests in this window is 20.
// Request 11: Dropped. At most 20 requests are allowed in ten seconds.
// Request 11: Dropped. At most 20 requests are allowed in ten seconds.
// Request 11: Dropped. At most 20 requests are allowed in ten seconds. Also, at most 3 requests are allowed per second.

// Hence, a total of 7 requests are dropped.
func droppedRequestsMain() {
	n := 27
	requestTime := []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7, 11, 11, 11, 11}

	fmt.Println(droppedRequests(n, requestTime))
	fmt.Println(droppedRequestsOptimized(n, requestTime))

	n = 21
	requestTime = []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7}

	fmt.Println(droppedRequests(n, requestTime))
	fmt.Println(droppedRequestsOptimized(n, requestTime))
}
