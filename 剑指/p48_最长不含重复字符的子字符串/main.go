// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

package main

func main() {

}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	res := 0
	var m [128]int
	i := 0
	for j := range s {
		lastIdx := m[s[j]]
		if lastIdx > i {
			i = lastIdx
		}
		res = max(res, j-i+1)
		m[s[j]] = j + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
