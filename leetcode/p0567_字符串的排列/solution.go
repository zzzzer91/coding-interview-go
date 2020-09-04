package p0567_字符串的排列

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	needChars := make([]int, 26)
	window := make([]int, 26)
	for i, _ := range s1 {
		needChars[s1[i]-'a']++
		window[s2[i]-'a']++
	}
	if isSameArray(needChars, window) {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		window[s2[i-len(s1)]-'a']--
		window[s2[i]-'a']++
		if isSameArray(needChars, window) {
			return true
		}
	}
	return false
}

func isSameArray(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
