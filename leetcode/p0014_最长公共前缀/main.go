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
	firstStr, remainStrs := strs[0], strs[1:]
	for i := 0; i < len(firstStr); i++ { // 迭代第一个字符串
		for j := 0; j < len(remainStrs); j++ { // 迭代剩余字符串在数组的索引
			if i == len(remainStrs[j]) || remainStrs[j][i] != firstStr[i] {
				return firstStr[:i]
			}
		}
	}
	return firstStr
}
