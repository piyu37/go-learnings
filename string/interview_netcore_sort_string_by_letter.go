package main

import (
	"fmt"
	"strings"
)

// Sort the input by letter/char without changing the case(upper/lower) in string
// i/p = Netcore Bangalore
// o/p = aaBceeeglnNoorrt
func sortStringByLetter() {
	input := "Netcore Bangalore"
	output := ""

	mapChar := make(map[string]int)

	for i := range input {
		c := string(input[i])

		if _, ok := mapChar[c]; ok {
			mapChar[c] += 1
			continue
		}

		mapChar[c] = 1
	}

	count := 0
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		c := string(v)
		if noOfTimes, exist := mapChar[c]; exist {
			loopCount := count + noOfTimes
			for i := count; i < loopCount; i++ {
				output += c
				count++
			}
		}

		upperC := strings.ToUpper(c)
		if noOfTimes, exist := mapChar[upperC]; exist {
			loopCount := count + noOfTimes
			for i := count; i < loopCount; i++ {
				output += upperC
				count++
			}
		}
	}

	fmt.Println(output)
}
