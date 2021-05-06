// https://leetcode-cn.com/problems/permutation-in-string/

package main

func main() {

}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// 统计字符频率，频率一样，即是排列，很棒的思路
	needChars := make([]int, 26)
	window := make([]int, 26)
	for i, _ := range s1 {
		needChars[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	for i := len(s1); i < len(s2); i++ {
		window[s2[i-len(s1)]-'a']--
		window[s2[i]-'a']++
		if isSameArray(needChars, window) {
			return true
		}
	}
	return isSameArray(needChars, window)
}

func isSameArray(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
