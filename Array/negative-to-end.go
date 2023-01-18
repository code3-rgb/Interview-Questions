package main

import "fmt"

var pl = fmt.Println

func negative_to_end(arr []int) {

	j := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			if i != j {
				temp := arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
			j++
		}
	}
}

func main() {

	arr := []int{-1, 2, -3, 4, 5, 6, -7, 8, 9}
	negative_to_end(arr)

	pl(arr)

}
