// 携程笔试
// 从前往后，每层字母组合总数，并判断是否有循环（字母之前出现过）

package main

import "fmt"

func main() {
	// f([]string{"a", "bc", "df", "gh"})
	f([]string{"a", "bc", "d", "eah", "f"})
}

func f(strs []string) {
	res := make([]byte, len(strs))
	m := make(map[byte]int)
	var dfs func(u int, cFlag bool)
	dfs = func(u int, cFlag bool) {
		if u == len(strs) {
			if cFlag {
				fmt.Printf("%s -- circular\n", string(res))
			} else {
				fmt.Printf("%s\n", string(res))
			}
			return
		}
		s := strs[u]
		for j := 0; j < len(s); j++ {
			c := s[j]
			tFlag := cFlag
			if v := m[c]; v > 0 {
				tFlag = true
			}
			m[c]++
			res[u] = c
			dfs(u+1, tFlag)
			m[c]--
		}
	}
	dfs(0, false)
}
