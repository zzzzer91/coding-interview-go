// 使用标准库自带的堆

package main

import (
	"container/heap"
	"fmt"
)

// 要实现 `heap.Interface` 接口
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) IsEmpty() bool {
	return len(*h) == 0
}

func (h *IntHeap) Size() int {
	return len(*h)
}

func main() {
	h := IntHeap{1, 5, 4, 7, 2, 6, 3} // 创建 slice
	heap.Init(&h)                     // 初始化 heap
	fmt.Println(h)
	fmt.Println(heap.Pop(&h)) // 调用 pop
	heap.Push(&h, 6)          // 调用 push
	fmt.Println(h)
	for len(h) > 0 {
		fmt.Printf("%d ", heap.Pop(&h))
	}
}
