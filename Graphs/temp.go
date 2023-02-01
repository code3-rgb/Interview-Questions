package main

import "fmt"

var pl = fmt.Println

var store = [][]string{
	{"i", "j"},
	{"k", "i"},
	{"m", "k"},
	{"k", "l"},
	{"o", "n"},
}

func convert(store [][]string) map[string][]string {
	tmp := map[string][]string{}

	for i := 0; i < len(store); i++ {
		a := store[i][0]
		b := store[i][1]

		if tmp[a] == nil {
			tmp[a] = []string{}
		}
		if tmp[b] == nil {
			tmp[b] = []string{}
		}

		tmp[a] = append(tmp[a], b)
		tmp[b] = append(tmp[b], a)

	}

	return tmp

}

func has_path(dest, src string) bool {
	graph := convert(store)

	stack := graph[src]

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if curr == dest {
			return true
		}

		for i := 0; i < len(graph[curr]); i++ {
			stack = append(stack, graph[curr][i])
		}

	}

	return false

}

func main() {
	pl(has_path("k", "n"))
}
