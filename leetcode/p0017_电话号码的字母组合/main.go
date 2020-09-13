// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

package main

func main() {

}

func letterCombinations(digits string) []string {
	var res []string
	if len(digits) == 0 {
		return res
	}
	m := []string{
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var comb []byte
	var dfs func(u int)
	dfs = func(u int) {
		if u == len(digits) {
			res = append(res, string(comb))
			return
		}
		s := m[digits[u]-'2']
		for i := 0; i < len(s); i++ {
			comb = append(comb, s[i])
			dfs(u + 1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(0)
	return res
}
