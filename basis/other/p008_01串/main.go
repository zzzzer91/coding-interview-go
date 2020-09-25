// 不含 110 的最长子串

package main

import "fmt"

func main() {
	s := "1101010110010110"
	fmt.Println(f(s))
}

func f(s string) int {
	res := 2
	count := 2
	for i := 2; i < len(s); i++ {
		if s[i-2] == '1' && s[i-1] == '1' && s[i] == '0' {
			res = max(res, count)
			count = 2
			continue
		}
		count++
	}
	res = max(res, count)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
