package main

import "fmt"

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type OrgReport struct {
	result *OrgChart
	count  int
}

/*
 * COMPLEXITY ANALYSIS
 *
 * TIME COMPLEXITY: O(N)
 * ----------------------------
 * Where N is the total number of people (nodes) in the organization chart.
 *
 * The total time is calculated by: (Number of states explored) * (Work done per state)
 *
 * 1. Number of States -> O(N)
 * - The algorithm performs a standard Depth-First Search (DFS) traversal
 * on the N-ary tree.
 * - In the worst-case scenario (where the two target reports are at the
 * very bottom of the hierarchy or in completely different branches), the
 * algorithm must visit every single node in the organization exactly once
 * to tally up the important reports.
 *
 * 2. Work Per State -> O(1)
 * - Inside the recursive `lcmHelper` function, the algorithm checks pointer
 * equality (`org == reportOne`), increments a counter, and returns a small
 * `OrgReport` struct.
 * - Note for the interviewer: While there is a `for` loop iterating over
 * the `DirectReports`, across the *entire* execution of the algorithm,
 * each manager-to-report link is traversed exactly once. Therefore, the
 * actual processing work done for each specific node evaluates to O(1)
 * constant time.
 *
 * Total Time = O(N) states * O(1) work per state = O(N)
 *
 * SPACE COMPLEXITY: O(H) or O(D)
 * ----------------------------
 * Where H (or D) is the maximum height (or depth) of the organization chart.
 *
 * 1. Recursion Stack: The algorithm explores the chart branch by branch. The
 * maximum depth of the recursion tree is perfectly equal to the longest
 * chain of command in the org chart. In a fairly flat/balanced organization,
 * this is O(log N). In the absolute worst case (a single vertical chain of
 * management), the stack goes N levels deep. Stating O(H) is the most
 * accurate and interview-appropriate answer.
 * 2. Auxiliary Space: The algorithm does not allocate any dynamic arrays,
 * slices, or hash maps. The only memory used is the fixed-size `OrgReport`
 * struct created during each stack frame, which takes O(1) auxiliary space.
 *
 * Total Space = Stack O(H) + Auxiliary O(1) = O(H)
 */
func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return lcmHelper(org, reportOne, reportTwo).result
}

func lcmHelper(org, reportOne, reportTwo *OrgChart) *OrgReport {
	numImpReports := 0

	for _, o := range org.DirectReports {
		or := lcmHelper(o, reportOne, reportTwo)
		if or.result != nil {
			return or
		}

		numImpReports += or.count
	}

	if org == reportOne || org == reportTwo {
		numImpReports++
	}

	var lcm *OrgChart

	if numImpReports == 2 {
		lcm = org
	}

	return &OrgReport{
		result: lcm,
		count:  numImpReports,
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/29_lowest_common_manager
func lowestCommonManager() {
	r1 := &OrgChart{
		Name: "E",
	}

	r2 := &OrgChart{
		Name: "I",
	}

	oc := &OrgChart{
		Name: "A",
		DirectReports: []*OrgChart{
			{
				Name: "B",
				DirectReports: []*OrgChart{
					{
						Name: "D",
						DirectReports: []*OrgChart{
							{
								Name: "H",
							},
							r2,
						},
					},
					r1,
				},
			},
			{
				Name: "C",
				DirectReports: []*OrgChart{
					{
						Name: "F",
					},
					{
						Name: "G",
					},
				},
			},
		},
	}

	fmt.Println(GetLowestCommonManager(oc, r1, r2))
}
