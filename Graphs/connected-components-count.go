package main

import "fmt"

var pl = fmt.Println

var graph = map[int][]int{
	0: {8, 1, 5},
	1: {0},
	5: {0, 8},
	8: {0, 5},
	2: {3, 4},
	3: {2, 4},
	4: {3, 2},
}

func explore(memo map[int]int, graph map[int][]int, src int) bool {

	if memo[src] != 0 {
		return false
	}

	memo[src] = src

	for i := 0; i < len(graph[src]); i++ {
		explore(memo, graph, graph[src][i])
	}

	return true

}

func connected_count(graph map[int][]int) int {
	count := 0
	memo := map[int]int{}

	for k := range graph {
		if explore(memo, graph, k) == true {
			count++
		}
	}

	return count
}

func main() {

	count := connected_count(graph)
	pl(count)

	// temp := map[int]int{}

	// pl(temp[32])

}
