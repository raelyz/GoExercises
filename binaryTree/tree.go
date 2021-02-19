package main

import "fmt"

func main() {
	data := []int{5, 7, 3, 6, 4, 8, 2, 9, 1}
	tree := BST{nil}
	for _, data := range data {
		tree.insert(data)
	}
	tree.postOrder(&tree.root)
}

type BinaryNode struct {
	item  int
	left  *BinaryNode
	right *BinaryNode
}

type BST struct {
	root *BinaryNode
}

func (bst *BST) insertNode(t **BinaryNode, item int) error {
	if *t == nil {
		newNode := &BinaryNode{
			item:  item,
			left:  nil,
			right: nil,
		}
		*t = newNode
		return nil
	}
	if item < (*t).item {
		bst.insertNode(&((*t).left), item)
	} else {
		bst.insertNode(&((*t).right), item)
	}
	return nil
}
func (bst *BST) insert(item int) {
	bst.insertNode(&bst.root, item)
}

func (bst *BST) inOrder(t **BinaryNode) {
	if (*t) == nil {
		return
	} else {
		bst.inOrder(&(*t).left)
		fmt.Println((*t))
		bst.inOrder(&(*t).right)

	}

}
func (bst *BST) preOrder(t **BinaryNode) {
	if (*t) != nil {
		fmt.Println(&(*t))

		bst.preOrder(&(*t).left)
		bst.preOrder(&(*t).right)

	}

}

func (bst *BST) postOrder(t **BinaryNode) {
	if (*t) == nil {
		return

	} else {

		bst.postOrder(&(*t).left)
		bst.postOrder(&(*t).right)
		fmt.Println((*t))
	}

}
