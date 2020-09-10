package main

import "unicode/utf8"

func main() {

}

// 输入字符串只有小写字母
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i, _ := range s {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}
	for _, i := range m {
		if i != 0 {
			return false
		}
	}
	return true
}

// 假设输入的是utf8
func isAnagram2(s string, t string) bool {
	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t) {
		return false
	}
	tTemp := []rune(t)
	m := make(map[rune]int)
	for i, c := range s {
		m[c]++
		m[tTemp[i]]--
	}
	for _, i := range m {
		if i != 0 {
			return false
		}
	}
	return true
}
