// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

package main

func main() {

}

func permutation(s string) []string {
	bytes := []byte(s)
	var res []string
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(bytes) {
			res = append(res, string(bytes))
			return
		}
		var visited [26]bool
		for i := u; i < len(bytes); i++ {
			if visited[bytes[i]-'a'] {
				continue
			}
			visited[bytes[i]-'a'] = true
			bytes[i], bytes[u] = bytes[u], bytes[i] // 非常巧妙
			dfs(u + 1)
			bytes[i], bytes[u] = bytes[u], bytes[i]
		}
	}
	dfs(0)
	return res
}
