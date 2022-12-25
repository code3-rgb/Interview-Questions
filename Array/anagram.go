package main

import (
	"fmt"
)

var pl = fmt.Println

func anangram(str string, key string) {
	str_len := len([]rune(str))
	key_len := len([]rune(key))

	var m = make(map[string]string)
	count := 0
	total := 0
	var tst string

	for i := 0; i < key_len; i++ {
		m[string(key[i])] = string(key[i])
	}
	for i := 0; i < str_len; i++ {
		tst = string(str[i])
		if m[tst] != "" {
			count++
		}
		if count >= key_len {
			total++
			count = 0
		}

	}
	pl(total)

}

func main() {
	// str := "gotxtogxotgxdogt"
	str := "racecare"
	key := "race"

	anangram(str, key)

}
