package main

import (
	"fmt"
)

func main() {
	fmt.Println("----- Implement warshall Algorithm -----")
	fmt.Printf("Enter number of vertices ---->\t")
	var v int
	fmt.Scan(&v)
	var g = make([][]int, v+1)

	// get matrix
	fmt.Printf("\nEnter weights , if there is no path add 999 instead of infinity---->\t")
	for i := 0; i < v; i++ {
		row := make([]int, v+1)
		for j := 0; j < v; j++ {
			fmt.Scan(&row[j])
		}
		g = append(g, row)
	}
	floyds(g, v)

}
func floyds(g [][]int, n int) {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					g[i][j] = 0
				} else {
					g[i][j] = min(g[i][j], g[i][k]+g[k][j])
				}
			}
		}
	}
	fmt.Println(g)
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b

}
