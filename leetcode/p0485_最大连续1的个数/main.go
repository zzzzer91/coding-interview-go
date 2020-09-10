package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	count := 0
	for _, n := range nums {
		if n == 1 {
			count++
		} else {
			res = max(res, count)
			count = 0
		}
	}
	res = max(res, count) // 数组最后一个元素是 1 的情况
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
