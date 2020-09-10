package main

func main() {

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	firstStr := strs[0]
	remainingStrs := strs[1:]

	minStrLen := len(firstStr)
	for _, s := range remainingStrs {
		if len(s) < minStrLen {
			minStrLen = len(s)
		}
	}

	for i := 0; i < minStrLen; i++ {
		c := firstStr[i]
		for _, s := range remainingStrs {
			if c != s[i] {
				return firstStr[:i]
			}
		}
	}
	return firstStr[:minStrLen]
}
