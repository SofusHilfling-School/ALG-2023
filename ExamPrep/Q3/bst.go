package main

import "fmt"

type BSTNode struct {
	Value       int
	Left, Right *BSTNode
}

type BinarySearchTree struct {
	Root *BSTNode
}

// Insert inserts a new value into the BST
func (bst *BinarySearchTree) Insert(value int) {
	newNode := &BSTNode{Value: value}

	// if root is empty, insert new item as a root element
	if bst.Root == nil {
		bst.Root = newNode
		return
	}

	curr := bst.Root
	for {
		// if value is less move down the left subtree
		// else move down the right subtree
		if value < curr.Value {
			// if child is null, no subtree exist so add node at that place
			if curr.Left == nil {
				curr.Left = newNode
				return
			}
			curr = curr.Left
		} else {
			// if child is null, no subtree exist so add node at that place
			if curr.Right == nil {
				curr.Right = newNode
				return
			}
			curr = curr.Right
		}
	}
}

// Search searches for a value in the BST
func (bst *BinarySearchTree) Search(value int) bool {
	curr := bst.Root

	for curr != nil {
		fmt.Println("node:", curr.Value)
		if value == curr.Value {
			return true
		} else if value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return false
}
