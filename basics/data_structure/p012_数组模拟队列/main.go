// 用数组模拟队列，比较简单，做题时常用

package main

import "fmt"

func main() {
	const N = 2
	q := make([]int, N)
	ff, rr := 0, -1 // ff 头部指针，rr 尾部指针

	// 从尾部压入队列
	rr++
	q[rr] = 11
	rr++
	q[rr] = 12

	// 队列判满
	if rr == N-1 {
		fmt.Println("队列为满")
	}

	// 从头部弹出队列
	ff++

	// 从尾部弹出队列
	rr--

	// 队列判空
	if ff > rr {
		fmt.Println("队列为空")
	}
}
