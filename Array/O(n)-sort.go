package main

import "fmt"

var pl = fmt.Println

func sort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp

			i = -1
		}
	}

	pl(arr)

}

func main() {
	var arr []int = []int{7, 10, 4, 3, 20, 15}
	sort(arr)

}
