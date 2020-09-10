// 点名，不在队列或者在队列的被点到去队首
// 思路：倒着数即可，某个数出现的最后一次是他的实际排序，前面的点名完全被覆盖掉，没卵用。
// 用一个unordered_set存一下出现过的数，出现过就直接跳过，没出现过就print出来。
package main

import "fmt"

const N = 1e9 + 10

var a [N]int

func main() {
	var m int
	var no int
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", no)
		fmt.Println(no)
		//if a[no] == no {
		//	a[no] = 0
		//	fmt.Println(no)
		//} else if a[no] == 0 {
		//	a[no] = no
		//}
	}
}
