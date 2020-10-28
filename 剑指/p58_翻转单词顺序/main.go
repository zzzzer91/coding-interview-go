// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	s = " " + strings.TrimSpace(s)
	var res []byte
	pos := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			res = append(res, s[i+1:pos+1]...)
			res = append(res, ' ')
			for i >= 0 && s[i] == ' ' {
				i--
			}
			pos = i
		}
	}
	return string(res[:len(res)-1])
}
