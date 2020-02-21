package queue

import "fmt"

type queue struct {
	data []int
}
type fqueue interface {
	createQueue()
	enqueue()
	dequeue()
}

// Initialize new queue
func (q queue) createQueue() *queue {
	return &q
}

// Enqueue -- Inserts element in Queue --- First in First out
func (q *queue) enqueue(n int) *queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue -- Remove element from Queue --- First in First Out
func (q *queue) dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
