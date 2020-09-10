// 给一串偶数个字符，只有大小写字母，求修改多少个字母可以让大小写数量相同。
// 思路：求出大小写个数，相减除以二即可。

package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)
	a, b := 0, 0
	for i := range s {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			a++
		} else if c >= 'A' && c <= 'Z' {
			b++
		}
	}
	fmt.Println(abs(a-b) / 2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
