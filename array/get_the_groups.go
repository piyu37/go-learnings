package main

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
