package main

import (
	"fmt"
	"sort"
)

var pl = fmt.Println

var values = []struct {
	Price  float32
	Weight float32
}{
	{60, 10},
	{100, 20},
	{120, 30},
}

func knapsack(W float32) float32 {

	var finalValue float32 = 0.0

	sort.SliceStable(values, func(i, j int) bool {
		v1 := values[i].Price / values[i].Weight
		v2 := values[j].Price / values[j].Weight
		return v1 > v2
	})

	for _, value := range values {

		if value.Weight <= W {
			W -= value.Weight
			finalValue += value.Price
		} else {
			finalValue += value.Price * (W / value.Weight)
			break
		}

	}

	return finalValue

}

func main() {

	temp := knapsack(50)

	pl(temp)

}
