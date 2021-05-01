// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
// 注意字符串中可能有重复字符

package main

func main() {

}

// 根据字符串排列的特点，考虑深度优先搜索所有排列方案。
// 即通过字符交换，先固定第 1 位字符（ n 种情况）、
// 再固定第 2 位字符（ n-1 种情况）、... 、最后固定第 n 位字符（ 1 种情况）。
func permutation(s string) []string {
	bytes := []byte(s)
	var res []string
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(bytes) {
			res = append(res, string(bytes))
			return
		}
		var visited [26]bool // 用于去除字符串中的重复字符
		for i := u; i < len(bytes); i++ {
			if visited[bytes[i]-'a'] {
				continue
			}
			visited[bytes[i]-'a'] = true
			bytes[i], bytes[u] = bytes[u], bytes[i]
			dfs(u + 1)
			bytes[i], bytes[u] = bytes[u], bytes[i]
		}
	}
	dfs(0)
	return res
}
