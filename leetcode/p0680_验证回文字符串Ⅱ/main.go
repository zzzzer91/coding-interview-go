// https://leetcode-cn.com/problems/valid-palindrome-ii/

package main

func main() {

}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isValid(s, l+1, r) || isValid(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isValid(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
