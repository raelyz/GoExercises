package main

import (
	"fmt"
	"strings"
)

func main() {
	stack := Stack{
		head: nil,
		size: 0,
	}
	mapper := make(map[string]string)
	mapper["{"] = "}"
	mapper["["] = "]"
	mapper["("] = ")"
	test := "{}{}{}{(})}"
	for _, item := range strings.Split(test, "") {

		if item == "}" || item == "]" || item == ")" {
			top := mapper[stack.pop().item]
			if top != item {
				fmt.Println("invalid string")
				break
			}
		} else {
			stack.push(item)
		}
	}
	stack.push("first")
	stack.push("2nd")
	stack.push("3")
	stack.push("4")
	stack.printAllNodes()
}

type Node struct {
	item string
	next *Node
}

type Stack struct {
	head *Node
	size int
}

func (s *Stack) push(item string) error {
	newItem := &Node{
		item: item,
		next: s.head,
	}
	if s.head == nil {
		s.head = newItem
	} else {
		newItem.next = s.head
		s.head = newItem
	}
	s.size++
	return nil
}

func (s *Stack) printAllNodes() error {
	if s.size == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	curr := s.head
	for curr != nil {
		fmt.Println(curr.item)
		curr = curr.next
	}
	return nil
}

func (s *Stack) pop() *Node {
	if s.size == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	pop := s.head
	s.head = s.head.next
	s.size--
	return pop
}
