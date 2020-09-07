// 输出前k多和前k少的字符串

package main

import "fmt"

const N = 1e5 + 10

var n, k int

func main() {
	fmt.Scanf("%d %d", &n, &k)
	m := make(map[string]int)
	var s string
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		m[s]++
	}
}
