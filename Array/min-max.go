package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func sortAndsolve() {
	var len int = 6

	minmax := []int{3, 2, 1, 56, 10000, 167}

	pl("Unsorted List: ", minmax)

	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if minmax[i] < minmax[j] {
				temp := minmax[i]
				minmax[i] = minmax[j]
				minmax[j] = temp
			}
		}
	}

	pl("Sorted List: ", minmax)

	lowest := minmax[len-1]
	highest := minmax[0]

	pl("highest: ", highest, " lowest: ", lowest)
}

func unsortedVersion() {
	minmax := []int{3, 2, 1, 56, 10000, 167}
	len := len(minmax)

	var small int
	var large int

	if minmax[0] > minmax[1] {
		small = minmax[1]
		large = minmax[0]
	} else {
		small = minmax[0]
		large = minmax[1]
	}

	for i := 2; i < len; i++ {
		if minmax[i] > large {
			large = minmax[i]
		} else if minmax[i] < small {
			small = minmax[i]
		}

	}

	pl(small)
	pl(large)
}

func main() {
	start := time.Now()

	// unsortedVersion()
	time.Sleep(1)
	sortAndsolve()

	timeElapsed := time.Since(start)
	fmt.Printf("The `for` loop took %s", timeElapsed)

}
