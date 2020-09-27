package main

import "fmt"

// 有 a 个物体，b 个隔板；每个箱子最多 k 个独立隔间，每个隔间最多放 v 个物品
// 问：需要的最小箱子数
var a, b, k, v int

func main() {
	for {
		_, err := fmt.Scanf("%d %d %d %d", &a, &b, &k, &v)
		if err != nil {
			break
		}
		fmt.Println(f())
	}
}

func f() int {
	res := 0
	for a > 0 {
		if b >= k-1 {
			a -= k * v
			b -= k - 1
			res++
		} else if b > 0 {
			a -= (b + 1) * v
			b = 0
			res++
		} else {
			a -= v
			res++
		}
	}
	return res
}
