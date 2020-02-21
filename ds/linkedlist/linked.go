package linkedlist

import "fmt"

type node struct {
	data interface{}
	next *node
}

type listOps interface {
	createList()
	insertNode()
	deletedNode()
	findElement()
}

func (start node) createList() *node {
	start.data = 0
	start.next = nil
	return &start
}

func (start *node) insertNode(n int) {
	var newNode *node
	newNode.data = n
	newNode.next = nil
	start = newNode
	fmt.Printf("--data inserted--")

}
func (start *node) deletedNode(a int) {
	var n, p *node
	n = start
	for n.next != nil {
		p.next = n
		if n.data == a {
			temp := n.next
			p.next = temp
		}
		n = n.next
	}
}
