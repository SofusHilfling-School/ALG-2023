package main

// type TwoThreeTreeNode struct {
// 	IsLeaf bool
// 	Keys   []int
// 	Child  []*TwoThreeTreeNode
// }

// type TwoThreeTree struct {
// 	Root *TwoThreeTreeNode
// }

// func NewNode(isLeaf bool) *TwoThreeTreeNode {
// 	return &TwoThreeTreeNode{
// 		IsLeaf: isLeaf,
// 		Keys:   make([]int, 0),
// 		Child:  make([]*TwoThreeTreeNode, 0),
// 	}
// }

// func NewTwoThreeTree() *TwoThreeTree {
// 	return &TwoThreeTree{
// 		Root: nil,
// 	}
// }

// func (t *TwoThreeTree) Search(key int) bool {
// 	return t.searchRecursive(t.Root, key)
// }

// func (t *TwoThreeTree) searchRecursive(node *TwoThreeTreeNode, key int) bool {
// 	if node == nil {
// 		return false
// 	}

// 	// loop through all keys in node
// 	for i := 0; i < len(node.Keys); i++ {
// 		// if key are equal return true (match found)
// 		// else if less check child at same position
// 		//   first key = left child
// 		//   second key = middle child
// 		if key == node.Keys[i] {
// 			return true
// 		} else if key < node.Keys[i] {
// 			if node.IsLeaf {
// 				return false
// 			}
// 			return t.searchRecursive(node.Child[i], key)
// 		}
// 	}

// 	// make sure there exist a child
// 	// if it does, it must be the last child
// 	if node.IsLeaf {
// 		return false
// 	}
// 	return t.searchRecursive(node.Child[len(node.Keys)], key)
// }

// func (t *TwoThreeTree) Insert(key int) {
// 	// if first element in tree add it as the root and return
// 	if t.Root == nil {
// 		// the root is also a leaf, since it has no children
// 		t.Root = NewNode(true)
// 		t.Root.Keys = append(t.Root.Keys, key)
// 		return
// 	}

// 	// // if root is a three node we must split it
// 	if len(t.Root.Keys) == 2 {
// 		// the root has children, so its not a leaf now
// 		newRoot := NewNode(false)
// 		// add the previous root as a child of the new
// 		newRoot.Child = append(newRoot.Child, t.Root)
// 		// update the root
// 		t.Root = newRoot
// 		// split the child (its a 3-node) at index 0 (the only child)
// 		t.splitChild(newRoot, 0)
// 	}

// 	t.insertNonFull(t.Root, key)
// }

// func (t *TwoThreeTree) splitChild(node *TwoThreeTreeNode, index int) {
// 	// get the child at index specified
// 	child := node.Child[index]
// 	newChild := NewNode(child.IsLeaf)

// 	// Move the third key and child to the new node
// 	newChild.Keys = append(newChild.Keys, child.Keys[2])
// 	child.Keys = child.Keys[:2]
// 	if !child.IsLeaf {
// 		newChild.Child = append(newChild.Child, child.Child[2])
// 		child.Child = child.Child[:2]
// 	}

// 	node.Keys = append(node.Keys, 0)
// 	node.Child = append(node.Child, nil)

// 	// Shift keys and children to the right to make space for newChild
// 	for i := len(node.Keys) - 1; i > index; i-- {
// 		node.Keys[i] = node.Keys[i-1]
// 	}
// 	for i := len(node.Child) - 1; i > index+1; i-- {
// 		node.Child[i] = node.Child[i-1]
// 	}

// 	// Insert newChild into node
// 	node.Keys[index] = child.Keys[1]
// 	node.Child[index+1] = newChild

// 	// Update the parent pointers of newChild's children
// 	if !newChild.IsLeaf {
// 		newChild.Child[0].Child[0] = newChild.Child[0]
// 		newChild.Child[1].Child[0] = newChild.Child[1]
// 	}

// 	child.Keys = child.Keys[:1]
// 	child.Child = child.Child[:2]
// }

// func (t *TwoThreeTree) insertNonFull(node *TwoThreeTreeNode, key int) {
// 	index := len(node.Keys) - 1

// 	// if node has no children add key to keys
// 	// we are sure that
// 	if node.IsLeaf {
// 		node.Keys = append(node.Keys, key)
// 		// swap
// 		if len(node.Keys) == 2 && node.Keys[0] > node.Keys[1] {
// 			node.Keys[0], node.Keys[1] = node.Keys[1], node.Keys[0]
// 		}
// 	} else {
// 		for index >= 0 && key < node.Keys[index] {
// 			index--
// 		}
// 		index++
// 		if len(node.Child[index].Keys) == 3 {
// 			t.splitChild(node, index)
// 			if key > node.Keys[index] {
// 				index++
// 			}
// 		}
// 		t.insertNonFull(node.Child[index], key)
// 	}
// }
