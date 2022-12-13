package main

import "fmt"

var pl = fmt.Println

func findPeak2() {
	// peak_array := []int{1, 1, 1, 4, 3}
	peak_array := []int{1, 2, 1, 1, 3, 5, 6, 4}

	size := len(peak_array)

	var left = 0
	var right = size - 1

	for left < right {
		var mid = left + (right-left)/2
		if peak_array[mid] < peak_array[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	pl(left)
}

func findPeak() {
	// peak_array := []int{1, 2, 1, 1, 3, 5, 6, 4}
	peak_array := []int{1, 2, 1, 1, 3}
	size := len(peak_array)
	var peak []int

	if size == 1 {
		pl(0)
		return
	}
	if peak_array[0] >= peak_array[1] {
		pl(0)
		return
	}
	if peak_array[size-1] >= peak_array[size-2] {
		pl(size - 1)
		return
	}

	for i := 1; i < size; i++ {
		if peak_array[i-1] <= peak_array[i] && peak_array[i+1] <= peak_array[i] {
			peak = append(peak, i)
		}
	}

	pl(peak)
}

func main() {
	// findPeak()
	findPeak2()
}
