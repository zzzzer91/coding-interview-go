// https://leetcode-cn.com/problems/valid-anagram/

package main

import "unicode/utf8"

func main() {

}

// 输入字符串只有小写字母
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}
	for i := range t {
		if m[t[i]-'a'] == 0 {
			return false
		}
		m[t[i]-'a']--
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
