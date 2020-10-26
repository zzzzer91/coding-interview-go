// https://www.acwing.com/problem/content/description/832/

package main

import (
	"fmt"
)

const N = 100010

var stack [N]int // 整个栈必然是单调递增的，因为入栈时，会把所有比当前元素大的从栈顶弹出
var top = -1

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var v int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &v)
		for top != -1 && stack[top] >= v { // 找到栈顶第一个比当前数小的
			top--
		}
		if top == -1 {
			fmt.Printf("%d ", -1)
		} else {
			fmt.Printf("%d ", stack[top])
		}
		top++
		stack[top] = v
	}
	fmt.Println()
}
