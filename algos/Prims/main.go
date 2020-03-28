package main

import "fmt"

func main() {
	var n, v, x, y int
	fmt.Printf("\n Enter number of nodes :  ")
	fmt.Scan(&n)
	var adj = make([][]int, 0)
	for i := 0; i < n; i++ {
		var row = make([]int, 0)
		for j := 0; j < n; j++ {
			fmt.Printf("\n Enter path from %v to %v :", i, j)
			fmt.Scan(&v)
			if v == 0 {
				v = 999
			}
			row = append(row, v)
		}
		adj = append(adj, row)
	}
	var visited = make([]int, n, n)
	visited[0] = 1

	ne := 0

	for ne < n-1 {
		min := 999
		x = 0
		y = 0
		for i := 0; i < n; i++ {
			if visited[i] == 1 {
				for j := 0; j < n; j++ {
					if visited[j] == 0 && adj[i][j] < min {
						min = adj[i][j]
						x = i
						y = j
					}
				}
			}
		}
		ne++
		visited[y] = 1
		fmt.Printf("\n edge %v , %v --> %v , weight: %v \n", ne, x+1, y+1, min)

	}

}
