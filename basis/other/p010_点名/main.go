// n 个人，循环点名，从 1 开始点，点到 m 的人出列，然后下一个继续从 1 开始点
// 问：剩下的人的初始编号

package main

import (
	"fmt"
)

var n, m = 5, 3

func main() {
	fmt.Println(f())
}

func f() int {
	a := make([]bool, n)
	count := n
	i := 0
	j := 0
	for count > 1 {
		if j == m-1 {
			a[i] = true
			count--
			j = 0
		} else {
			j++
		}
		i = (i + 1) % n
		for a[i] {
			i = (i + 1) % n
		}
	}
	for i := 0; i < n; i++ {
		if !a[i] {
			return i + 1
		}
	}
	return -1
}

func f2() int {
	a := make([]bool, n+1)
	count := n
	i := 1
	j := 1
	for count > 1 {
		if j == m {
			a[i] = true
			count--
			j = 1
		} else {
			j++
		}
		i = (i + 1) % n
		if i == 0 {
			i = n
		}
		for a[i] {
			i = (i + 1) % n
			if i == 0 {
				i = n
			}
		}
	}
	for i := 1; i <= n; i++ {
		if !a[i] {
			return i + 1
		}
	}
	return -1
}
