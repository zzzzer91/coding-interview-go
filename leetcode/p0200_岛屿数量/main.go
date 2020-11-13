// https://leetcode-cn.com/problems/number-of-islands/

package main

func main() {

}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			return false
		}
		grid[x][y] = '0'
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
		return true
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j) {
				res++
			}
		}
	}
	return res
}
