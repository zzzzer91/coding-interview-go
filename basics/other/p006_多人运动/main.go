// https://github.com/azl397985856/leetcode/issues/347

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type SortData [][]int

func (d SortData) Len() int {
	return len(d)
}

func (d SortData) Less(i, j int) bool {
	return d[i][0] < d[j][0]
}

func (d SortData) Swap(i, j int) {
	d[i], d[j] = d[j], d[j]
}

type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[j]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	tmp := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tmp
}

func main() {
	fmt.Println(np([][]int{{0, 30}, {5, 10}, {15, 20}})) // 2
}

func np(intervals [][]int) int {
	sort.Sort(SortData(intervals))
	var h MinHeap
	heap.Push(&h, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if h[0][1] < intervals[i][0] {
			heap.Pop(&h)
		}
		heap.Push(&h, intervals[i])
	}
	return h.Len()
}
