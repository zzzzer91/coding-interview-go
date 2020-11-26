// 基数排序，必须都是正数

package main

import "fmt"

func main() {
	// 都是正数
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 99, 10, 1111}
	fmt.Println(RadixSort(a))
}

func RadixSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	buf := make([]int, n)
	maxVal := max(nums...)
	for exp := 1; exp <= maxVal; exp *= 10 {
		cnt := [10]int{}
		for _, v := range nums {
			digit := v / exp % 10
			cnt[digit]++
		}
		for i := 1; i < 10; i++ { // 用于下面分散到 buf 中， cnt[10] == len(nums)
			cnt[i] += cnt[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			digit := nums[i] / exp % 10
			buf[cnt[digit]-1] = nums[i]
			cnt[digit]--
		}
		copy(nums, buf)
	}
	return nums
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}
