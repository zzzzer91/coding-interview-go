// https://leetcode-cn.com/problems/valid-palindrome/

package main

func main() {

}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		lc, rc := s[l], s[r]
		if !isAlphanumeric(lc) {
			l++
			continue
		}
		if !isAlphanumeric(rc) {
			r--
			continue
		}
		if toLowerCase(lc) != toLowerCase(rc) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c | 0x20
	}
	return c
}
