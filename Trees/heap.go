package main

import "fmt"

var pl = fmt.Println
var arr = []int{}

func insert(k int) {
	arr = append(arr, k)
	size := len(arr) - 1
	heapifyUp(size)

}

func heapifyUp(index int) {
	for hasParent(index) && arr[getParent(index)] > arr[index] {
		swap(getParent(index), index)
		heapifyUp(getParent(index))
	}
}

func hasParent(index int) bool {
	return getParent(index) >= 0
}
func getParent(index int) int {
	return index / 2
}

func swap(index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp

}

func main() {

	insert(16)
	insert(54)
	insert(19)
	insert(67)
	insert(87)

	pl(arr)

}
