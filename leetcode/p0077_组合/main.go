// https://leetcode-cn.com/problems/combinations/

package main

func main() {

}

func combine(n int, k int) [][]int {
	var res [][]int
	comb := make([]int, 0, k)
	var dfs func(u int)
	dfs = func(u int) {
		if len(comb) == k {
			temp := make([]int, k)
			copy(temp, comb)
			res = append(res, temp)
			return
		}
		// i <= n-(k-len(comb))+1 剪枝
		// 如果 n = 7, k = 4，
		// 从 5 开始搜索就已经没有意义了，
		// 这是因为：即使把 5 选上，后面的数只有 6 和 7，一共就 3 个候选数，凑不出 4 个数的组合。
		for i := u; i <= n-(k-len(comb))+1; i++ {
			comb = append(comb, i)
			dfs(i + 1)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(1)
	return res
}
