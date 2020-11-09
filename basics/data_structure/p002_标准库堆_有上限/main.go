// 使用标准库自带的堆

package main

import (
	"container/heap"
	"fmt"
)

type MinHeap struct {
	data     []int
	capacity int
}

func (h MinHeap) Len() int {
	return len(h.data)
}

func (h MinHeap) Less(i, j int) bool {
	return h.data[i] > h.data[j] // 大顶堆
}

func (h MinHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MinHeap) Push(x interface{}) {
	if h.capacity == h.Len() {
		x := x.(int)
		if x < h.data[0] {
			h.data[0] = x
			heap.Fix(h, 0) // 重新调整堆，效率比 pop 后 push 要快
		}
		return
	}
	h.data = append(h.data, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	tmp := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return tmp
}

func main() {
	fmt.Println(getLeastNumbers([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
}

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	res := MinHeap{data: arr[:k], capacity: k}
	heap.Init(&res)
	for _, v := range arr[k:] {
		heap.Push(&res, v)
	}
	return res.data
}
