package main

func main() {

}

// O(nlogk)
func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	// O(klogk)
	buildHeap(heap)
	// O((n-k)logk)
	for _, num := range nums[k:] {
		if num > heap[0] {
			heap[0] = num
			heapify(heap, 0)
		}
	}
	return heap[0]
}

func buildHeap(heap []int) {
	for i := len(heap)/2 - 1; i >= 0; i-- {
		heapify(heap, i)
	}
}

func heapify(heap []int, idx int) {
	minIdx := idx
	l := 2*idx + 1
	r := l + 1
	if l < len(heap) && heap[minIdx] > heap[l] {
		minIdx = l
	}
	if r < len(heap) && heap[minIdx] > heap[r] {
		minIdx = r
	}
	if minIdx != idx {
		swap(heap, idx, minIdx)
		heapify(heap, minIdx) // 放判断里面，是因为堆已经建成，后面的元素必定符合要求了
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
