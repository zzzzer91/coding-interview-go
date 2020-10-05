// https://leetcode-cn.com/problems/3sum-closest/

package main

import "sort"

func main() {

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var res int
	minDis := int(1e9)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			sum := nums[i] + nums[l] + nums[r]
			tmp := abs(sum - target)
			if tmp < minDis {
				res = sum
				minDis = tmp
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return sum
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
