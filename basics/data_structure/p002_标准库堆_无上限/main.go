// 使用标准库自带的堆

package main

import (
	"container/heap"
	"fmt"
)

// 要实现 `heap.Interface` 接口
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

// 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
// 注意大顶堆和小顶堆的区别是，当前数组首元素是堆中最大值还是最小值
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	// 注意这个堆是没有容量上限的，
	// 要想有上限，应该定义一个结构体，然后有个字段表示容量
	h := MinHeap{1, 5, 4, 7, 2, 6, 3} // 创建 slice
	heap.Init(&h)                     // 初始化 heap，堆为空或没有数据可以不做这一步
	fmt.Println(h)
	fmt.Println(heap.Pop(&h)) // 调用 pop
	fmt.Println(heap.Pop(&h))
	heap.Push(&h, 6) // 调用 push
	fmt.Println(h)
	for len(h) > 0 {
		fmt.Printf("%d ", heap.Pop(&h))
	}
}
