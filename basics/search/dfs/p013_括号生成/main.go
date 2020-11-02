// https://leetcode-cn.com/problems/generate-parentheses/

package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var res []string
	var comb []byte
	var dfs func(lc, rc int)
	dfs = func(lc, rc int) {
		if len(comb) == n*2 {
			res = append(res, string(comb))
			return
		}
		if lc < n {
			comb = append(comb, '(')
			dfs(lc+1, rc)
			comb = comb[:len(comb)-1]
		}
		if rc < lc {
			comb = append(comb, ')')
			dfs(lc, rc+1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0, 0)
	return res
}
