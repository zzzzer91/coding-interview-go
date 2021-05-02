// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/
// 注意题目中的有序条件，首先想到二分

package main

func main() {

}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r { // 注意等于条件必须加上，举例：[0]
		m := l + (r-l)>>1
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
