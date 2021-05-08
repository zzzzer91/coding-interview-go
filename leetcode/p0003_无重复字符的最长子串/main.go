// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

package main

func main() {

}

// 只能是 ascii 码
// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	res := 0
	// 存储上一次出现位置，为 i+1
	var m [128]int
	for l, r := 0, 0; r < len(s); r++ {
		l = max(l, m[s[r]])
		res = max(res, r-l+1)
		// 以 原索引+1 为索引，因为数组默认初始化 0
		m[s[r]] = r + 1
	}
	return res
}

// 任意字符
func lengthOfLongestSubstring2(s string) int {
	ret := 0
	showedIdx := make(map[rune]int)
	cursor := 0
	for i, c := range s {
		cursor = max(cursor, showedIdx[c])
		ret = max(ret, i-cursor+1)
		showedIdx[c] = i + 1
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
