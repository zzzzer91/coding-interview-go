// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(getLeastNumbers2([]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8))
}

// 利用快排性质来做，最好O(n)，最坏O(n^2)
func getLeastNumbers(arr []int, k int) []int {
	f(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func f(arr []int, i, j, k int) {
	if j-i+1 <= k {
		return
	}
	pivotIdx := partition(arr, i, j)
	if pivotIdx-i > k-1 {
		f(arr, i, pivotIdx-1, k)
	} else if pivotIdx-i < k-1 {
		f(arr, pivotIdx+1, j, k-1-(pivotIdx-i))
	}
}

func partition(arr []int, i, j int) int {
	l, r := i, j
	x := arr[l]
	for l < r {
		for l < r && arr[r] >= x {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] <= x {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = x
	return l
}

func getLeastNumbers2(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	res := MaxHeap{data: arr[:k], capacity: k}
	heap.Init(&res)
	for _, v := range arr[k:] {
		heap.Push(&res, v)
	}
	return res.data
}

type MaxHeap struct {
	data     []int
	capacity int
}

func (h MaxHeap) Len() int {
	return len(h.data)
}

func (h MaxHeap) Less(i, j int) bool {
	return h.data[i] > h.data[j]
}

func (h MaxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *MaxHeap) Push(x interface{}) {
	if h.capacity == h.Len() {
		if x.(int) < h.data[0] {
			heap.Pop(h) // 注意要用标准库的函数，而不能直接用 Pop() 方法
		} else {
			return
		}
	}
	h.data = append(h.data, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	tmp := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return tmp
}
