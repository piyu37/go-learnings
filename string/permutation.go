package main

import (
	"fmt"
	"strings"
)

func permute(words []string, start int, perMap map[string]bool) map[string]bool {
	if start == len(words)-1 {
		value := strings.Join(words, "")
		perMap[value] = true
		return perMap
	}

	for i := start; i < len(words); i++ {
		words[start], words[i] = words[i], words[start] // Swap elements
		perMap = permute(words, start+1, perMap)
		words[start], words[i] = words[i], words[start] // Restore the original order
	}

	return perMap
}

func findSubstring(s string, words []string) []int {
	perMap := permute(words, 0, map[string]bool{})

	permWordLen := len(words) * len(words[0])

	result := make([]int, 0)

	if permWordLen > len(s) {
		return result
	}

	i := 0
	for i <= len(s)-permWordLen {
		word := s[i : i+permWordLen]

		if _, ok := perMap[word]; ok {
			result = append(result, i)
			if s[i] == s[i+1] {
				i++
			} else {
				i += len(words[0])
			}
		} else {
			i++
		}
	}

	return result
}

func findSubstring2(s string, words []string) []int {
	result := make([]int, 0)

	wordsMap := make(map[string]int)

	for i := range words {
		wordsMap[words[i]]++
	}

	permWordLen := len(words) * len(words[0])

	if permWordLen > len(s) {
		return result
	}

	for right := permWordLen - 1; right < len(s); right++ {
		left := right - permWordLen + 1
		flag := false

		wordsMapCopy := make(map[string]int)

		for k, v := range wordsMap {
			wordsMapCopy[k] = v
		}

		for left <= right {
			word := s[left : left+len(words[0])]

			if count, ok := wordsMapCopy[word]; ok {
				if count == 0 {
					flag = true
					break
				}

				wordsMapCopy[word]--
				left += len(words[0])
			} else {
				flag = true
				break
			}
		}

		if !flag {
			totalCount := 0
			for _, v := range wordsMapCopy {
				totalCount += v
			}

			if totalCount == 0 {
				result = append(result, right-permWordLen+1)
			}
		}
	}

	return result
}

func findSubstringBest(s string, words []string) []int {
	result := make([]int, 0)

	wordsMap := make(map[string]int)

	for i := range words {
		wordsMap[words[i]]++
	}

	wordLen := len(words[0])

	permWordLen := len(words) * wordLen

	if permWordLen > len(s) {
		return result
	}

	for i := 0; i < wordLen; i++ {
		right := i
		left := i
		wordsUsed := 0
		excessWord := false
		tmpWordsMap := make(map[string]int)
		for right <= len(s)-wordLen {

			word := s[right : right+wordLen]

			if _, ok := wordsMap[word]; !ok {
				tmpWordsMap = make(map[string]int)
				wordsUsed = 0
				excessWord = false
				right += wordLen
				left = right
				continue
			} else {
				for right-left == permWordLen || excessWord {
					leftWord := s[left : left+wordLen]
					left += wordLen
					tmpWordsMap[leftWord]--

					if wordsMap[leftWord] == tmpWordsMap[leftWord] {
						excessWord = false
					} else {
						wordsUsed--
					}
				}

				tmpWordsMap[word]++
				if tmpWordsMap[word] <= wordsMap[word] {
					wordsUsed++
				} else {
					excessWord = true
				}

				if wordsUsed == len(words) && !excessWord {
					result = append(result, left)
				}
			}

			right += wordLen
		}
	}

	return result
}

// https://leetcode.com/problems/substring-with-concatenation-of-all-words/
func permutation() {
	words := []string{"word", "good", "best", "word"}
	fmt.Println(findSubstring("aaaaaaaaaaaaaa", words))
	fmt.Println(findSubstring2("barfoothefoobarman", words))
	fmt.Println(findSubstringBest("wordgoodgoodgoodbestword", words))
}
