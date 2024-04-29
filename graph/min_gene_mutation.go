package main

import "fmt"

func minMutation(startGene string, endGene string, bank []string) int {
	type geneMutation struct {
		gene  string
		steps int
	}

	queue := make([]geneMutation, 0)
	seen := make(map[string]struct{})

	queue = append(queue, geneMutation{startGene, 0})
	seen[startGene] = struct{}{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.gene == endGene {
			return curr.steps
		}

		for _, r := range "ACGT" {
			for i := range curr.gene {
				neighbour := curr.gene[:i] + string(r) + curr.gene[i+1:]
				_, ok := seen[neighbour]
				presentInBank := isGeneInBank(bank, neighbour)

				if !ok && presentInBank {
					queue = append(queue, geneMutation{neighbour, curr.steps + 1})
					seen[neighbour] = struct{}{}
				}
			}
		}
	}

	return -1
}

func isGeneInBank(bank []string, val string) bool {
	for _, v := range bank {
		if v == val {
			return true
		}
	}

	return false
}

// https://leetcode.com/problems/minimum-genetic-mutation/description/?envType=study-plan-v2&envId=top-interview-150
func minGeneMutationMain() {
	start := "AACCGGTT"
	end := "AAACGGTA"

	bank := []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}

	fmt.Println(minMutation(start, end, bank))
}
