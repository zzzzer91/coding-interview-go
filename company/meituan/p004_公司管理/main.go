//. 出一串子树的节点个数（包含自身），求是否能构成合法树，要求每个非叶节点至少有两个子节点。
// 思路：n=24明显提示我们是dfs+剪枝，先把输入从大到小排序。
// 首先预检查：1. arr[0]必须等于n 2. arr[i]不能等于2

package main

import "fmt"

const N = 30

var n int
var a [N]int

func main() {
	for {
		c, _ := fmt.Scanf("%d", &n)
		if c <= 0 {
			break
		}
		for i := n - 1; i >= 0; i-- {
			fmt.Scanf("%d", &a[i])
		}
		flag := false
		for i := n/2 - 1; i >= 0; i-- {
			l := i*2 + 1
			r := l + 1
			if l < n && r < n {
				if a[i] != n-i || a[i] != a[l]+a[r]+1 {
					fmt.Println("NO")
					flag = true
					break
				}
			} else if l < n {
				fmt.Println("NO")
				flag = true
				break
			} else {
				if a[i] != 1 {
					fmt.Println("NO")
					flag = true
					break
				}
			}
		}
		if !flag {
			fmt.Println("YES")
		}
	}
}
