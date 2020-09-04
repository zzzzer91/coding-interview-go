package p0058_最后一个单词的长度

import "strings"

func lengthOfLastWord(s string) int {
	length := 0
	foundWord := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if foundWord {
				break
			}
			continue
		}
		length++
		foundWord = true
	}
	return length
}

func lengthOfLastWord2(s string) int {
	s = strings.Trim(s, " ")
	index := strings.LastIndexByte(s, ' ')
	return len(s) - index - 1
}
