package main

import "fmt"

var pl = fmt.Println

type node struct {
	data  int
	left  *node
	right *node
}

type Tree struct {
	tree *node
}

func (t *Tree) Insert(data int) {
	if t.tree == nil {
		t.tree = &node{data, nil, nil}
		return
	}
	insert(data, t.tree)
}
func insert(data int, n *node) {
	if data < n.data {
		if n.left == nil {
			n.left = &node{data, nil, nil}
			return
		}
		insert(data, n.left)
	} else if data > n.data {
		if n.right == nil {
			n.right = &node{data, nil, nil}
			return
		}
		insert(data, n.right)
	}

}

func (t *Tree) Delete(data int) {
	if t.tree == nil {
		return
	}
	delete(data, t.tree)
}
func delete(data int, n *node) *node {
	if nil == n {
		return n
	}
	if n.data > data {
		n.left = delete(data, n.left)
	}
	if n.data < data {
		n.right = delete(data, n.right)
	}

	if data == n.data {
		if n.left == nil && n.right == nil {
			n = nil
			return n
		}
		if n.left == nil && n.right != nil {
			temp := n.right
			n = nil
			n = temp
			return n
		}
		if n.left != nil && n.right == nil {
			temp := n.left
			n = nil
			n = temp
			return n
		}

		tempNode := min_value(n.right)
		n.data = tempNode.data
		n.right = delete(data, n.right)

	}
	return n
}

func min_value(n *node) *node {
	temp := n
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}

func max_depth(n *node) int {
	if n == nil {
		return 0
	}
	if n.left == nil && n.right == nil {
		return 1
	}
	l := max_depth(n.left)
	r := max_depth(n.right)

	if l >= r {
		return l + 1
	} else {
		return r + 1
	}

}

func inorder(n *node) {
	if n == nil {
		return
	}
	pl(n.data)
	inorder(n.left)
	inorder(n.right)
}

func main() {
	t := Tree{}

	t.Insert(12)
	t.Insert(15)
	t.Insert(5)
	t.Insert(3)
	t.Insert(17)
	t.Insert(19)
	t.Insert(7)
	t.Insert(9)

	pl("AFTER deletion:")
	inorder(t.tree)

	pl()
	pl("AFTER deletion:")

	t.Delete(7)
	t.Delete(17)
	t.Delete(9)
	t.Delete(12)

	inorder(t.tree)
	h := max_depth(t.tree)

	pl("height is:", h)

}
