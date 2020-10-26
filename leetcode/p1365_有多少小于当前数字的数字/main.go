// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

package main

func main() {

}

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	const N = 101
	count := make([]int, N) // 0 <= nums[i] <= 100
	for _, v := range nums {
		count[v]++
	}
	for i := 1; i < N; i++ {
		count[i] += count[i-1]
	}
	for i, v := range nums {
		if v > 0 {
			res[i] = count[v-1] // 这里要减1
		}
	}
	return res
}
