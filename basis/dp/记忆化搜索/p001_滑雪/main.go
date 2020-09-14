// https://www.acwing.com/problem/content/description/903/
// 和 https://leetcode-cn.com/problems/word-search/ 有些异曲同工之妙

package main

import (
	"fmt"
)

const N = 310

var r, c int
var g [N][N]int
var dp [N][N]int
var dx, dy = [4]int{-1, 0, 1, 0}, [4]int{0, 1, 0, -1}

func main() {
	fmt.Scanf("%d %d", &r, &c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scanf("%d", &g[i][j])
			dp[i][j] = -1
		}
	}
	res := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			res = max(res, dfs(i, j))
		}
	}
	fmt.Println(res)
}

func dfs(x, y int) int {
	if dp[x][y] != -1 { // 记忆化搜索关键一步，代表之前已经搜过，不需要再搜了
		return dp[x][y]
	}
	dp[x][y] = 1 // 刚开始不动也算一种方案
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < r && b >= 0 && b < c && g[a][b] < g[x][y] {
			dp[x][y] = max(dp[x][y], dfs(a, b)+1)
		}
	}
	return dp[x][y]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
