// https://leetcode-cn.com/problems/word-break/

package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	// dp[i]表示字符串s的前i个字符能否拆分成wordDict。
	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordSet := make(map[string]bool, len(wordDict))
	for _, s := range wordDict {
		wordSet[s] = true
	}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
