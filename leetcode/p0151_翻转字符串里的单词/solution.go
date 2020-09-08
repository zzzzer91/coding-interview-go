package main

import "strings"

func main() {

}

func reverseWords(s string) string {
	var stack []string
	foundWord := false
	pos := 0
	for i, _ := range s {
		if s[i] == ' ' {
			if foundWord {
				stack = append(stack, s[pos:i])
				foundWord = false
			}
			continue
		}
		if !foundWord {
			pos = i
			foundWord = true
		}
	}
	// 字符串结尾没有空格时，需要再判断下
	if foundWord {
		stack = append(stack, s[pos:])
	}
	return strings.Join(reverse(stack), " ")
}

func reverse(strs []string) []string {
	l, r := 0, len(strs)-1
	for l < r {
		strs[l], strs[r] = strs[r], strs[l]
		l++
		r--
	}
	return strs
}
