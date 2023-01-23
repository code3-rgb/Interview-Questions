package main

import "fmt"

var pl = fmt.Println

type node struct {
	store []int
	size  int
}

func (n *node) removeMin() int {

	// if n.store == nil {
	// 	return 0
	// }

	last := len(n.store) - 1

	n.swap(0, last)
	deletedNode := n.store[last]
	aux := n.store[:last]
	n.store = aux
	n.size -= 1

	n.heapifyDown(0)

	return deletedNode
}
func (n *node) heapifyDown(index int) {
	for n.RightChildExist(index) {
		smaller := getLeftChildIndex(index)
		left := n.store[getLeftChildIndex(index)]
		right := n.store[getRightChildIndex(index)]
		if n.RightChildExist(index) && right < left {
			smaller = getRightChildIndex(index)
		}
		if n.store[index] < n.store[smaller] {
			break
		} else {
			n.swap(index, smaller)
		}
		index = smaller
	}
}

func (n *node) insert(data int) {
	n.store = append(n.store, data)
	n.size++
	n.heapifyUp(n.size - 1)
}
func (n *node) heapifyUp(index int) {
	parent := getParent(index)
	for n.parentExist(index) && n.store[index] > n.store[parent] {
		n.swap(parent, index)
		n.heapifyUp(parent)
	}
}

func (n *node) swap(index1, index2 int) {
	temp := n.store[index1]
	n.store[index1] = n.store[index2]
	n.store[index2] = temp
}

func (n *node) parentExist(index int) bool {
	return n.store[index] >= 0
}
func getParent(index int) int {
	return index / 2
}

func (n *node) LeftChildExist(index int) bool {
	return getLeftChildIndex(index) <= len(n.store)-1
}
func (n *node) RightChildExist(index int) bool {
	return getRightChildIndex(index) <= len(n.store)-1
}

func getLeftChildIndex(index int) int {
	return index*2 + 1
}
func getRightChildIndex(index int) int {
	return index*2 + 2
}

func Init() *node {
	heap := &node{}
	return heap
}

func (heap *node) minify() []int {
	temp := []int{}
	size := heap.size
	for i := 0; i < size; i++ {
		val := heap.removeMin()
		pl(val)
		temp = append(temp, val)
	}
	return temp
}

func main() {

	heap := Init()

	heap.size = 0

	heap.insert(3)
	heap.insert(2)
	heap.insert(10)
	heap.insert(4)

	pl(heap.store)

	pl(heap.removeMin())

	pl(heap.store)

	pl(heap.removeMin())

	pl(heap.store)

	pl(heap.removeMin())

	pl(heap.store)

	pl(heap.removeMin())

	pl(heap.store)

}
