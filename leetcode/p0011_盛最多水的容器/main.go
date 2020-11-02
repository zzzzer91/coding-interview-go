// https://leetcode-cn.com/problems/container-with-most-water/

package main

func main() {

}

func maxArea(height []int) int {
	res := 0
	l, r := 0, len(height)-1
	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
