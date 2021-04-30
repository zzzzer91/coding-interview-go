// https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof/

package main

func main() {

}

func exist(board [][]byte, word string) bool {
	rowLen, colLen := len(board), len(board[0])
	st := make([]bool, rowLen*colLen)
	var dfs func(x, y, u int) bool
	dfs = func(x, y, u int) bool {
		if u == len(word) {
			return true
		}
		stIdx := x*colLen + y
		if x < 0 || x >= rowLen || y < 0 || y >= colLen || word[u] != board[x][y] || st[stIdx] {
			return false
		}
		st[stIdx] = true
		defer func() { st[stIdx] = false }()
		u += 1
		return dfs(x-1, y, u) || dfs(x, y+1, u) || dfs(x+1, y, u) || dfs(x, y-1, u)
	}
	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 省略一个额外数组
func exist2(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	m, n := len(board), len(board[0])

	var dfs func(u, x, y int) bool
	dfs = func(u, x, y int) bool {
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[u] {
			return false
		}
		if u == len(word)-1 {
			return true
		}
		u += 1
		temp := board[x][y]
		board[x][y] = '#'
		res := dfs(u, x+1, y) || dfs(u, x-1, y) || dfs(u, x, y+1) || dfs(u, x, y-1)
		board[x][y] = temp
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(0, i, j) {
				return true
			}
		}
	}
	return false
}
