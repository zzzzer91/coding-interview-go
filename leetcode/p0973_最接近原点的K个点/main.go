// https://leetcode-cn.com/problems/k-closest-points-to-origin/

package main

import (
	"container/heap"
	"math/rand"
)

func main() {

}

type pair struct {
	dist  int
	point []int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].dist > h[j].dist }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 堆
func kClosest(points [][]int, k int) (ans [][]int) {
	h := make(hp, k)
	for i, p := range points[:k] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h) // O(k) 初始化堆
	for _, p := range points[k:] {
		if dist := p[0]*p[0] + p[1]*p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0) // 效率比 pop 后 push 要快
		}
	}
	for _, p := range h {
		ans = append(ans, p.point)
	}
	return
}

// 快速排序思想
func kClosest2(points [][]int, k int) (ans [][]int) {
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})

	var quickSelect func(left, right int)
	quickSelect = func(left, right int) {
		if left == right {
			return
		}
		pivot := points[right] // 取当前区间 [left,right] 最右侧元素作为切分参照
		lessCount := left
		for i := left; i < right; i++ {
			if less(points[i], pivot) {
				points[i], points[lessCount] = points[lessCount], points[i]
				lessCount++
			}
		}
		// 循环结束后，有 lessCount 个元素比 pivot 小
		// 把 pivot 交换到 points[lessCount] 的位置
		// 交换之后，points[lessCount] 左侧的元素均小于 pivot，points[lessCount] 右侧的元素均不小于 pivot
		points[right], points[lessCount] = points[lessCount], points[right]
		if lessCount+1 == k {
			return
		} else if lessCount+1 < k {
			quickSelect(lessCount+1, right)
		} else {
			quickSelect(left, lessCount-1)
		}
	}
	quickSelect(0, len(points)-1)
	return points[:k]
}

func less(p, q []int) bool {
	return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
}
