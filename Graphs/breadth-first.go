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

func breadth(graph [][]int, src int) {
	stack := graph[src]

	pl(src)

	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		pl(curr)
		for i := 0; i < len(graph[curr]); i++ {
			stack = append(stack, graph[curr][i])
		}

	}

}

func main() {
	breadth(graph, 0)

}
