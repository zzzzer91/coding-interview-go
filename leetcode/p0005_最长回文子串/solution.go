package main

func main() {

}

// O(n^2) 时间复杂度，使用马拉车算法，可以达到 O(n)，但是不通用
func longestPalindrome(s string) string {
	var res string
	var subS string
	for i := 0; i < len(s); i++ {
		// 奇数子串中心
		subS = findPalindrome(s, i, i)
		res = maxS(subS, res)
		// 偶数子串中心
		subS = findPalindrome(s, i, i+1)
		res = maxS(subS, res)
	}
	return res
}

func maxS(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func findPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
