package main

import "fmt"

var pl = fmt.Println

var graph = [][]int{
	0: {2, 1},
	1: {3},
	2: {4},
	3: {5},
	4: {},
	5: {},
}

func depth(graph [][]int, src int) {
	store := graph[src]

	pl(src)

	for len(store) > 0 {
		curr := store[len(store)-1]
		store = store[:len(store)-1]

		pl(curr)

		for i := 0; i < len(graph[curr]); i++ {
			store = append(store, graph[curr][i])
		}
	}

}

func main() {
	depth(graph, 0)
}
