package main

import (
	"fmt"
	"sort"
)

var pl = fmt.Println

func kth_element(arr []int, k int) int {

	sort.Ints(arr)

	return arr[k-1]

}

func main() {
	var arr []int = []int{7, 10, 4, 3, 20, 15}

	kth := kth_element(arr, 4)

	pl(kth)

}
