package main

import "fmt"

var pl = fmt.Println

func search(arr []int, left int, right int, data int) {

	if left > right {
		pl("Not found")
		return
	}
	mid := (left + right) / 2
	if arr[mid] == data {
		pl("Found at ", mid, " index")
		return
	}
	if data < arr[mid] {
		search(arr, left, mid-1, data)
	} else if data > arr[mid] {
		search(arr, mid+1, right, data)
	}

}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	search(arr, 0, len(arr)-1, 9)

}
