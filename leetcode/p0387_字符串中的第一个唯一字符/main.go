// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

package main

func main() {

}

func firstUniqChar(s string) int {
	var m [26]int
	for i := range s {
		m[s[i]-'a']++
	}
	for i := range s {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
