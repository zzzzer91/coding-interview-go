// https://leetcode-cn.com/problems/freedom-trail/

package main

import "math"

func main() {

}

func findRotateSteps(ring string, key string) int {
	indices := [26][]int{}
	for i := range ring {
		p := ring[i] - 'a'
		indices[p] = append(indices[p], i)
	}

	memo := make([][]int, len(ring))
	for i := range memo {
		memo[i] = make([]int, len(key))
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j == len(key) {
			return 0
		}
		if memo[i][j] > 0 {
			return memo[i][j] // 记忆化递归
		}
		p := key[j] - 'a'
		memo[i][j] = math.MaxInt32
		for _, index := range indices[p] {
			dist := abs(i - index)
			dist = min(dist, len(ring)-dist) // 顺序针和逆时针找个距离短的
			memo[i][j] = min(memo[i][j], dist+dfs(index, j+1))
		}
		return memo[i][j]
	}

	return dfs(0, 0) + len(key)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
