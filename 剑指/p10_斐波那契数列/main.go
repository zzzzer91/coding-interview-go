// https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/

package main

func main() {

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}
