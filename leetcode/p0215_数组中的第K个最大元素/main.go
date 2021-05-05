// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

package main

func main() {

}

// O(nlogk)
func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	// O(klogk)
	initHeap(heap)
	// O((n-k)logk)
	for _, n := range nums[k:] {
		if n > heap[0] {
			heap[0] = n
			heapify(heap, 0)
		}
	}
	return heap[0]
}

func initHeap(a []int) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		heapify(a, i)
	}
}

func heapify(a []int, i int) {
	minIdx := i
	lc := i*2 + 1
	rc := lc + 1
	if lc < len(a) && a[minIdx] > a[lc] {
		minIdx = lc
	}
	if rc < len(a) && a[minIdx] > a[rc] {
		minIdx = rc
	}
	if minIdx != i {
		a[minIdx], a[i] = a[i], a[minIdx]
		heapify(a, minIdx)
	}
}
