// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

package main

func main() {

}

func singleNumbers(nums []int) []int {
	n := 0
	for _, v := range nums {
		n ^= v
	}
	n = n & (-n) // 找到两个不同的数最右边不同的那一位，以此来分组
	res := make([]int, 2)
	for _, v := range nums {
		if v&n != 0 { // 分组
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
