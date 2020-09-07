// 黑红棋,每次可以互换相邻的棋子,求使得每种颜色棋子递增排序的最小交换次数。

package main

import "fmt"

const N = 3010

var a [N * 2]int

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var s string
	fmt.Scanf("%s", &s)
	for i := 0; i < 2*n; i++ {
		fmt.Scanf("%d", a[i])
	}
	fmt.Println(n)
}
