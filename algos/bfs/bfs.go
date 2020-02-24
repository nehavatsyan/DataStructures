package main

import "fmt"

var front, rear int = -1, -1
var queue = make([]int, 10)

func main() {
	var n, num, v int // n is the size of the dimension of the matrix
	fmt.Println("Enter number of vertex")
	fmt.Scanf("%d", &n)
	visit := make([]int, 0, 10)
	adj := make([][]int, 0, 10) // the adjacency matrix
	row := make([]int, 0, 10)   // each row of the matrix
	// Input for the adjacency matrix
	fmt.Println("Enter 1 if edge exists else 0")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("\n%v ----> %v", i, j)
			fmt.Scanf("%d", &num)
			row = append(row, num)
		}
		adj = append(adj, row)
		row = make([]int, 0, 10)
	}
	// Initializing the visit matrix to zero
	for i := 0; i < n; i++ {
		visit = append(visit, 0)
	}
	fmt.Println("Enter start vertex")
	fmt.Scan(&v)
	// Printing the number of components
	callBfs(adj, visit, v, n)
}

func callBfs(adj [][]int, visit []int, v int, n int) {

	pushQueue(v)
	for isEmpty() == 0 {
		v = popQueue()
		if visit[v] == 1 {
			continue
		}
		fmt.Printf("%v -->", v)
		visit[v] = 1
		for i := 0; i < n; i++ {
			if adj[v][i] == 1 && visit[i] == 0 {
				pushQueue(i)
			}
		}
	}
	fmt.Println()
}

// <--------------------------Push element into queue----------------------->

func pushQueue(v int) {
	if rear == 9 {
		fmt.Println("Queue Overflow")
	} else {
		if front == -1 {
			front = 0
		}
		rear = rear + 1
		queue[rear] = v
	}
}

// <-------------------------- check if queue is empty----------------------->

func isEmpty() int {
	if front == -1 || front > rear {
		return 1
	}
	return 0
}

// <-------------------------- Pop element of queue ----------------------->
func popQueue() int {
	if front == -1 || front > rear {
		fmt.Printf("Queue Underflow\n")
		return 0
	}

	deleteItem := queue[front]
	front = front + 1
	return deleteItem
}

/* func bfs(a [][]int, v int, visit []int) {
	visit[v] = 1
	for w := 0; w < len(a); w++ {
		if a[v][w] == 1 && visit[w] == 0 {
			r++
			queue = append(queue, w)
		}
	}
	for f <= r {
		adjv := queue[f]
		f++
		bfs(a, adjv, visit)
	}
} */
