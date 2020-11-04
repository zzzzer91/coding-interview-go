// https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/

package main

func main() {

}

func reverseWords(s string) string {
	res := []byte(s)
	pos := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			reverse(res, pos, i-1)
			for i < len(s) && res[i] == ' ' {
				i++
			}
			pos = i
		} else if i == len(s)-1 {
			reverse(res, pos, i)
		}
	}
	return string(res)
}

func reverse(s []byte, l, r int) []byte {
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}
