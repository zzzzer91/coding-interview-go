// 点名，不在队列或者在队列的被点到去队首
// 思路：倒着数即可，某个数出现的最后一次是他的实际排序，前面的点名完全被覆盖掉，没卵用。
// 用一个unordered_set存一下出现过的数，出现过就直接跳过，没出现过就print出来。
package main

import "fmt"

const N = 1e9 + 10

var a [N]int

func main() {
	var m int
	fmt.Scanf("%d", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &a[i])
	}
	set := make(map[int]struct{}, m)
	for i := m - 1; i >= 0; i-- {
		tmp := a[i]
		if _, ok := set[tmp]; !ok {
			fmt.Println(tmp)
			set[tmp] = struct{}{}
		}

	}
}
