// https://leetcode-cn.com/problems/letter-case-permutation/

package main

func main() {

}

func letterCasePermutation(S string) []string {
	var res []string
	s := []byte(S)
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(s) {
			res = append(res, string(s))
			return
		}
		dfs(u + 1)
		if s[u] >= 'a' && s[u] <= 'z' {
			s[u] -= 32
			dfs(u + 1)
		} else if s[u] >= 'A' && s[u] <= 'Z' {
			s[u] += 32
			dfs(u + 1)
		}
	}
	dfs(0)
	return res
}
