// https://leetcode-cn.com/problems/n-queens/

package main

func main() {

}

func solveNQueens(n int) [][]string {
	var res [][]string
	colSt, digSt, uDigSt := make([]bool, n), make([]bool, n*2), make([]bool, n*2)
	comb := make([]byte, n*n)
	for i := range comb {
		comb[i] = '.'
	}
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			temp := make([]string, n)
			for i := range temp {
				temp[i] = string(comb[i*n : (i+1)*n])
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if colSt[i] || digSt[i+u] || uDigSt[i-u+n] {
				continue
			}
			colSt[i], digSt[i+u], uDigSt[i-u+n] = true, true, true
			comb[u*n+i] = 'Q'
			dfs(u + 1)
			comb[u*n+i] = '.'
			colSt[i], digSt[i+u], uDigSt[i-u+n] = false, false, false
		}
	}
	dfs(0)
	return res
}

// 状态压缩，优化空间
func solveNQueens2(n int) [][]string {
	var res [][]string
	colSt, digSt, uDigSt := 0, 0, 0 // 状态压缩
	comb := make([]byte, n*n)
	for i := range comb {
		comb[i] = '.'
	}
	var dfs func(u int)
	dfs = func(u int) {
		if u == n {
			temp := make([]string, n)
			for i := range temp {
				temp[i] = string(comb[i*n : (i+1)*n])
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if colSt&(1<<i) != 0 || digSt&(1<<(i+u)) != 0 || uDigSt&(1<<(i-u+n)) != 0 {
				continue
			}
			colSt ^= 1 << i
			digSt ^= 1 << (i + u)
			uDigSt ^= 1 << (i - u + n)
			comb[u*n+i] = 'Q'
			dfs(u + 1)
			comb[u*n+i] = '.'
			colSt ^= 1 << i
			digSt ^= 1 << (i + u)
			uDigSt ^= 1 << (i - u + n)
		}
	}
	dfs(0)
	return res
}
