// https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/

package main

func main() {

}

func reverseLeftWords1(s string, n int) string {
	tmp := []byte(s)
	reverse(tmp[:n])
	reverse(tmp[n:])
	reverse(tmp)
	return string(tmp)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

func reverseLeftWords2(s string, n int) string {
	return (s + s)[n : len(s)+n]
}

func reverseLeftWords3(s string, n int) string {
	return s[n:] + s[:n]
}
