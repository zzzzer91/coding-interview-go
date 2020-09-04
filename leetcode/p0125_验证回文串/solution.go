package p0125_验证回文串

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		cLeft, cRight := s[left], s[right]
		if !isAlphanumeric(cLeft) {
			left++
			continue
		}
		if !isAlphanumeric(cRight) {
			right--
			continue
		}
		if toLowerCase(cLeft) != toLowerCase(cRight) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	if (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
		return true
	}
	return false
}

func toLowerCase(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c | 0x20
	}
	return c
}
