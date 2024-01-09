package main

import "fmt"

func getTheGroups(n int32, queryType []string, students1 []int32, students2 []int32) []int32 {
	friendGroupMap := make(map[int32]*[]int32)
	result := make([]int32, 0)

	for i, v := range queryType {
		s1 := students1[i]
		s2 := students2[i]

		if v == "Friend" {
			_, ok1 := friendGroupMap[s1]
			_, ok2 := friendGroupMap[s2]

			if ok1 || ok2 {
				if ok1 && ok2 {
					*friendGroupMap[s1] = append(*friendGroupMap[s1], *friendGroupMap[s2]...)
					friendGroupMap[s2] = friendGroupMap[s1]
				} else if ok1 {
					*friendGroupMap[s1] = append(*friendGroupMap[s1], s2)
					friendGroupMap[s2] = friendGroupMap[s1]
				} else {
					*friendGroupMap[s2] = append(*friendGroupMap[s2], s1)
					friendGroupMap[s1] = friendGroupMap[s2]
				}
			} else {
				friendGroupMap[s1] = &[]int32{s1, s2}
				friendGroupMap[s2] = friendGroupMap[s1]
			}
		} else {
			_, okS1 := friendGroupMap[s1]
			if !okS1 {
				friendGroupMap[s1] = &[]int32{s1}
			}

			_, okS2 := friendGroupMap[s2]
			if !okS2 {
				friendGroupMap[s2] = &[]int32{s2}
			}

			total := len(*friendGroupMap[s1]) + len(*friendGroupMap[s2])

			result = append(result, int32(total))
		}
	}

	return result
}

// WAP to return length of friends whenever "Total" gets encounter;
// create/merge friend list when "Friend" come in qt
// e.g. qt = friend; 30 is friend of 62 & viceversa
// qt = total; if exist then return length of map; else create new friend list from s1 & s2 and return 2
func getTheGroupsMain() {
	var n int32 = 100
	qt := []string{"Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total", "Friend", "Total"}
	s1 := []int32{30, 49, 33, 79, 61, 13, 73, 1, 23, 14, 22, 1, 38, 82, 80, 75, 45, 64, 3, 19, 67, 68, 40, 10, 53, 37, 2, 47, 75, 49, 83, 63, 43, 8, 70, 37, 44, 41, 3, 34, 36, 23, 75, 13, 29, 53, 11, 60, 55, 22, 4, 48, 44, 21, 82, 69, 18, 77, 64, 19, 36, 61, 77, 65, 69, 4, 21, 78, 59, 43, 32, 30, 84, 70, 77, 17, 28, 21, 60, 67, 5, 44, 12, 57}
	s2 := []int32{62, 19, 5, 83, 56, 78, 18, 35, 72, 33, 21, 8, 54, 6, 57, 46, 40, 10, 61, 51, 42, 51, 12, 47, 52, 67, 26, 23, 17, 52, 3, 81, 8, 68, 41, 52, 83, 48, 31, 62, 44, 38, 1, 71, 56, 33, 8, 37, 24, 26, 57, 20, 22, 14, 51, 40, 47, 20, 23, 27, 20, 13, 32, 78, 82, 51, 77, 65, 52, 5, 15, 69, 69, 7, 58, 20, 12, 28, 31, 76, 82, 64, 4, 29}
	fmt.Println(getTheGroups(n, qt, s1, s2))
}
