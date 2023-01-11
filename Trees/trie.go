package main

import "fmt"

var pl = fmt.Println

type node struct {
	children [26]*node
	isEnd    bool
}
type trie struct {
	tst *node
}

func Init() *trie {
	trie := &trie{tst: &node{}}
	return trie
}

func (trie *trie) insert(str string) {
	current := trie.tst

	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &node{}
		}
		current = current.children[index]
	}
	current.isEnd = true
}

// func (trie *trie) find(str string) []string {
// 	current := trie.tst
// 	strs := []string{}
// 	temp := ""

// 	for i := 0; i < len(str); i++ {
// 		index := str[i] - 'a'
// 		temp += string(str[i])
// 		if current.children[index] == nil {
// 			return []string{}
// 		}
// 		current = current.children[index]
// 	}
// 	if current.isEnd {
// 		return strs
// 	}
// 	return []string{}

// }

func (trie *trie) search(str string) bool {

	current := trie.tst

	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.isEnd {
		return true
	}
	return false
}

func main() {
	trie := Init()

	trie.insert("john")

	t := trie.search("john")
	pl(t)

	t = trie.search("mercy")
	pl(t)

	// pl(trie.find("john"))

}
