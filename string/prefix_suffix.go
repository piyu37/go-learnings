package main

import (
	"fmt"
	"sort"
)

func calculateScore(text, prefixString, suffixString string) string {
	result := make(map[int][]string)
	lenText := len(text)

	for lenText > 0 {
		for i := 0; i < len(text)+1-lenText; i++ {
			substring := text[i : i+lenText]

			// calc the pre_text_pattern_score
			preTextPatternScore := min(len(prefixString), len(substring))

			subva1 := substring[:preTextPatternScore]
			preva1 := prefixString[len(prefixString)-preTextPatternScore:]
			for preTextPatternScore > 0 && subva1 != preva1 {
				preTextPatternScore--
				subva1 = substring[:preTextPatternScore]
				preva1 = prefixString[len(prefixString)-preTextPatternScore:]
			}

			// calc the post_text_pattern_score
			postTextPatternScore := min(len(suffixString), len(substring))

			subva1 = substring[len(substring)-postTextPatternScore:]
			sufva1 := suffixString[:postTextPatternScore]
			for postTextPatternScore > 0 && subva1 != sufva1 {
				postTextPatternScore--
				subva1 = substring[len(substring)-postTextPatternScore:]
				sufva1 = suffixString[:postTextPatternScore]
			}

			// calculate the pattern_score
			patternScore := preTextPatternScore + postTextPatternScore

			if _, ok := result[patternScore]; !ok {
				// resets the dictionary key
				result[patternScore] = make([]string, 0)
			}

			result[patternScore] = append(result[patternScore], substring)
		}

		lenText-- // reduce lenText by 1
	}

	// store the highest key, so we can sort the right item to return
	maximumPatternScore := max(getKeys(result))

	maxScore := result[min(maximumPatternScore, len(text))]

	sort.Strings(maxScore)
	return maxScore[0]
}

func max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}

	return v1
}

func getKeys(m map[int][]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// https://leetcode.com/discuss/interview-question/787504/navis-interview-question
func prefixSuffix() {
	text := "nothing"
	prefixString := "bruno"
	suffixString := "ingenious"
	score := calculateScore(text, prefixString, suffixString)
	fmt.Println(score)
}
