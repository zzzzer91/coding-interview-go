// https://leetcode-cn.com/problems/daily-temperatures/

package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var stack []int // 存放下标
	for i, t := range T {
		for len(stack) > 0 && t > T[stack[len(stack)-1]] {
			stackTopV := stack[len(stack)-1]
			res[stackTopV] = i - stackTopV
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
