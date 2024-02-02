package main

import (
	"fmt"
	"math"
	"sort"
)

type movie struct {
	releaseTime int32
	duration    int32
}

func minimumTimeSpent(comedyReleaseTime, comedyDuration, dramaReleaseTime, dramaDuration []int32) int32 {
	comedy := combine(comedyReleaseTime, comedyDuration)
	drama := combine(dramaReleaseTime, dramaDuration)

	sort.Slice(comedy, func(i, j int) bool {
		return comedy[i].releaseTime < comedy[j].releaseTime
	})

	sort.Slice(drama, func(i, j int) bool {
		return drama[i].releaseTime < drama[j].releaseTime
	})

	totalTime := int32(math.MaxInt32)

	for i := range comedy {
		comedyMovie := comedy[i]
		time := comedyMovie.releaseTime + comedyMovie.duration

		for j := range drama {
			dramaMovie := drama[j]
			time += dramaMovie.duration
			if time < totalTime {
				totalTime = time
			}
		}
	}

	if totalTime == int32(math.MaxInt32) {
		return 0
	}

	return totalTime
}

func combine(releaseTime, duration []int32) []movie {
	var combined []movie
	for i := 0; i < len(releaseTime); i++ {
		combined = append(combined, movie{releaseTime[i], duration[i]})
	}
	return combined
}

// Amazon Prime Video has movies in the category 'comedy' or 'drama'.
// Determine the earliest time you can finish at least one movie from each category.
// The release schedule and durations of the movies are provided.
// You can start watching a movie at the time it is released or later.
// If you begin a movie at time t, it ends at time t + duration.
// If a movie ends at time t + duration, the second movie can start at that time,
// t + duration, or later. The movies can be watched in any order.

// Example:
// comedyReleaseTime = [1, 4]
// comedyDuration = [3, 2]
// dramaReleaseTime = [5, 2]
// dramaDuration = [2, 2]
// Duration and release times are aligned by index.

// Example Answer:
// Two of the best ways to finish watching one movie from each category at the earliest time are as follows:
// 1. Start watching comedy movie1 at time t = 1 and until t = 1 + 3 = 4. Then, watch the drama movie2 from time t = 4 to t = 4 + 2 = 6.
// 2. Start watching drama movie2 at time t = 2 and until t = 2 + 2 = 4.Then, watch the comedy movie2 from time t = 4 to t = 4 + 2 = 6.
// The earliest finish time and also the answer is 6.
func watchMovies() {
	comedyReleaseTime := []int32{1, 1}
	comedyDuration := []int32{1, 1}
	dramaReleaseTime := []int32{1, 1, 1}
	dramaDuration := []int32{1, 1, 1}

	result := minimumTimeSpent(comedyReleaseTime, comedyDuration, dramaReleaseTime, dramaDuration)
	fmt.Println("The earliest finish time and also the answer is:", result)
}
