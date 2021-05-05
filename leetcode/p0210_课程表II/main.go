// https://leetcode-cn.com/problems/course-schedule-ii/

package main

import "container/list"

func main() {

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	var res []int
	adjacency := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, p := range prerequisites {
		adjacency[p[1]] = append(adjacency[p[1]], p[0])
		inDegree[p[0]]++
	}
	q := list.New()
	for i, n := range inDegree {
		if n == 0 {
			q.PushBack(i)
		}
	}
	for q.Len() > 0 {
		v := q.Remove(q.Front()).(int)
		res = append(res, v)
		for _, successor := range adjacency[v] {
			inDegree[successor]--
			if inDegree[successor] == 0 {
				q.PushBack(successor)
			}
		}
	}
	if len(res) != numCourses {
		return nil
	}
	return res
}
