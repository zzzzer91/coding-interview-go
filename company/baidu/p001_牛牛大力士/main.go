// n 个大力士，m 场比试，求 p 的排名
// 第二行之后，会输入战绩，a 战胜了 b

package main

import "fmt"

const N = 110

var n, m, p int
var g [N][N]bool
var visited1, visited2 [N]bool

/*
4 5 4
1 2
2 3
4 3
1 4
2 4
*/
func main() {
	fmt.Scanf("%d%d%d", &n, &m, &p)
	var a, b int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d%d", &a, &b)
		g[a][b] = true
	}
	start, end := dfs1(p)+1, n-dfs2(p)
	for i := start; i <= end; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func dfs1(x int) int {
	visited1[x] = true
	res := 0
	for i := 1; i <= n; i++ {
		if !visited1[i] && g[i][x] {
			res += dfs1(i) + 1
		}
	}
	return res
}

func dfs2(x int) int {
	visited2[x] = true
	res := 0
	for i := 1; i <= n; i++ {
		if !visited2[i] && g[x][i] {
			res += dfs2(i) + 1
		}
	}
	return res
}
