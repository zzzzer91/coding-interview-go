// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

package main

func main() {

}

func movingCount(m int, n int, k int) int {
	st := make([]bool, m*n)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || st[x*n+y] || !canEnter(x, y, k) {
			return 0
		}
		st[x*n+y] = true
		return dfs(x-1, y) + dfs(x, y+1) + dfs(x+1, y) + dfs(x, y-1) + 1
	}
	return dfs(0, 0)
}

func canEnter(x, y, k int) bool {
	res := 0
	for x > 0 {
		res += x % 10
		x /= 10
	}
	for y > 0 {
		res += y % 10
		y /= 10
	}
	return res <= k
}
