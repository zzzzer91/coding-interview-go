// https://leetcode-cn.com/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/

package main

func main() {

}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	pos := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) != 0 && popped[pos] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			pos++
		}
	}
	return len(stack) == 0
}
