// https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/

package main

func main() {

}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] == m {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
