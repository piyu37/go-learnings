package main

import "fmt"

func matchWordsInBoggle(board [][]string, words []string) []string {
	letterMap := make(map[string][][]int)
	result := []string{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			letterMap[board[i][j]] = append(letterMap[board[i][j]], []int{i, j})
		}
	}

	visited := make([][]bool, len(board))

	for i := range visited {
		col := make([]bool, len(board[0]))
		visited[i] = col
	}

	for _, word := range words {
		if len(word) == 1 {
			if _, exist := letterMap[word]; exist {
				result = append(result, word)
				continue
			}
		}

		if list, exist := letterMap[string(word[0])]; exist {
			for _, pos := range list {
				i, j := pos[0], pos[1]
				nxtLetterList := letterMap[string(word[1])]
				visited[i][j] = true
				wordMatched := matchWord(board, word, string(word[0]), i, j, nxtLetterList, letterMap,
					visited, 1)
				visited[i][j] = false
				if wordMatched {
					result = append(result, word)
					break
				}
			}
		} else {
			continue
		}
	}

	return result
}

func matchWord(board [][]string, word, matchingWord string, prevI, prevJ int,
	list [][]int, letterMap map[string][][]int, visited [][]bool, wordIdx int) bool {
	for _, pos := range list {
		i, j := pos[0], pos[1]

		if visited[i][j] {
			continue
		}

		isNeighbor := max(mod(i-prevI), mod(j-prevJ)) <= 1
		if !isNeighbor {
			continue
		}

		visited[i][j] = true

		matchingWord += board[i][j]

		if len(word) == len(matchingWord) {
			visited[i][j] = false
			return word == matchingWord
		}

		nxtLetterList, exist := letterMap[string(word[wordIdx+1])]
		if !exist {
			return false
		}

		isWordMatched := matchWord(board, word, matchingWord, i, j, nxtLetterList,
			letterMap, visited, wordIdx+1)
		visited[i][j] = false
		if isWordMatched {
			return true
		}
	}

	return false
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func mod(v1 int) int {
	if v1 >= 0 {
		return v1
	}

	return -v1
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/21_boggle_board
func boggleBoard() {
	board := [][]string{
		{"t", "h", "i", "s", "i", "s", "a"},
		{"s", "i", "m", "p", "l", "e", "x"},
		{"b", "x", "x", "x", "x", "e", "b"},
		{"x", "o", "g", "g", "l", "x", "o"},
		{"x", "x", "x", "D", "T", "r", "a"},
		{"R", "E", "P", "E", "A", "d", "x"},
		{"x", "x", "x", "x", "x", "x", "x"},
		{"N", "O", "T", "R", "E", "-", "P"},
		{"x", "x", "D", "E", "T", "A", "E"},
	}

	words := []string{
		"this", "is", "not", "a", "simple", "boggle",
		"board", "test", "REPEATED", "NOTRE-PEATED",
	}

	fmt.Println(matchWordsInBoggle(board, words))
}
