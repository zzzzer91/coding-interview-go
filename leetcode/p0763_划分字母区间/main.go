// https://leetcode-cn.com/problems/partition-labels/

package main

func main() {

}

func partitionLabels(S string) []int {
	var res []int
	var m [26]int
	for i := range S {
		m[S[i]-'a'] = i
	}
	start, end := 0, 0
	for i := range S {
		end = max(end, m[S[i]-'a'])
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
