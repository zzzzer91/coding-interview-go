// https://leetcode-cn.com/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/

package main

func main() {

}

func findContinuousSequence(target int) [][]int {
	var res [][]int
	nums := make([]int, target)
	for i := range nums {
		nums[i] = i + 1
	}
	sum := 0
	i, j := 0, 0
	for j < len(nums) {
		if sum < target {
			sum += nums[j]
			j++
		} else {
			if sum == target {
				res = append(res, nums[i:j])
			}
			sum -= nums[i]
			i++
			if i == j {
				break
			}
		}
	}
	return res
}

// 优化
func findContinuousSequence2(target int) [][]int {
	var res [][]int
	nums := make([]int, target+1)
	for i := range nums {
		nums[i] = i
	}
	i, j := 1, 2
	for i < j {
		sum := (i + j) * (j - i + 1) / 2 // 等差数列求和 (a+b)*n/2
		if sum < target {
			j++
		} else {
			if sum == target {
				res = append(res, nums[i:j+1])
			}
			i++
		}
	}
	return res
}
