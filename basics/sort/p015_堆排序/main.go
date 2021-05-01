// 堆排序
// 属于不稳定排序
// 平均时 O(nlog₂n)，最坏时 O(nlog₂n)，最好时 O(nlog₂n)
// 空 O(1)

package main

import (
	"fmt"
)

func main() {
	a := []int{-1, 3, -5, 7, -9, 2, -4, 6, -8, 10, 0}
	heapSort(a)
	fmt.Println(a)
}

// 堆排序的基本思想是：
// 将待排序序列构造成一个大顶堆，
// 此时，整个序列的最大值就是堆顶的根节点。
// 将其与末尾元素进行交换，此时末尾就为最大值。
// 然后将剩余n-1个元素重新构造成一个堆，
// 这样会得到n个元素的次小值。
// 如此反复执行，便能得到一个有序序列了。
func heapSort(a []int) {
	initHeap(a)
	for i := len(a) - 1; i >= 1; i-- {
		swap(a, 0, i)
		heapify(a[:i], 0)
	}
}

// initHeap 初始化堆
func initHeap(a []int) {
	// 从右往左第一个非叶子结点的坐标 len(a)/2-1
	// 记住这个规律
	for i := len(a)/2 - 1; i >= 0; i-- {
		// 从第一个非叶子节点往上调整
		heapify(a, i)
	}
}

// 大顶堆
// 大顶堆：a[p] >= a[lc] && a[p] >= a[rc]
// 小顶堆：a[p] <= a[lc] && a[p] <= a[rc]
// heapify 调整堆，从上往下
func heapify(a []int, i int) {
	// i, lc, rc 三个点中找到最大那个
	maxIdx := i
	// 设当前节点坐标为 i，
	// 其 lChild 为 2i+1，rChild 为 2i+2。
	// 其 parent 为 (i-1)/2，当然这里没用到
	lc := i*2 + 1
	rc := lc + 1
	// 如果 lc >= len(a)，证明没有叶子节点
	if lc < len(a) && a[maxIdx] < a[lc] {
		maxIdx = lc
	}
	if rc < len(a) && a[maxIdx] < a[rc] {
		maxIdx = rc
	}
	if i != maxIdx {
		swap(a, i, maxIdx)
		heapify(a, maxIdx) // 父节点和子节点发生了交换，那么要继续递归调整子节点，使之符合堆条件
	}
}

// swap 交换数组中两个元素值
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
