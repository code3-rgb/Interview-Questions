package main

import "fmt"

var pl = fmt.Println

func min_max(arr []float32, k float32) {

	var count float32 = 0
	var total float32 = 0
	var min float32 = 0
	var max float32 = 0
	j := 0

	for i := 0; i < len(arr); i++ {
		total = total + arr[i]
		// pl(total, " <-total")

		if count == k {
			avg := total / k
			// pl(avg)
			if avg > max {
				if min > max {
					min = max
				}
				max = avg
			} else if avg < max {
				if avg < min || min == 0 {
					min = avg
				}
			}
			i = j - 1
			j++

			// pl(i, " <-i")

			count = 0
			total = 0

		}
		count++
	}
	pl("Max:", max, "\nMin:", min)
	pl("Difference of max-min: ", max-min)
}

func main() {
	arr := []float32{3, 8, 9, 15}
	var k float32 = 2

	min_max(arr, k)

}
