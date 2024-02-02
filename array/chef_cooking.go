package main

import (
	"fmt"
	"sort"
	"strings"
)

func evaluateCookingDay(ings []string, toi, toc int) string {
	ingMap := make(map[int]int)
	catArr := make([]int, 5)

	output := ""

	for i := range ings {
		if strings.Contains(ings[i], "fat") {
			ingMap[i] = i + 1
			catArr[0]++
		} else if strings.Contains(ings[i], "carb") {
			ingMap[i] = i + 2
			catArr[1]++
		} else if strings.Contains(ings[i], "protein") {
			ingMap[i] = i + 3
			catArr[2]++
		} else if strings.Contains(ings[i], "fiber") {
			ingMap[i] = i + 3
			catArr[3]++
		} else if strings.Contains(ings[i], "seasoning") {
			ingMap[i] = i + 4
			catArr[4]++
		}

		for k, v := range ingMap {
			if v < i {
				delete(ingMap, k)
				if strings.Contains(ings[k], "fat") {
					catArr[0]--
				} else if strings.Contains(ings[k], "carb") {
					catArr[1]--
				} else if strings.Contains(ings[k], "protein") {
					catArr[2]--
				} else if strings.Contains(ings[k], "fiber") {
					catArr[3]--
				} else if strings.Contains(ings[k], "seasoning") {
					catArr[4]--
				}
			}
		}

		totalCat := 0
		for cIdx := range catArr {
			if catArr[cIdx] > 0 {
				totalCat++
			}
		}

		if len(ingMap) == toi && totalCat == toc {
			arr := make([]int, 0)
			for k := range ingMap {
				arr = append(arr, k)

				delete(ingMap, k)
				if strings.Contains(ings[k], "fat") {
					catArr[0]--
				} else if strings.Contains(ings[k], "carb") {
					catArr[1]--
				} else if strings.Contains(ings[k], "protein") {
					catArr[2]--
				} else if strings.Contains(ings[k], "fiber") {
					catArr[3]--
				} else if strings.Contains(ings[k], "seasoning") {
					catArr[4]--
				}
			}

			sort.Ints(arr)

			for idx, ingsIdx := range arr {
				if idx < len(arr)-1 {
					output += ings[ingsIdx] + ":"
				} else {
					output += ings[ingsIdx]
				}
			}
		} else {
			output += "#"
		}
	}

	return output
}

// https://coderanch.com/t/726461/java/Solving-complex-programming
// With above question, the other requirement was chef has 5 categories
// each ingredient category expires in following times:
// fat - 1, carb - 2, protein-3, fiber-3, seasoning-4

// i/p = p1, s1, c1, s2, f1, f2, p2
// for cat store, arr
// for expire store, map key name/index => expiry
func chefCooking() {
	ings := []string{"fat1", "fiber1", "carb1", "fat2", "fiber2"}

	toi := 3
	toc := 3

	fmt.Println(evaluateCookingDay(ings, toi, toc))

}
