// https://leetcode-cn.com/problems/longest-common-prefix/

package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstStr := strs[0]
	remainingStrs := strs[1:]

	minStrLen := len(firstStr)
	for _, s := range remainingStrs {
		if len(s) < minStrLen {
			minStrLen = len(s)
		}
	}

	for i := 0; i < minStrLen; i++ {
		c := firstStr[i]
		for _, s := range remainingStrs {
			if c != s[i] {
				return firstStr[:i]
			}
		}
	}
	return firstStr[:minStrLen]
}

// 更简洁
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ { // 迭代第一个字符串
		for j := 1; j < len(strs); j++ { // 迭代剩余字符串在数组的索引
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
