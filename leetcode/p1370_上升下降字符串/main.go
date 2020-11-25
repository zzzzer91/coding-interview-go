// https://leetcode-cn.com/problems/increasing-decreasing-string/

package main

func main() {

}

func sortString(s string) string {
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}
	var res []byte
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if m[i] > 0 {
				res = append(res, byte('a'+i))
				m[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if m[i] > 0 {
				res = append(res, byte('a'+i))
				m[i]--
			}
		}
	}
	return string(res)
}
