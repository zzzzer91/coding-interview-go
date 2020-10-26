// https://leetcode-cn.com/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/

package main

import "container/list"

func main() {

}

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
