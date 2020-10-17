// https://leetcode-cn.com/problems/n-queens-ii/

package main

func main() {

}

func totalNQueens(n int) int {
	res := 0
	colSt, digSt, uDigSt := 0, 0, 0 // 状态压缩
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			res++
			return
		}
		for i := 0; i < n; i++ {
			if colSt&(1<<i) != 0 || digSt&(1<<(i+u)) != 0 || uDigSt&(1<<(i-u+n)) != 0 {
				continue
			}
			colSt ^= 1 << i
			digSt ^= 1 << (i + u)
			uDigSt ^= 1 << (i - u + n)
			dfs(u + 1)
			colSt ^= 1 << i
			digSt ^= 1 << (i + u)
			uDigSt ^= 1 << (i - u + n)
		}
	}
	dfs(0)
	return res
}

// 大佬的代码，没看懂
func totalNQueens2(n int) int {
	count := 0
	var dfs func(row, col, pie, na int)
	dfs = func(row, col, pie, na int) {
		if row >= n {
			count++
			return
		}
		bits := (^(col | pie | na)) & ((1 << n) - 1)
		for bits > 0 {
			p := bits & -bits
			bits &= bits - 1
			dfs(row+1, col|p, (pie|p)<<1, (na|p)>>1)
		}
	}
	dfs(0, 0, 0, 0)
	return count
}
