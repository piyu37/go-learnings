package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func CorrectPath(s string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano())
	for {
		route := make([]rune, 0)
		tracepos := make(map[int]map[int]bool)
		x, y, answer := 1, 5, 1

		for _, i := range s {
			if i == '?' {
				v := r.Intn(4)
				str := "lrud"[v]
				i = rune(str)
			}

			if i == 'u' {
				y++
			} else if i == 'd' {
				y--
			} else if i == 'r' {
				x++
			} else if i == 'l' {
				x--
			}

			if _, ok := tracepos[x]; !ok {
				tracepos[x] = make(map[int]bool)
			}

			if tracepos[x][y] {
				answer = 0
				break
			} else {
				tracepos[x][y] = true
			}

			route = append(route, i)

			if x == 6 || x == 0 || y == 0 || y == 6 {
				answer = 0
				break
			}
		}

		if x == 5 && y == 1 && answer == 1 {
			result := string(route)
			result = removeCharacters(result, "gp27f91")
			if result == "" {
				return "EMPTY"
			}
			return result
		}
	}
}
func removeCharacters(str, charsToRemove string) string {
	for _, char := range charsToRemove {
		str = strings.ReplaceAll(str, string(char), "")
	}
	return str
}

// https://github.com/Liangmp/coderbyte-solution/blob/master/coderbyte-CorrectPath.py
func generatePath() {
	fmt.Println(CorrectPath("drdr??rrddd?"))
}
