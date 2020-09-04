package p0387_字符串中的第一个唯一字符

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, c := range s {
		m[c-'a']++
	}
	for i, c := range s {
		if m[c-'a'] == 1 {
			return i
		}
	}
	return -1
}
