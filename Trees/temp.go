package main

import "fmt"

var pl = fmt.Println

type node struct {
	data  int
	left  *node
	right *node
}
type tree struct {
	root *node
}

func (t *tree) insert(data int) {
	if t.root == nil {
		t.root = &node{data, nil, nil}
		return
	}
	current := t.root
	for current != nil {
		// left child
		if data < current.data {
			if current.left == nil {
				current.left = &node{data, nil, nil}
				return
			}
			current = current.left
		}
		//right child
		if data > current.data {
			if current.right == nil {
				current.right = &node{data, nil, nil}
				return
			}
			current = current.right
		}
	}

}

func (t *tree) delete(key int) {
	if t.root == nil {
		return
	}
	Delete(t.root, key)
}
func Delete(n *node, key int) *node {

	if n == nil {
		return n
	}

	if key < n.data {
		n.left = Delete(n.left, key)
	}
	if key > n.data {
		n.right = Delete(n.right, key)
	}

	if n.data == key {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil && n.right != nil {
			return n.right
		}
		if n.left != nil && n.right == nil {
			return n.left
		}
		smallest := nodeHelper(n.right)
		n.data = smallest.data
		n.right = Delete(smallest.right, key)

	}
	return n

}

func nodeHelper(n *node) *node {
	left := n
	for left != nil && left.left != nil {
		left = left.left
	}
	return left
}

func Init() *tree {
	return &tree{nil}
}
func (t *tree) traverse() {
	Inorder(t.root)
}
func Inorder(t *node) {
	if t == nil {
		return
	}
	Inorder(t.left)
	pl(t.data)
	Inorder(t.right)

}

func main() {
	Tree := Init()

	Tree.insert(33)
	Tree.insert(31)
	Tree.insert(3)
	Tree.insert(43)

	Tree.traverse()

	Tree.delete(33)

	pl("After deletion:")

	Tree.traverse()

}
