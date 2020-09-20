// https://leetcode-cn.com/problems/jump-game-iv/

package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(minJumps([]int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}))
}

func minJumps(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	res := int(1e9)
	count := 0
	q := list.New()
	q.PushBack(0)
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		if u == 0 {
			res = min(res, count)
		}
		for i, n := range arr {
			if (u-1 >= 0 && i == u-1) || (u+1 < len(arr) && i == u+1) || n == arr[u] {
				q.PushBack(i)
			}
		}
		count++
	}
	return res
}

func minJumps2(arr []int) int {
	visited := make([]bool, len(arr))
	var dfs func(u int) int
	dfs = func(u int) int {
		if u == len(arr)-1 {
			return 0
		}
		visited[u] = true
		var minCount int = 1e9
		for i, n := range arr {
			if visited[i] {
				continue
			}
			if (u-1 >= 0 && i == u-1) || (u+1 < len(arr) && i == u+1) || n == arr[u] {
				minCount = min(minCount, dfs(i)+1)
			}
		}
		return minCount
	}
	return dfs(0)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
