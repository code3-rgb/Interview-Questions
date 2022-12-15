package main

import "fmt"

var pl = fmt.Sprintf

func subarry(arr []int, k int) {

	count := 0
	tst := 0
	lst := 0

	for i := 0; i < len(arr); i++ {
		if count < k {
			count += arr[i]
		}
		if count > k {
			i = tst
			tst += 1

			count = 0
		}
		if count == k {
			lst = i
			break
		}
		if tst >= len(arr)-1 {
			pl("Not found")
			return
		}

	}

	if count < k || count > k {
		pl("Not found")
		return
	}
	result := pl("Your no:%d is from index:%d  to index:%d", count, tst, lst)

	fmt.Println(result)

}

func main() {
	arr := []int{1, 2, 3, 4, 3, 6, 5, 8}

	subarry(arr, 13)

}
