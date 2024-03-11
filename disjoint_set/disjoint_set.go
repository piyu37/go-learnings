package main

import "fmt"

type disjointSet struct {
	rank   []int
	parent []int
	size   []int
}

func initDisjointSet(len int) *disjointSet {
	r := make([]int, len)
	p := make([]int, len)
	s := make([]int, len)

	for i := range p {
		p[i] = i
		s[i] = 1
	}

	return &disjointSet{r, p, s}
}

func (ds *disjointSet) findUPar(node int) int {
	if node == ds.parent[node] {
		return node
	}

	parent := ds.findUPar(ds.parent[node])
	ds.parent[node] = parent

	return parent
}

func (ds *disjointSet) unionByRank(u, v int) {
	ulpU := ds.findUPar(u)
	ulpV := ds.findUPar(v)

	if ulpU == ulpV {
		return
	}

	ru := ds.rank[ulpU]
	rv := ds.rank[ulpV]

	if ru == rv {
		ds.rank[ulpU]++
	}

	if ru >= rv {
		ds.parent[ulpV] = ulpU
	} else {
		ds.parent[ulpU] = ulpV
	}
}

func (ds *disjointSet) unionBySize(u, v int) {
	ulpU := ds.findUPar(u)
	ulpV := ds.findUPar(v)

	if ulpU == ulpV {
		return
	}

	su := ds.size[ulpU]
	sv := ds.size[ulpV]

	if su >= sv {
		ds.parent[ulpV] = ulpU
		ds.size[ulpU] += sv
	} else {
		ds.parent[ulpU] = ulpV
		ds.size[ulpV] += su
	}
}

func disjointSetMain() {
	dj := initDisjointSet(8) // here 8 tells no. of nodes(0 to 7); if nodes are not sequentially, then use map
	dj.unionByRank(1, 2)
	dj.unionByRank(2, 3)
	dj.unionByRank(4, 5)
	dj.unionByRank(6, 7)
	dj.unionByRank(5, 6)

	if dj.findUPar(3) == dj.findUPar(7) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	dj.unionByRank(3, 7)

	if dj.findUPar(3) == dj.findUPar(7) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	dj2 := initDisjointSet(8) // here 7 tells no. of nodes; if nodes are not sequentially, then use map
	dj2.unionBySize(1, 2)
	dj2.unionBySize(2, 3)
	dj2.unionBySize(4, 5)
	dj2.unionBySize(6, 7)
	dj2.unionBySize(5, 6)

	if dj2.findUPar(3) == dj2.findUPar(7) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	dj2.unionBySize(3, 7)

	if dj2.findUPar(3) == dj2.findUPar(7) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
