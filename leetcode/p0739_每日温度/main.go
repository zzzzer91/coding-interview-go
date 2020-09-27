// https://leetcode-cn.com/problems/daily-temperatures/

package main

func main() {

}

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	var stack []int
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
