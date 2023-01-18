package main

import (
	"fmt"
)

var pl = fmt.Println

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func knapsack(wt []int, val []int, Wt int, N int) int {
	if Wt == 0 || N == 0 {
		return 0
	}
	if wt[N-1] > Wt {
		return knapsack(wt, val, Wt, N-1)
	} else {
		return max(val[N-1]+knapsack(wt, val, Wt-wt[N-1], N-1), knapsack(wt, val, Wt, N-1))
	}

}

func main() {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}

	Wt := 50

	total := knapsack(wt, val, Wt, len(val))

	pl(total)
}
