// https://leetcode-cn.com/problems/isomorphic-strings/

package main

func main() {

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := [128]byte{}
	t2s := [128]byte{}
	for i := range s {
		c1, c2 := s[i], t[i]
		if s2t[c1] != t2s[c2] {
			return false
		}
		s2t[c1] = c2
		t2s[c2] = c2
	}
	return true
}
