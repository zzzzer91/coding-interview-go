// https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/

package main

import "fmt"

func main() {
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
}

func findPairs(nums []int, k int) int {
	numSet := make(map[int]struct{})
	diffSet := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := numSet[v-k]; ok {
			diffSet[v-k] = struct{}{}
		}
		if _, ok := numSet[v+k]; ok {
			diffSet[v] = struct{}{}
		}
		numSet[v] = struct{}{}
	}
	return len(diffSet)
}

func findPairs2(nums []int, k int) int {
	var res int
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	if k == 0 {
		for _, v := range m {
			if v >= 2 {
				res++
			}
		}
	} else {
		for key := range m {
			if m[key+k] > 0 {
				res++
			}
		}
	}
	return res
}
