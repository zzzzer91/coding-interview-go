// https://leetcode-cn.com/problems/word-break-ii/

package main

import "strings"

func main() {

}

// 记忆化，memo[i] 代表记录 s[i:] 的结果
func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	memo := make([][][]string, n)
	var backtrack func(u int) [][]string
	backtrack = func(u int) [][]string {
		if memo[u] != nil {
			return memo[u]
		}
		// 很奇怪，这里不能用 var wordList [][]string，否则会超时
		wordList := [][]string{}
		for i := u + 1; i < n; i++ {
			word := s[u:i]
			if wordSet[word] {
				for _, nextWords := range backtrack(i) {
					// append([]string{word}, nextWords...)，数组头部追加
					wordList = append(wordList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[u:]
		if wordSet[word] {
			wordList = append(wordList, []string{word})
		}
		memo[u] = wordList
		return wordList
	}
	var res []string
	for _, words := range backtrack(0) {
		res = append(res, strings.Join(words, " "))
	}
	return res
}

// dfs 超时，忽略了以下情况
// s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
// wordDict = ["a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa"]
func wordBreak2(s string, wordDict []string) []string {
	var res []string
	var comb []string
	wordSet := make(map[string]bool, len(s))
	for _, v := range wordDict {
		wordSet[v] = true
	}
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(s) {
			res = append(res, strings.Join(comb, " "))
			return
		}
		for i := u + 1; i <= len(s); i++ {
			if wordSet[s[u:i]] {
				comb = append(comb, s[u:i])
				dfs(i)
				comb = comb[:len(comb)-1]
			}
		}
	}
	dfs(0)
	return res
}
