// https://leetcode-cn.com/problems/island-perimeter/

package main

func main() {

}

func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
			return 1
		}
		if grid[x][y] == 2 {
			return 0
		}
		grid[x][y] = 2
		return dfs(x-1, y) + dfs(x, y+1) + dfs(x+1, y) + dfs(x, y-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return dfs(i, j)
			}
		}
	}
	return -1
}
