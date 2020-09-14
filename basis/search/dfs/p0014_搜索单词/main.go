// https://leetcode-cn.com/problems/word-search/
// 同一元素只能被使用一次

package main

func main() {

}

func exist(board [][]byte, word string) bool {
	dx, dy := []int{-1, 0, 1, 0}, []int{0, -1, 0, 1} // 四个方向
	rowLen := len(board)
	colLen := len(board[0])
	st := make([][]bool, rowLen) // 当前位置是否已被访问过
	for i := range st {
		st[i] = make([]bool, colLen)
	}
	var dfs func(x, y, u int) bool
	dfs = func(x, y, u int) bool {
		if board[x][y] != word[u] {
			return false
		}
		if u == len(word)-1 {
			return true
		}
		st[x][y] = true
		for i := 0; i < 4; i++ {
			a, b := x+dx[i], y+dy[i]
			if a >= 0 && a < rowLen && b >= 0 && b < colLen && !st[a][b] {
				if dfs(a, b, u+1) {
					return true
				}
			}
		}
		st[x][y] = false
		return false
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

// 代码更简洁
func exist2(board [][]byte, word string) bool {
	rowLen := len(board)
	colLen := len(board[0])
	st := make([][]bool, rowLen) // 当前位置是否已被访问过
	for i := range st {
		st[i] = make([]bool, colLen)
	}
	var dfs func(x, y, u int) bool
	dfs = func(x, y, u int) bool {
		if u == len(word) {
			return true
		}
		if x < 0 || x >= rowLen || y < 0 || y >= colLen || st[x][y] || board[x][y] != word[u] {
			return false
		}
		st[x][y] = true
		res := dfs(x-1, y, u+1) || dfs(x+1, y, u+1) || dfs(x, y-1, u+1) || dfs(x, y+1, u+1)
		st[x][y] = false
		return res
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
