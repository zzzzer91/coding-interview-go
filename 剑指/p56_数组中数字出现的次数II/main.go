// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/

package main

func main() {

}

// 如果一个数字出现3次，它的二进制每一位也出现的3次。
// 如果把所有的出现三次的数字的二进制表示的每一位都分别加起来，那么每一位都能被3整除。
// 我们把数组中所有的数字的二进制表示的每一位都加起来。
// 如果某一位能被3整除，那么这一位对只出现一次的那个数的这一肯定为0。
// 如果某一位不能被3整除，那么只出现一次的那个数字的该位置一定为1。
func singleNumber(nums []int) int {
	res := 0
	a := make([]int, 32)
	for _, n := range nums {
		for i := 0; i < 32; i++ {
			a[i] += n >> i & 1
		}
	}
	for i, n := range a {
		if n%3 == 1 {
			res |= 1 << i
		}
	}
	return res
}
