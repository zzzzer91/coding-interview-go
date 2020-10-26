// https://leetcode-cn.com/problems/sliding-window-maximum/

package main

import "container/list"

func main() {

}

// 标准库是链表，性能比较低
func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	q := list.New()
	for i, n := range nums {
		if q.Len() != 0 && i-q.Front().Value.(int) > k-1 {
			q.Remove(q.Front())
		}
		for q.Len() != 0 && n >= nums[q.Back().Value.(int)] {
			q.Remove(q.Back())
		}
		q.PushBack(i)
		if i >= k-1 {
			res = append(res, nums[q.Front().Value.(int)])
		}
	}
	return res
}

// 使用数组模拟，快大概30%
func maxSlidingWindow2(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	q := make([]int, 0, len(nums))
	for i, n := range nums {
		if len(q) != 0 && i-q[0] > k-1 {
			q = q[1:]
		}
		for len(q) != 0 && n >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
