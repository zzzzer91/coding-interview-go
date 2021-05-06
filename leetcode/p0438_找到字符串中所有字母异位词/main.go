// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 相似：https://leetcode-cn.com/problems/permutation-in-string/

package main

func main() {

}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	var res []int
	charCounts := make([]int, 26)
	windows := make([]int, 26)
	for i := range p {
		charCounts[p[i]-'a']++
		windows[s[i]-'a']++
	}
	if isSameArray(charCounts, windows) {
		res = append(res, 0)
	}
	for i := len(p); i < len(s); i++ {
		windows[s[i-len(p)]-'a']--
		windows[s[i]-'a']++
		if isSameArray(charCounts, windows) {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func isSameArray(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
