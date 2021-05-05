// https://leetcode-cn.com/problems/counting-bits/

package main

func main() {

}

/*
i & (i - 1)可以去掉i最右边的一个1（如果有），
因此 i & (i - 1）是比 i 小的，
而且i & (i - 1)的1的个数已经在前面算过了，
所以i的1的个数就是 i & (i - 1)的1的个数加上1。
*/
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ { //注意要从1开始，0不满足
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
