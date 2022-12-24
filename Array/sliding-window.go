package main

import "fmt"

var pl = fmt.Println

func subarry_sum(arr []int, k int) {
	j := 0
	total := 0
	count := 0
	ar := []int{}

	if k > len(arr) || k < 1 {
		return
	}

	for i := 0; i < len(arr); i++ {
		total += arr[i]
		// pl(total)

		if count == k {
			ar = append(ar, total)
			total = 0
			j++
			i = j - 1
			pl(i)
			count = 0
		} else {
			count++
		}
	}
	pl(ar)
}

func main() {

	arr := []int{3, 3, 5, 2, 9, 7, 1}

	subarry_sum(arr, 2)
}
