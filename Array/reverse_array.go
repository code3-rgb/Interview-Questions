package main

import "fmt"

var pl = fmt.Println

func main() {

	var array = []int{1, 2, 3, 4, 5, 6, 7}
	size := len(array)
	start := 0
	end := size - 1

	for start < end {
		temp := array[start]
		array[start] = array[end]
		array[end] = temp
		start++
		end--
	}

	pl(array)

}
