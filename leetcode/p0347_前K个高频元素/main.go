// https://leetcode-cn.com/problems/top-k-frequent-elements/

package main

import "container/heap"

func main() {

}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	h := make(Heap, 0, k+1)
	for key, val := range m {
		heap.Push(&h, [2]int{key, val})
		if h.Len() > k { // 注意要放在 push 后面，确保弹出的是 k+1 个元素中最小的，而不是 k 个
			heap.Pop(&h)
		}
	}
	res := make([]int, 0, k)
	for _, v := range h {
		res = append(res, v[0])
	}
	return res
}

type Heap [][2]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
