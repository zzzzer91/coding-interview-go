// https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/

package main

func main() {

}

// 将出现过的元素作为下标，将数组中对应元素标记为负的，
// 若数组某元素为正的，说明其下标+1没在数组出现过，就是我们要找的值
func findDisappearedNumbers(nums []int) []int {
	for _, n := range nums {
		idx := abs(n) - 1 // abs 不能少，因为可能已被其他的改成负的
		nums[idx] = -abs(nums[idx])
	}
	var ret []int
	for i, num := range nums {
		if num >= 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

// 交换位置，第一种更好
func findDisappearedNumbers2(nums []int) []int {
	var res []int
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, n := range nums {
		if i != n-1 {
			res = append(res, i+1)
		}
	}
	return res
}
