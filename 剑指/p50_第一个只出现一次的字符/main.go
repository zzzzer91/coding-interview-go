// https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/

package main

func main() {

}

func firstUniqChar(s string) byte {
	var m [128]int
	for i := range s {
		m[s[i]]++
	}
	for i := range s {
		if m[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
