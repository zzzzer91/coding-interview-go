//给出A，B两国想要的土地，输出只有A国想要的土地数，只有B国想要的土地数，两个国家都想要的土地数。
//思路：简单的求交集大小，甚至不用给出交集，先遍历A，用set存下来，再遍历B，遇到set中有的就cnt++，最后输出A-cnt, B-cnt, cnt

package main

import "fmt"

const N = 1e5 + 10

var n, p, q int

var a, b [N]int

func main() {
	fmt.Scanf("%d %d %d", &n, &p, &q)
	m := make(map[int]struct{})
	count := 0
	for i := 0; i < p; i++ {
		fmt.Scanf("%d", &a[i])
		m[a[i]] = struct{}{}
	}
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &b[i])
		if _, ok := m[b[i]]; ok {
			count++
		}
	}
	fmt.Println(p-count, q-count, count)
}
