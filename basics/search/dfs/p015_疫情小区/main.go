// 求 1 聚集的个数
/*
3 3
1 1 0
1 0 1
0 0 0

3
5
1 1 0 0 0
1 0 1 1 0
0 1 0 0 1
*/

package main

import "fmt"

const N = 110

var m, n int
var g [N][N]int
var dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

func main() {
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &g[i][j])
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func dfs(x, y int) bool {
	if g[x][y] == 0 {
		return false
	}
	g[x][y] = 0
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < m && b >= 0 && b < n {
			dfs(a, b)
		}
	}
	return true
}
