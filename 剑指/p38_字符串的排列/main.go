// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/

package main

func main() {

}

func permutation(s string) []string {
	bytes := []byte(s)
	var res []string
	var comb []byte
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(bytes) {
			res = append(res, string(comb))
			return
		}
		var visited [26]bool // 集合去重
		for i := u; i < len(bytes); i++ {
			if visited[bytes[i]-'a'] {
				continue
			}
			visited[bytes[i]-'a'] = true
			comb = append(comb, bytes[i])
			bytes[i], bytes[u] = bytes[u], bytes[i] // 省一个数组去表示已经用过的元素，很巧秒
			dfs(u + 1)
			comb = comb[:len(comb)-1]
			bytes[i], bytes[u] = bytes[u], bytes[i]
		}
	}
	dfs(0)
	return res
}
