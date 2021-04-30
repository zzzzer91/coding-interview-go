// https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/

package main

func main() {

}

func movingCount(m int, n int, k int) int {
	st := make([]bool, m*n)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || st[x*n+y] || !isCanEnter(x, y, k) {
			return 0
		}
		st[x*n+y] = true
		return dfs(x-1, y) + dfs(x, y+1) + dfs(x+1, y) + dfs(x, y-1) + 1
	}
	return dfs(0, 0)
}

// 向上和向左的移动其实可以省略
func movingCount2(m int, n int, k int) int {
	st := make([]bool, m*n)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || st[x*n+y] {
			return 0
		}
		st[x*n+y] = true
		if !isCanEnter(x, y, k) {
			return 0
		}
		return dfs(x+1, y) + dfs(x, y+1) + 1
	}
	return dfs(0, 0)
}

func isCanEnter(x, y, k int) bool {
	return compute(x)+compute(y) <= k
}

func compute(n int) int {
	res := 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
