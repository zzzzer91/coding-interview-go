// https://leetcode-cn.com/problems/remove-k-digits/

package main

import "strings"

func main() {

}

func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		res = "0"
	}
	return res
}
