package main

import "fmt"

func myOnCallDays(T int32, K int32, D int32) []int32 {
	result := make([]int32, 0)

	var tCount int32 = 0
	var i int32
	var originalK int32 = K
	var reflctionK int32 = K
	originalFlag := true

	for i = 1; i <= D; i++ {
		if originalFlag {
			tCount++

			if tCount == originalK {
				result = append(result, i)
			}
		} else {
			tCount--

			if tCount == reflctionK {
				result = append(result, i)
			}
		}

		if originalFlag && tCount == T {
			tCount++
			originalFlag = false
		}

		if !originalFlag && tCount == 1 {
			tCount--
			originalFlag = true
		}
	}

	return result
}

// onCallDays simulates on-call days for two teams, toggling between the original and
// reflection teams and tracking the consecutive on-call days for each team
// T = The total number of days in the simulation/exchange between teams.
// K = The number of consecutive days on call for each team.
// D = The total number of on-call days for each team.
// T=1, K=2, D=10 => 1 4 5 8 9
// T=2, K=2, D=10 => 2 3 6 7 10
func onCallDays() {
	T := 2
	K := 2
	D := 10

	fmt.Println(myOnCallDays(int32(T), int32(K), int32(D)))
}
