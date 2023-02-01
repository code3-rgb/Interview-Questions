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

func hasPath(dest, src string, graph map[string][]string, memo map[string]string) bool {

	if src == dest {
		return true
	}
	if memo[src] != "" {
		return false
	}
	memo[src] = src

	var status bool

	for i := 0; i < len(graph[src]); i++ {
		status = hasPath(dest, graph[src][i], graph, memo)
		if status {
			return true
		}
	}
	return false

}

func main() {

	graphs := convert(store)

	pl(graphs)

	memo := map[string]string{}
	status := hasPath("n", "k", graphs, memo)

	pl(status)

}
