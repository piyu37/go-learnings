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
