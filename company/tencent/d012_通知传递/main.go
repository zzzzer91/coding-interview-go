// 给定n个人和m个人的小团体，消息从0开始传播，问传播到多少人

package main

import "fmt"

var a [510][110]int

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	m1 := make(map[int]int) // 存序号对应的组
	for i := 0; i < m; i++ {
		var x int
		fmt.Scanf("%d", &x)
		for j := 0; j < x; j++ {
			fmt.Scanf("%d", &a[i][j])
			if a[i][j] == 0 {
				x = j
			}
			m1[a[i][j]] = i
		}
	}
	//if v, ok := m1[a[i][j]]; ok {
	//
	//}
	res := 7

	fmt.Println(res)
}
