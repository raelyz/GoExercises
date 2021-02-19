package main

import "fmt"

func main() {
	queue1 := Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
	queue1.enqueue("first", 0)
	queue1.enqueue("second", 1)
	queue1.enqueue("third", 2)
	queue1.enqueue("fourth", 1)
	queue1.printNodes()
}

type Node struct {
	item     string
	next     *Node
	priority int
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func (q *Queue) enqueue(item string, prio int) error {
	newNode := &Node{
		item:     item,
		next:     nil,
		priority: prio,
	}
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		currNode := q.head
		for currNode.next.priority >= prio {
			currNode = currNode.next
		}
		newNode.next = currNode.next
		currNode.next = newNode
		if newNode.next == nil {
			q.tail = newNode
		}
	}

	q.size++
	return nil
}

func (q *Queue) dequeue() (string, error) {
	if q.head == nil {
		fmt.Println("empty queue")
	}
	frontItem := q.head.item
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		fmt.Println(q.head.item)
		q.head = q.head.next
	}
	q.size--
	return frontItem, nil
}

func (q *Queue) printNodes() error {
	currNode := q.head
	for currNode != nil {
		fmt.Println(currNode.item)
		currNode = currNode.next
	}
	return nil
}
