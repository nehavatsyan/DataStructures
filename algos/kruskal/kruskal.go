package main

import (
	"fmt"
)

var v, num, a, b, mincost int
var ne int = 0
var visit = make([]int, 10)

func main() {
	fmt.Println("----- Kruskal Algorithm Implementation -----")
	fmt.Println(" Enter number of vertices ")
	fmt.Scan(&v) //get number of vertices
	fmt.Println(" Enter weight if edge present else 0")
	adj := make([][]int, 0, 10) // the adjacency matrix
	row := make([]int, 0, 10)

	// Get data

	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			fmt.Printf("%v ----> %v ", i, j)
			fmt.Scanf("%d", &num)
			if num == 0 {
				num = 999
			}
			row = append(row, num)
		}
		fmt.Println()
		adj = append(adj, row)
		row = make([]int, 0, 10)
	}

	fmt.Println("The edges of minimum cost spanning tree are ")
	min := 999
	for ne < v {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if adj[i][j] < min {
					min = adj[i][j]
					a = i
					b = j
				}
			}
		}

		ne++
		u := find(a)
		v := find(b)
		//fmt.Print(u, v)
		ne = ne + 1
		adj[a][b] = 999
		adj[b][a] = 999
		//fmt.Print(uni(u, v))
		if uni(u, v) > 0 {
			fmt.Printf("%v edge (%v,%v) =%v\n", ne, a, b, min)
			mincost += min
		}

	}
	fmt.Printf("Minimum cost is %v", mincost)

}
func find(i int) int {
	if visit[i] > 0 {
		i = visit[i]
	}
	//fmt.Print(i)
	return i
}
func uni(i int, j int) int {
	if i != j {
		visit[j] = i
		return 1
	}
	return 0
}
