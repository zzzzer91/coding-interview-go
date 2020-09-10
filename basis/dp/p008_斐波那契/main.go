package main

import "fmt"

func fibonacci(n int) int {
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(fibonacci(4))
}
