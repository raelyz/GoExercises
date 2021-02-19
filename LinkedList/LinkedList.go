package main

import "fmt"

func main() {

	linkList := LinkedList{
		head: nil,
		size: 0,
	}
	linkList.printAllNodes()
	linkList.addNode("first")
	linkList.addNode("second")
	linkList.addNode("third")
	linkList.addNode("fourth")
	linkList.addNode("5")
	linkList.addNode("6")
	linkList.printAllNodes()
	linkList.reverseI()
	linkList.printAllNodes()
	linkList.head = linkList.reverseR(linkList.head)
	linkList.printAllNodes()
}

type Node struct {
	item string
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (p *LinkedList) addNode(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.head == nil {
		p.head = newNode
	} else {
		currNode := p.head
		for currNode.next != nil {
			currNode = currNode.next
		}
		currNode.next = newNode
	}
	p.size++
	return nil
}

func (p *LinkedList) removeNode(index int) error {
	currNode := p.head
	nextNode := currNode.next
	var position int = 1
	if position == index {
		p.head = nextNode
		return nil
	}
	for position+1 != index {
		position++
		currNode = currNode.next
		nextNode = currNode.next
	}
	currNode.next = nextNode.next
	p.size--
	return nil
}

func (p *LinkedList) printAllNodes() error {
	if p.head == nil {
		fmt.Println("list is empty")
	}
	var index int
	currNode := p.head
	for currNode != nil {
		fmt.Println("Item", index, currNode.item)

		currNode = currNode.next
		index++
	}
	fmt.Println("LinkedList of size", p.size)
	return nil
}

func (p *LinkedList) get(index int) error {
	currNode := p.head
	var position int = 1
	for position != index {
		position++
		currNode = currNode.next
	}
	fmt.Println(currNode.item)
	return nil
}

func (p *LinkedList) reverseI() error {
	if p.size < 2 {
		return nil
	}
	currNode := p.head
	nextNode := currNode.next
	currNode.next = nil

	for nextNode != nil {
		nextnext := nextNode.next
		nextNode.next = currNode
		currNode = nextNode
		nextNode = nextnext
	}
	p.head = currNode
	return nil

}

func (p *LinkedList) reverseR(node *Node) *Node {
	if node.next == nil || node == nil {
		return node
	}
	newNode := p.reverseR(node.next)
	node.next.next = node
	node.next = nil
	return newNode
}
