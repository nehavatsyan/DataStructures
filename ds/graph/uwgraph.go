package graph

import "fmt"

type uwGraph struct {
	numNodes int
	edges    [][]int
}
type opsUwgraph interface {
	createUwg()
}

func (g uwGraph) createUwg(n int) *uwGraph {
	row := make([]int, n)
	fmt.Println("enter 1 if edge exists else 0 ")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v , %v \n", i, j)
			fmt.Scan(&row[j])
		}
		g.edges = append(g.edges, row)
	}
	return &g
}
