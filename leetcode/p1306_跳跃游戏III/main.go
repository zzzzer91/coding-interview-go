// https://leetcode-cn.com/problems/jump-game-iii/submissions/
// bfs 和 dfs 两种做法

package main

import "container/list"

func main() {

}

// bfs
func canReach(arr []int, start int) bool {
	if arr[start] == 0 {
		return true
	}
	q := list.New()
	visited := make([]bool, len(arr))
	q.PushBack(start)
	visited[start] = true
	for q.Len() > 0 {
		u := q.Remove(q.Front()).(int)
		l, r := u-arr[u], u+arr[u]
		if l >= 0 && !visited[l] {
			if arr[l] == 0 {
				return true
			}
			q.PushBack(l)
			visited[l] = true
		}
		if r < len(arr) && !visited[r] {
			if arr[r] == 0 {
				return true
			}
			q.PushBack(r)
			visited[r] = true
		}
	}
	return false
}

// dfs
func canReach2(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	var dfs func(u int) bool
	dfs = func(u int) bool {
		if u < 0 || u >= len(arr) || visited[u] {
			return false
		}
		if arr[u] == 0 {
			return true
		}
		visited[u] = true
		return dfs(u-arr[u]) || dfs(u+arr[u])
	}
	return dfs(start)
}
