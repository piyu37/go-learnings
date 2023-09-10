package main

import (
	"encoding/json"
	"log"
)

type node struct {
	Name     string  `json:"name"`
	Children []*node `json:"children,omitempty"`
}

func parse(v string) (*node, error) {
	root := &node{}

	parseChildren(v, root)

	return root, nil
}

func parseChildren(v string, n *node) {
	if v == "" {
		return
	}

	name := ""
	i := 0
	// getting name(parent) string from the string
	if v[i] != '[' {
		for ; i < len(v); i++ {
			if v[i] == '[' {
				break
			}

			name += string(v[i])
		}
	}

	n.Name = name

	arr := []string{}
	k := i + 1
	// bracket to count open & close brackets
	bracket := 0
	childString := ""

	// spliting the string into array if we found ','
	for i = i + 1; i < len(v)-1; i++ {
		if v[i] == '[' {
			bracket++
		}
		if v[i] == ']' {
			bracket--
		}
		if v[i] == ',' && bracket == 0 {
			childString = v[k:i]
			arr = append(arr, childString)
			k = i + 1
		}
	}

	// this will become true when we have no brackets & comma in string(completing the recursion)
	if k == i {
		return
	}

	childString = v[k:i]
	arr = append(arr, childString)

	// calling parseChildren for each string in array
	for i := 0; i < len(arr); i++ {
		n.Children = append(n.Children, &node{})

		parseChildren(arr[i], n.Children[i])
	}
}

var examples = []string{
	"[a,b,c]",
	"[a[aa[aaa],ab,ac],b,c[ca,cb,cc[cca]]]",
}

func string_tree_conversion() {
	for i, example := range examples {
		result, err := parse(example)
		if err != nil {
			panic(err)
		}
		j, err := json.MarshalIndent(result, " ", " ")
		if err != nil {
			panic(err)
		}

		log.Printf("Example %d: %s - %s", i, example, string(j))
	}
}
