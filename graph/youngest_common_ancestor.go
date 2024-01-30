package main

import "fmt"

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	d1Height, d2Height := -1, -1

	d1 := descendantOne
	for d1 != nil {
		d1 = d1.Ancestor
		d1Height++
	}

	d2 := descendantTwo
	for d2 != nil {
		d2 = d2.Ancestor
		d2Height++
	}

	if d1Height > d2Height {
		for d1Height != d2Height {
			descendantOne = descendantOne.Ancestor
			d1Height--
		}
	} else if d1Height < d2Height {
		for d1Height != d2Height {
			descendantTwo = descendantTwo.Ancestor
			d2Height--
		}
	}

	for descendantOne != nil {
		if descendantOne == descendantTwo {
			return descendantOne
		}

		descendantOne = descendantOne.Ancestor
		descendantTwo = descendantTwo.Ancestor
	}

	return nil
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/20_youngest_common_ancestor
func youngestCommonAncestor() {
	A := &AncestralTree{
		Name: "A",
	}

	B := &AncestralTree{
		Name:     "B",
		Ancestor: A,
	}

	I := &AncestralTree{
		Name: "I",
		Ancestor: &AncestralTree{
			Name:     "D",
			Ancestor: B,
		},
	}

	E := &AncestralTree{
		Name:     "E",
		Ancestor: B,
	}

	fmt.Println(GetYoungestCommonAncestor(A, I, E))
}
