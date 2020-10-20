// https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/

package main

func main() {

}

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}
