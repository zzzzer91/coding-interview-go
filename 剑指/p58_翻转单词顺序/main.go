// https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/

package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	var res []byte
	s = " " + strings.TrimSpace(s)
	for i, j := len(s)-1, len(s)-1; i >= 0; {
		if s[i] != ' ' {
			i--
		} else {
			res = append(res, s[i:j+1]...)
			for i >= 0 && s[i] == ' ' {
				i--
			}
			j = i
		}
	}
	return string(res[1:])
}
