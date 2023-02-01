package main

import "fmt"

var pl = fmt.Println

var graph = map[string][]string{
	"f": {"g", "i"},
	"g": {"h"},
	"h": {},
	"i": {"g", "k"},
	"j": {"i"},
	"k": {},
}

// has path
func has_path(graph map[string][]string, dst, src string) bool {
	stack := graph[src]

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if dst == curr {
			return true
		}

		for i := 0; i < len(graph[curr]); i++ {

			stack = append(stack, graph[curr][i])
		}

	}

	return false

}

func main() {

	pl(has_path(graph, "h", "g"))

}
