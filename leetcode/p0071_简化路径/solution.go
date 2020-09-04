package p0071_简化路径

import "strings"

func simplifyPath(path string) string {
	var stack []string
	temp := strings.Split(path, "/")
	for _, s := range temp {
		if s == "" || s == "." {
			continue
		}
		if s == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, s)
	}
	return "/" + strings.Join(stack, "/")
}
