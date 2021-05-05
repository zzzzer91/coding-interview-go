// https://leetcode-cn.com/problems/course-schedule/

package main

import "container/list"

func main() {

}

// BFS
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 邻接矩阵
	adjacency := make([][]int, numCourses)
	// 入度数组
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		// p[0] 依赖 p[1]，即 p[1] -> p[0]
		// 所以 p[1] 作 key，p[0] 作 value
		adjacency[p[1]] = append(adjacency[p[1]], p[0])
		inDegree[p[0]]++
	}
	q := list.New()
	// 首先遍历一遍，把所有入度为 0 的结点加入队列
	for i, n := range inDegree {
		if n == 0 {
			q.PushBack(i)
		}
	}
	count := 0
	for q.Len() > 0 {
		v := q.Remove(q.Front()).(int)
		count++
		for _, successor := range adjacency[v] {
			inDegree[successor] -= 1      // 排除 v 对 successor 的入度
			if inDegree[successor] == 0 { // 入度为 0，则加入队列
				q.PushBack(successor)
			}
		}
	}
	return count == numCourses
}
