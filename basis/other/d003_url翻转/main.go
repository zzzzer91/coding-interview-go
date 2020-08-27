package main

import (
	"fmt"
	"strings"
)

func reverse(url []byte) {
	l, r := 0, len(url)-1
	for l < r {
		url[l], url[r] = url[r], url[l]
		l++
		r--
	}
}

func reverseUrl(url []byte) {
	pos := 0
	for i := 0; i < len(url); i++ {
		if url[i] == '.' {
			reverse(url[pos:i])
			pos = i + 1
		}
	}
	reverse(url[pos:])
	reverse(url)
}

func reverseUrl2(url string) string {
	var stack []string
	pos := 0
	for i := 0; i < len(url); i++ {
		if url[i] == '.' {
			stack = append(stack, url[pos:i])
			pos = i + 1
		}
	}
	stack = append(stack, url[pos:])
	l, r := 0, len(stack)-1
	for l < r {
		stack[l], stack[r] = stack[r], stack[l]
		l++
		r--
	}
	return strings.Join(stack, ".")
}

func main() {
	url := "www.toutiao.com"
	fmt.Printf("%s\n", reverseUrl2(url))
}
