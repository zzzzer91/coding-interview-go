// 思路：观察到异或的两个性质：1. 可交换 2. 任何数异或自身为零。
// 观察到mod的一个性质：循环。因此我们可以竖着求，而不是横着求。也就是我们遍历j去求，而不是i。

package main

import "fmt"

const N = 1e5 + 10

var n int
var a [N]int

func main() {
	fmt.Scanf("%d", &n)
	var res int
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &a[i])
		t := a[i]
		for j := 1; j <= n; j++ {
			if i < j {
				t ^= i
			} else {
				t ^= i % j
			}
		}
		res ^= t
	}
	fmt.Println(res)
}
