// https://leetcode-cn.com/problems/palindromic-substrings/

package main

func main() {

}

func countSubstrings(s string) int {
	res := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for j := range s {
		for i := 0; i <= j; i++ {
			if i == j {
				dp[i][j] = true
				res++
			} else if j-i == 1 && s[i] == s[j] {
				dp[i][j] = true
				res++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}
