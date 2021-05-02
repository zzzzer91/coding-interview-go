// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/

package main

func main() {

}

// 如果我们可以把所有数字分成两组，使得：
// 两个只出现一次的数字在不同的组中；
// 相同的数字会被分到相同的组中。
// 那么对两个组分别进行异或操作，即可得到答案的两个数字。这是解决这个问题的关键。
func singleNumbers(nums []int) []int {
	res := make([]int, 2)
	n := 0
	for _, v := range nums {
		n ^= v
	}
	n = n & (-n) // 得到最右边一个 1，其余都是 0，以此来分组
	for _, v := range nums {
		if v&n != 0 { // 相同的数字必然被分在同一组
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
