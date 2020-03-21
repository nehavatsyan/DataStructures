package main

import "fmt"

var cost int

func main() {
	var n, s, v int
	fmt.Print("Enter number of nodes ---> ")
	fmt.Scan(&n)
	fmt.Print("\n Enter start point ---> ")
	fmt.Scan(&s)
	var dist = make([][]int, 0, 0)
	for i := 0; i < n; i++ {
		var a = make([]int, 0, 0)
		for j := 0; j < n; j++ {
			fmt.Printf("\n Enter cost %v ---> %v =  ", i+1, j+1)
			fmt.Scan(&v)
			a = append(a, v)
		}
		dist = append(dist, a)
	}

	fmt.Printf("\n Entered cost matrix is : \n")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%v ", dist[i][j])
		}
		fmt.Println()
	}
	var city = make([]int, n)
	path(n, s-1, dist, city)
	fmt.Printf("\nThe total cost is %v \n", cost)
}

func path(n int, s int, dist [][]int, city []int) {
	city[s] = 1
	fmt.Printf("%v ->", s+1)
	ncity := least(n, s, dist, city)
	if ncity == 999 {
		ncity = 0
		fmt.Printf("%v", ncity+1)
		cost = cost + dist[s][ncity]
		//fmt.Println(cost)
		return
	}

	path(n, ncity, dist, city)

}
func least(n int, s int, dist [][]int, city []int) int {
	nc := 999
	min := 999
	var kmin int

	for i := 0; i < n; i++ {
		if dist[s][i] != 0 && city[i] == 0 {
			if dist[s][i]+dist[i][s] < min {
				min = dist[i][s] + dist[s][i]
				kmin = dist[s][i]
				nc = i
			}
		}
	}

	if min != 999 {
		//fmt.Printf("\n%v\n", kmin)
		cost = cost + kmin
	}
	return nc
}
