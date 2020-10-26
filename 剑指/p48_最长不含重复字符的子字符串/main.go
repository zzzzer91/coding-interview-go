// https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/

package main

func main() {

}

func lengthOfLongestSubstring(s string) int {
	res := 0
	var m [128]int
	pos := 0
	for i := range s {
		lastIdx := m[s[i]]
		if lastIdx > pos {
			pos = lastIdx
		}
		res = max(res, i-pos+1)
		m[s[i]] = i + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
