package main

import "fmt"

var pl = fmt.Println

func fibo1(num int) {
	fib := []int{0, 1}

	var temp int

	for i := 1; i < num; i++ {
		temp = fib[i-1] + fib[i]
		fib = append(fib, temp)
	}

	pl("Fibonacci of", num, "is", fib[num])
}

func fibo2(num int) int {

	if num <= 2 {
		return 1
	}
	return fibo2(num-1) + fibo2(num-2)
}

func main() {
	fibo1(14)
	pl(fibo2(14))

}
