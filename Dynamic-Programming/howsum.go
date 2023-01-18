package main

import "fmt"

var pl = fmt.Println

func howsum(numbers []int, target int) []int {
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	for i := 0; i < len(numbers); i++ {
		remainder := target - numbers[i]
		result := howsum(numbers, remainder)
		if result != nil {
			temp := []int{numbers[i]}
			temp = append(temp, result...)
			return temp
		}
	}

	return nil
}

func main() {

	pl(howsum([]int{2, 3, 6, 4, 7}, 7))

}
